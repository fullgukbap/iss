package http

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Router struct {
	*fiber.App
}

func NewRouter(pictureHandler PictureHandler) (*Router, error) {
	app := fiber.New(fiber.Config{
		ErrorHandler: customErrorHandler,
	})
	app.Use(recover.New())
	app.Use(helmet.New())
	app.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Content-Length", "Accept-Language", "Accept-Encoding", "Connection", "Access-Control-Allow-Origin"},
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"},
	}))

	app.Use(limiter.New(limiter.Config{
		Max:               20,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	// Or extend your config for customization
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	app.Use(etag.New())

	app.Post("/api/picutre", pictureHandler.Create)
	app.Get("/api/picutre", pictureHandler.Find)
	app.Put("/api/picutre", pictureHandler.Update)
	app.Delete("/api/picutre", pictureHandler.Delete)

	var r *Router
	r.App = app
	return r, nil
}
