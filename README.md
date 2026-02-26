# 🔨 The Forge - Blacksmith Authentication System

**Tempered by fire, secured by steel**

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://go.dev/)
[![Fiber](https://img.shields.io/badge/Fiber-v2-00ACD7?style=for-the-badge&logo=go)](https://gofiber.io/)
[![GORM](https://img.shields.io/badge/GORM-SQLite-4B8BBE?style=for-the-badge)](https://gorm.io/)
[![Alpine.js](https://img.shields.io/badge/Alpine.js-3.x-8BC0D0?style=for-the-badge&logo=alpine.js)](https://alpinejs.dev/)
[![Tailwind CSS](https://img.shields.io/badge/Tailwind-CSS-38B2AC?style=for-the-badge&logo=tailwind-css)](https://tailwindcss.com/)
[![Render](https://img.shields.io/badge/Deploy%20on-Render-46E3B7?style=for-the-badge&logo=render)](https://render.com)

---

## 🔥 Overview

**The Forge** is an immersive 3D authentication system inspired by the ancient art of blacksmithing. Perfect for celebrating strength, craftsmanship, and mastery — ideal for men's day themes (Defender's Day, February 23rd) with universal appeal.

Built with **Go Fiber**, **GORM**, **SQLite**, **Alpine.js**, and **Tailwind CSS** — this system features stunning visual effects including molten metal glow, flying sparks, forge flames, and hammer strike animations.

### ✨ Key Features

- 🔨 **Immersive 3D Effects**: Pulsing forge glow, animated sparks, dynamic flames
- ⚡ **Smart Validation**: Real-time email availability check, password strength meter
- 🔐 **Secure**: Bcrypt password hashing, session-based authentication
- 🎨 **Beautiful UI**: Glassmorphism cards, gradient effects, responsive design
- 🚀 **No CGO**: Uses pure Go SQLite driver (`github.com/glebarez/sqlite`)
- 💨 **Async Operations**: No page reloads, smooth UX with Alpine.js
- 🎯 **Production Ready**: Optimized for Render deployment

---

## 🎬 Demo

**Default credentials:**
- 📧 Email: `forge@example.com`
- 🔑 Password: `password123`

---

## 🛠️ Tech Stack

| Technology | Purpose |
|------------|---------|
| **Go 1.21+** | Backend runtime |
| **Fiber v2** | Fast web framework |
| **GORM** | ORM for database operations |
| **SQLite** | Lightweight embedded database |
| **Alpine.js** | Reactive frontend framework |
| **Tailwind CSS** | Utility-first CSS framework |
| **Bcrypt** | Password hashing |

---

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher
- Git

### Installation

```bash
git clone https://github.com/smart-developer1791/go-fiber-auth-forge
cd go-fiber-auth-forge
```

Initialize dependencies and run:

```bash
go mod tidy
go run .
```

Open [http://localhost:3000](http://localhost:3000) 🔥

---

## 📂 Project Structure

```text
go-fiber-auth-forge/
├── main.go              # Core application with Fiber routes, GORM models
├── templates/
│   ├── login.html       # Login page with forge effects
│   ├── register.html    # Registration with password strength meter
│   └── dashboard.html   # Protected dashboard area
├── go.mod               # Go dependencies
├── .gitignore          # Git ignore rules
├── render.yaml         # Render deployment config
└── README.md           # This file
```

---

## 🎨 Features Breakdown

### 🔐 Authentication

- **Login**: Email + password with async validation
- **Registration**: Email uniqueness check, password strength meter, confirmation matching
- **Sessions**: Secure session-based auth with Fiber middleware
- **Logout**: Clean session destruction

### 🎭 Visual Effects

| Effect | Description |
|--------|-------------|
| **Forge Glow** | Pulsing orange/red radial gradient simulating hot metal |
| **Flying Sparks** | Dynamic particles rising from the bottom with random trajectories |
| **Flame Animation** | Flickering flame effects on both sides |
| **Hammer Strike** | Button animation on submit (rotation + translation) |
| **Metal Shine** | Gradient overlays creating metallic luster |

### 🧠 Smart Validation

- ✅ **Email**: Format validation + real-time availability check via API
- ✅ **Password**: 6+ characters, strength meter (5 levels: Weak → Excellent)
- ✅ **Confirmation**: Real-time matching with visual feedback
- ✅ **Async**: All validations happen without page reload

---

## 🔧 API Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| `GET` | `/` | Redirect to login | ❌ |
| `GET` | `/login` | Login page | ❌ |
| `GET` | `/register` | Registration page | ❌ |
| `GET` | `/dashboard` | Protected dashboard | ✅ |
| `POST` | `/api/login` | Login handler | ❌ |
| `POST` | `/api/register` | Registration handler | ❌ |
| `POST` | `/api/logout` | Logout handler | ✅ |
| `GET` | `/api/check-email` | Email availability | ❌ |
| `GET` | `/api/user` | Get current user | ✅ |

---

## 🎯 Password Strength Levels

| Strength | Criteria |
|----------|----------|
| **Weak** | Less than 6 characters |
| **Fair** | 6+ characters |
| **Good** | 10+ characters OR mixed case |
| **Strong** | Mixed case + numbers |
| **Excellent** | Mixed case + numbers + special characters |

---

## 🌍 Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `3000` | Server port |

---

## 🎨 Customization

### Colors

The forge theme uses a warm palette:

```css
/* Primary: Orange/Red gradient */
from-orange-500 via-red-500 to-orange-600

/* Background: Dark zinc/neutral */
from-zinc-900 via-neutral-900 to-black

/* Accents: Metallic silver */
border-orange-500/20
```

### Animation Timing

Adjust animation speeds in `<style>` blocks:

```css
@keyframes forge-glow {
    /* Change duration: 3s → 2s for faster pulsing */
    animation: forge-glow 2s ease-in-out infinite;
}
```

---

## 🎖️ Perfect For

- 🎉 **Defender's Day** (February 23rd) celebrations
- 👨 **Men's Day** themed applications
- ⚔️ **Strength & Craftsmanship** branding
- 🏭 **Industrial/Manufacturing** platforms
- 🎮 **Gaming** authentication (RPG/crafting themes)

---

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## 🙏 Acknowledgments

- 🔨 Inspired by the ancient art of blacksmithing
- 🔥 Tailwind CSS for the amazing utility classes
- ⚡ Alpine.js for reactive simplicity
- 🚀 Fiber for blazing fast Go web framework
- 💎 GORM for elegant database operations

---

## Deploy in 10 seconds

[![Deploy to Render](https://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy)
