# ZoeLabs.de

A minimalist, terminal-themed personal website for Zoe Hanke built with Go and vanilla HTML/CSS/JavaScript.

## 🚀 Features

- **Terminal UI**: Interactive command-line interface
- **Dark Mode**: Sleek dark theme with Trans pride colors
- **Responsive Design**: Works on all devices
- **Portfolio Gateway**: Recruiter-friendly commands for experience, skills, AI-assisted engineering, and public portfolio links
- **Fast Loading**: Compressed main document stays under 14kb, keeping the first view small enough for a typical initial TCP window
- **Docker Support**: Easy deployment with Docker Compose
- **Hot Reload**: Development with Air for instant feedback

## 🛠️ Tech Stack

- **Backend**: Go 1.24 with Gin framework
- **Frontend**: Vanilla HTML/CSS/JavaScript
- **Templating**: Go's `html/template`
- **Containerization**: Docker & Docker Compose
- **Development**: Air for hot reload

## 📁 Project Structure

```
zoelabs.de/
├── main.go                 # Go server entry point
├── go.mod                  # Go module definition
├── go.sum                  # Go dependencies checksum
├── Dockerfile              # Docker image definition
├── docker-compose.yml      # Docker Compose configuration
├── Makefile                # Development commands
├── .air.toml              # Air hot reload configuration
├── templates/              # HTML templates
│   ├── base.html          # Base template
│   ├── home.html          # Home page template
│   └── error.html         # Error page template
├── static/                 # Static assets
│   └── css/
│       └── style.css      # Main stylesheet
└── README.md              # This file
```

## 🚀 Quick Start

### Prerequisites

- Go 1.24+
- Docker & Docker Compose (for containerized deployment)

### Local Development

1. **Clone the repository**
   ```bash
   git clone git@github.com:zoeliterally/zoelabs.de.git
   cd zoelabs.de
   ```

2. **Install dependencies**
   ```bash
   make deps
   ```

3. **Start development server**
   ```bash
   make dev
   ```

4. **Open in browser**
   ```
   http://localhost:8080
   ```

### Alternative Commands

- `make hot` - Start with hot reload (Air)
- `make build` - Build the application
- `make clean` - Clean build artifacts
- `make help` - Show all available commands

### Docker Development

```bash
# Start with Docker Compose
docker compose up --build

# Or use the dev service
docker compose up dev
```

## 🐳 Production Deployment

### Docker Compose

```bash
# Build and start production
docker compose -f docker-compose.yml up --build -d

# Stop services
docker compose down
```

### Manual Build

```bash
# Build binary
make build

# Run binary
./zoelabs.de
```

## 🎮 Terminal Commands

The website features an interactive terminal with these commands:

- `help` - Show available commands
- `about` - About Zoe Hanke
- `contact` - Contact information
- `projects` - Current projects
- `portfolio` - Public portfolio link
- `ai` - AI-assisted engineering workflows
- `skills` - Technical skills
- `experience` - Work experience
- `leadership` - Technical leadership profile
- `timeline` - Career timeline
- `github` - GitHub profile
- `linkedin` - LinkedIn profile
- `philosophy` - Working philosophy
- `focus` - Current focus areas
- `clear` - Clear terminal
- `reboot` - Restart the system
- `ls` - List files
- `pwd` - Show current directory
- `whoami` - Show username

## 🎨 Design

- **Theme**: Dark mode with terminal aesthetics
- **Colors**: Trans pride gradient (#5BCEFA, #F5A9B8, #FFFFFF)
- **Font**: DejaVu Sans Mono for terminal authenticity
- **Layout**: 4:3 aspect ratio terminal window
- **Responsive**: Adapts to all screen sizes

## 🔧 Configuration

### Environment Variables

- `PORT` - Server port (default: 8080)
- `GIN_MODE` - Gin mode (default: release)

### Air Configuration

Hot reload is configured in `.air.toml`:
- Watches Go files and templates
- Automatically restarts on changes
- Excludes build artifacts

## 📝 License

This project is private and proprietary.

## 👤 Author

**Zoe Hanke** - Head of IT · Platform Engineer · Full Stack Developer
- Website: [zoelabs.de](https://zoelabs.de)
- GitHub: [@zoeliterally](https://github.com/zoeliterally)
- Portfolio: [github.com/zoeliterally/portfolio](https://github.com/zoeliterally/portfolio)

---

Built with ❤️ and ☕
