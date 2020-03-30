package main

import "fmt"
import "github.com/gofiber/fiber"

func main() {
  app := fiber.New()

  app.Static("/public", "./public")

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

  app.Listen(3000)
}
