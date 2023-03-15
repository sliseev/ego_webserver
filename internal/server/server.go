package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
}

func DbMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func New(db *gorm.DB) *Server {
	router := gin.Default()

	router.Use(DbMiddleware(db))

	router.POST("/driver", createDriver)
	router.GET("/driver", getDrivers)
	router.GET("/driver/:id", getDriver)
	router.PUT("/driver/:id", updateDriver)
	router.DELETE("/driver/:id", deleteDriver)

	return &Server{
		router: router,
	}
}

func (s *Server) Run(host string, port int) error {
	addr := fmt.Sprintf("%s:%d", host, port)
	return s.router.Run(addr)
}
