package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID:   "1269435",
		Key:     "fbd694ff3462c16966c2",
		Secret:  "25be050b20ccc45a5455",
		Cluster: "ap2",
		Secure:  true,
	}

	data := map[string]string{"message": "hello world"}
	pusherClient.Trigger("my-channel", "my-event", data)

	// app.Post("/api/messages", func(c *fiber.Ctx) error {
	// 	var data map[string]string

	// 	if err := c.BodyParser(&data); err != nil {
	// 		return err
	// 	}

	// 	pusherClient.Trigger("ChatApp-by-Seema", "message", data)
	// 	return c.JSON([]string{})
	// })

	app.Listen(":8000")
}
