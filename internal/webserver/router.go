package webserver

import "github.com/gofiber/fiber/v2"

// create http router
func (ws *WebServer) router() {
	ws.server.Get("retailer/:id", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!" + "id: " + c.Params("id"))
	})

	return
}
