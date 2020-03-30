package main

import "fmt"
import "github.com/gofiber/fiber"

func main() {
  app := fiber.New()

  app.Static("/public", "./public")

  // Set custom headers to and endpoint
  app.Use("/", func(c *fiber.Ctx) {
    c.Set("X-Custom-Header", "random-header")
    c.Next()
  })

  app.Use(func (c *fiber.Ctx){
    fmt.Println("My first middleware")
    c.Next()
  })

  app.Get("/", func(c *fiber.Ctx) {
    c.Send("Hello world from fiber")
  })

  app.Get("/:name", func(c *fiber.Ctx) {
    fmt.Printf("Hello %s!\n", c.Params("name"))
  })

  app.Get("/:name/id?", func(c *fiber.Ctx) {
    fmt.Printf("Hello %s! Id: %s\n", c.Params("name"), c.Params("id"))
  })

  api := app.Group("/api")

  v1 := api.Group("/v1")
  v1.Get("/list", func(c *fiber.Ctx) {
    c.Send("Response from /api/v1/list")
  })

  app.Listen(3000)
}
