package main

import "fmt"
import "github.com/gofiber/fiber"

func main() {
  app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) {
    c.Send("Hello world from fiber")
  })

  app.Get("/:name", func(c *fiber.Ctx) {
    fmt.Printf("Hello %s!\n", c.Params("name"))
  })

  app.Listen(3000)
}
