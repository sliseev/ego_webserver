package server

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
	l      *zap.Logger
}

func middleware(db *gorm.DB, l *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Set("log", l)
		c.Next()
	}
}

func zapLoggerWrapper(l *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		timestamp := time.Now()
		latency := timestamp.Sub(start)
		client := c.ClientIP()
		method := c.Request.Method
		status := c.Writer.Status()
		size := c.Writer.Size()

		l.Info("Request",
			zap.Int("Status", status),
			zap.Duration("Latency", latency),
			zap.String("Client", client),
			zap.String("Method", method),
			zap.String("Path", path),
			zap.String("Params", raw),
			zap.Int("Size", size),
		)
	}
}

func New(db *gorm.DB, l *zap.Logger) *Server {
	router := gin.New()
	router.Use(zapLoggerWrapper(l), gin.Recovery())
	router.Use(middleware(db, l))

	router.POST("/driver", createDriver)
	router.GET("/driver", getDrivers)
	router.GET("/driver/:id", getDriver)
	router.PUT("/driver/:id", updateDriver)
	router.DELETE("/driver/:id", deleteDriver)
	router.GET("/driver/count", getDriversCount)

	return &Server{router, l}
}

func (s *Server) Run(host string, port int) error {
	addr := fmt.Sprintf("%s:%d", host, port)
	return s.router.Run(addr)
}
