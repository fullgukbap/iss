package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Router struct {
	*fiber.App
}

func NewRouter(pictureHandler *PictureHandler) (*Router, error) {
	app := fiber.New(fiber.Config{
		ErrorHandler: customErrorHandler,
	})
	app.Use(recover.New())
	app.Use(helmet.New())
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins: "*",
		// AllowCredentials: true,
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// app.Use(limiter.New(limiter.Config{
	// 	Max:               20,
	// 	Expiration:        30 * time.Second,
	// 	LimiterMiddleware: limiter.SlidingWindow{},
	// }))
	// Or extend your config for customization
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	app.Use(etag.New())

	// Initialize default config
	app.Use(logger.New())

	// Or extend your config for customization
	// Logging remote IP and Port
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	api := app.Group("/v1/api")

	api.Post("/picture", pictureHandler.Create)
	api.Get("/picture/:id", pictureHandler.Find)
	api.Put("/picture/:id", pictureHandler.Update)
	api.Delete("/picture/:id", pictureHandler.Delete)

	return &Router{App: app}, nil
}

func (r *Router) Run(port string) error {
	return r.Listen(port)
}
