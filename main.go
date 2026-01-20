package main

import (
	"compress/gzip"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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

	// Serve static files with cache headers
	r.StaticFS("/static", http.Dir("./static"))
	r.Use(staticCacheMiddleware())

	// Middleware for logging, compression, and security headers
	r.Use(loggingMiddleware())
	r.Use(gzipMiddleware())
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

func staticCacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add cache headers for static assets
		path := c.Request.URL.Path
		if len(path) >= 8 && path[:8] == "/static/" {
			c.Header("Cache-Control", "public, max-age=31536000, immutable")
			c.Header("Expires", time.Now().Add(365*24*time.Hour).Format(http.TimeFormat))
		}
		c.Next()
	}
}

func gzipMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if client accepts gzip encoding
		if !strings.Contains(c.GetHeader("Accept-Encoding"), "gzip") {
			c.Next()
			return
		}

		// Skip compression for already compressed or binary content
		path := c.Request.URL.Path
		if strings.HasSuffix(path, ".svg") || strings.HasSuffix(path, ".ico") || strings.HasSuffix(path, ".png") || strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".jpeg") || strings.HasSuffix(path, ".gif") || strings.HasSuffix(path, ".webp") {
			c.Next()
			return
		}

		// Set gzip writer
		c.Header("Content-Encoding", "gzip")
		c.Header("Vary", "Accept-Encoding")
		
		gz := gzip.NewWriter(c.Writer)
		defer func() {
			gz.Close()
		}()

		c.Writer = &gzipResponseWriter{ResponseWriter: c.Writer, Writer: gz}
		c.Next()
	}
}

type gzipResponseWriter struct {
	gin.ResponseWriter
	Writer *gzip.Writer
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *gzipResponseWriter) WriteString(s string) (int, error) {
	return w.Writer.Write([]byte(s))
}
