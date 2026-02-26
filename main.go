package main

import (
	"log"
	"os"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User model
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Email     string    `gorm:"unique;not null"`
	Phone     string    `gorm:"unique"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

var (
	db    *gorm.DB
	store *session.Store
)

func main() {
	// Initialize database
	var err error
	db, err = gorm.Open(sqlite.Open("forge.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate
	db.AutoMigrate(&User{})

	// Create default user: +1 (555) 123-4567 / password123
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	db.FirstOrCreate(&User{
		Email:    "forge@example.com",
		Phone:    "+1 (555) 123-4567",
		Password: string(hashedPassword),
	}, User{Email: "forge@example.com"})

	// Initialize session store
	store = session.New(session.Config{
		Expiration: 24 * time.Hour,
	})

	// Initialize Fiber
	app := fiber.New(fiber.Config{
		Views:       nil,
		ViewsLayout: "",
	})

	// Static file serving for templates
	app.Static("/", "./templates")

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./templates/login.html")
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.SendFile("./templates/login.html")
	})

	app.Get("/register", func(c *fiber.Ctx) error {
		return c.SendFile("./templates/register.html")
	})

	app.Get("/dashboard", authRequired, func(c *fiber.Ctx) error {
		return c.SendFile("./templates/dashboard.html")
	})

	// API Routes
	app.Post("/api/login", handleLogin)
	app.Post("/api/register", handleRegister)
	app.Post("/api/logout", handleLogout)
	app.Get("/api/check-email", handleCheckEmail)
	app.Get("/api/user", authRequired, handleGetUser)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🔨 Forge Authentication running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}

// Middleware: Auth required
func authRequired(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	userID := sess.Get("user_id")
	if userID == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	return c.Next()
}

// Handler: Login
func handleLogin(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Find user
	var user User
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Create session
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Session error"})
	}

	sess.Set("user_id", user.ID)
	sess.Set("email", user.Email)
	if err := sess.Save(); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save session"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Login successful",
		"user": fiber.Map{
			"email": user.Email,
			"phone": user.Phone,
		},
	})
}

// Handler: Register
func handleRegister(c *fiber.Ctx) error {
	type RegisterRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Validate email
	if req.Email == "" || len(req.Email) < 5 {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid email address"})
	}

	// Validate password
	if len(req.Password) < 6 {
		return c.Status(400).JSON(fiber.Map{"error": "Password must be at least 6 characters"})
	}

	// Check if user exists
	var existingUser User
	if err := db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return c.Status(409).JSON(fiber.Map{"error": "Email already registered"})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	// Create user
	user := User{
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	// Create session
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Session error"})
	}

	sess.Set("user_id", user.ID)
	sess.Set("email", user.Email)
	if err := sess.Save(); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save session"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Registration successful",
		"user": fiber.Map{
			"email": user.Email,
		},
	})
}

// Handler: Logout
func handleLogout(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Session error"})
	}

	if err := sess.Destroy(); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to destroy session"})
	}

	return c.JSON(fiber.Map{"success": true, "message": "Logged out"})
}

// Handler: Check email availability
func handleCheckEmail(c *fiber.Ctx) error {
	email := c.Query("email")
	if email == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Email required"})
	}

	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return c.JSON(fiber.Map{"available": true})
	}

	return c.JSON(fiber.Map{"available": false})
}

// Handler: Get current user
func handleGetUser(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Session error"})
	}

	userID := sess.Get("user_id")
	email := sess.Get("email")

	return c.JSON(fiber.Map{
		"id":    userID,
		"email": email,
	})
}
