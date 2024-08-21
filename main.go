package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)


type Todo struct {
	ID        int `json:"id"`
	Completed bool `json:"completed"`
	Body      string `json:"body"`
}

func main(){
	fmt.Println("hello world!!")

	app := fiber.New()


	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(Todo{})
		fmt.Println("requestが来た")
		return c.Status(202).JSON(fiber.Map{"msg": "Hello, World"})
	})

	app.Post("/todos", func(c *fiber.Ctx) error {

		todo := &Todo{}


		if err := c.BodyParser(todo); err != nil {
			fmt.Println(err)
			return err
		}

		if todo.Body == ""{
			return c.Status(400).JSON(fiber.Map{"error": "Body is empty"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})



	
	log.Fatal(app.Listen(":4000"))

}