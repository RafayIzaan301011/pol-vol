package server

import (
	"net/http"
	"os"
	"pvms/middleware"

	"github.com/gin-gonic/gin"
)

// RunHTTPServer starts an HTTP server on the port specified by the PORT environment variable.
func RunHTTPServer(createHandler func(router *gin.Engine) http.Handler) {
	RunHTTPServerOnAddr(":"+os.Getenv("PORT"), createHandler)
}

// RunHTTPServerOnAddr starts an HTTP server on the specified address.
func RunHTTPServerOnAddr(addr string, createHandler func(router *gin.Engine) http.Handler) {
	router := gin.Default()

	// global middlewares
	setMiddlewares(router)

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

	middleware.AddCorsMiddleware(router)
}
