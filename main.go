package main

import (
	"fmt"
	"hokapi/app/data"
	"hokapi/app/middleware"
	"hokapi/app/route"
	"hokapi/scheduler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	f := fiber.New()
	go scheduler.Scheduler()

	json, err := data.InitJson()
	if err != nil {
		fmt.Println(err)
	}

	f.Use(cors.New())

	f.Use(middleware.RequestLimitMiddleware)

	route.InitRoute(f, json)
	f.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	f.Listen(":8080")
}
