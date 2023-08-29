package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// creating struct to house our todos
type Todo struct {
	ID     int    `json:"id"`
	Word   string `json:"word"`
	Done   bool   `json:"done"`
	Answer string `json:"answer"`
}

//we are just storing this in memory

func main() {
	fmt.Println("hello world")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	todos := []Todo{}
	//now can append slice and return it

	//health check endpoint
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("Server Ok")
	})
	// * create todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		err := c.BodyParser(todo)
		if err != nil {
			return err
		}
		todo.ID = len(todos) + 1 //get len of todos and plus one

		todos = append(todos, *todo)
		return c.JSON(todos)
	})
	//* updated todos
	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
		//get id out of url
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}

		for i, t := range todos {
			if t.ID == id {
				todos[i].Done = true
				break
			}
		}
		return c.JSON(todos)
	})
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})
	//* get all todos
	log.Fatal(app.Listen(":4000")) //if app.listen throws error, it will say it here

}
