package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// PageData represents the data passed to templates
type PageData struct {
	Title           string
	Description     string
	ErrorCode       string
	ErrorMessage    string
	ErrorDescription string
}

func main() {
	// Set Gin to release mode for production
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Load HTML templates
	r.LoadHTMLGlob("templates/*.html")

	// Serve static files
	r.Static("/static", "./static")

	// Middleware for logging and security headers
	r.Use(loggingMiddleware())
	r.Use(securityHeadersMiddleware())

	// Routes
	r.GET("/", homeHandler)
	r.GET("/health", healthHandler)

	// 404 handler
	r.NoRoute(notFoundHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(r.Run(":" + port))
}

func homeHandler(c *gin.Context) {
	data := PageData{
		Title:       "ZoeLabs",
		Description: "Zoe Hanke - IT Specialist & Full Stack Developer",
	}
	c.HTML(http.StatusOK, "home.html", data)
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

func notFoundHandler(c *gin.Context) {
	data := PageData{
		Title:           "404 - Page Not Found",
		ErrorCode:       "404",
		ErrorMessage:    "Page Not Found",
		ErrorDescription: "The requested page does not exist.",
	}
	c.HTML(http.StatusNotFound, "error.html", data)
}

func loggingMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s %s %v\n",
			param.Method,
			param.Path,
			param.Latency,
		)
	})
}

func securityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Next()
	}
}
