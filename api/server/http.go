package server

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// RunHTTPServer2 starts an HTTP server on the port specified by the PORT environment variable.
func RunHTTPServer2(createHandler func(router *gin.Engine) http.Handler) {
	RunHTTPServerOnAddr(":"+os.Getenv("PORT"), createHandler)
}

// RunHTTPServerOnAddr starts an HTTP server on the specified address.
func RunHTTPServerOnAddr(addr string, createHandler func(router *gin.Engine) http.Handler) {
	router := gin.Default()

	// global middlewares

	handler := createHandler(router)

	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

// setMiddlewares sets global middlewares for the Gin router.
func setMiddlewares(router *gin.Engine) {
	// Add global middlewares here
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("X-Frame-Options", "DENY")
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		c.Writer.Header().Set("Referrer-Policy", "no-referrer")
		c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'")
		c.Next()
	})

	addCorsMiddleware(router)

}

// --- CORS SETUP ---
func addCorsMiddleware(router *gin.Engine) {
	allowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ";")
	if len(allowedOrigins) == 0 {
		return
	}

	config := cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(config))
}
