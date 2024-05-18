package main

import (
	"fmt"
	"hokapi/app/data"
	"hokapi/app/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	f := fiber.New()

	
	json,err := data.InitJson()
    if err != nil {
        fmt.Println(err)
    }

	f.Use(cors.New())
    
    route.InitRoute(f,json)
	f.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	f.Listen(":8080")
}