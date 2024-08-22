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
		return c.Status(202).JSON(todos)
	})

	app.Post("/todos", func(c *fiber.Ctx) error {

		// 忘却録
		// variable
		x  := 1;
		// pointer xのアドレス
		p  := &x
		// pointerの中身
		fmt.Println(*p)

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

	app.Patch("/todos/:id",func(c *fiber.Ctx) error{
		id := c.Params("id")
		fmt.Println(id)

		for i,todo := range todos{
			if  fmt.Sprint(todo.ID) == id{
				todos[i].Completed = !todos[i].Completed
				return c.Status(200).JSON(todos[i])
			}
		}

		fmt.Println(todos)


		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	app.Delete("/todos/:id",func(c *fiber.Ctx)error{
		id := c.Params("id")
	
		for i,todo := range todos{
			if fmt.Sprint(todo.ID) == id {
				todos := append(todos[:i],todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"msg": "success"})
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})



	
	log.Fatal(app.Listen(":4000"))

}