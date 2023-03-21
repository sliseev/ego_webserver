package server

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
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

func promLowCardinalityAdapter() func(c *gin.Context) string {
	return func(c *gin.Context) string {
		url := c.Request.URL.Path
		for _, p := range c.Params {
			if p.Key == "id" {
				url = strings.Replace(url, p.Value, ":id", 1)
				break
			}
		}
		return url
	}
}

func New(db *gorm.DB, l *zap.Logger) *Server {
	router := gin.New()
	router.Use(zapLoggerWrapper(l), gin.Recovery())
	router.Use(middleware(db, l))

	p := ginprometheus.NewPrometheus("gin")
	p.ReqCntURLLabelMappingFn = promLowCardinalityAdapter()
	p.Use(router)

	router.POST("/driver", createDriver)
	router.GET("/driver", getDrivers)
	router.GET("/driver/:id", getDriver)
	router.PUT("/driver/:id", updateDriver)
	router.DELETE("/driver/:id", deleteDriver)
	router.GET("/driver/count", getDriversCount)

	router.POST("/testapi/drivers", generateDrivers)

	return &Server{router, l}
}

func (s *Server) Run(host string, port int) error {
	addr := fmt.Sprintf("%s:%d", host, port)
	return s.router.Run(addr)
}
