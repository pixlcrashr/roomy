package api

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/pixlcrashr/roomy/web"
	"gorm.io/gorm"
)

type Server struct {
	app *fiber.App
	db  *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	app := fiber.New(fiber.Config{
		AppName:     "Accounter API",
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	log.SetLevel(log.LevelTrace)

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/livez",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			sqlDB, err := db.DB()
			if err != nil {
				return false
			}

			if err := sqlDB.Ping(); err != nil {
				return false
			}

			return true
		},
		ReadinessEndpoint: "/readyz",
	}))

	web.RegisterRoutes(app)

	return &Server{
		app: app,
		db:  db,
	}
}

func (s *Server) Listen(addr string) error {
	return s.app.Listen(addr)
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
