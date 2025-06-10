package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bhubhik/todo-go/utils"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

func main() {
	const filename = "todos.json"

	todos, err := utils.LoadTodos(filename)
	if err != nil {
		fmt.Println("Error loading todos:", err)
		return
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		utils.PrintTodos(todos)
		var text string

		for {
			fmt.Print("\nType 'add' to add todo, 'delete' to delete or 'exit' to quit.")
			scanner.Scan()
			cmd := strings.TrimSpace(scanner.Text())

			switch cmd {
			case "add":
				fmt.Print("Enter your todo:")
				scanner.Scan()
				text = strings.TrimSpace(scanner.Text())
				if err := utils.AddTodo(text, &todos); err != nil {
					fmt.Println("Error adding todo:", err)
				}
				err = utils.SaveTodos(filename, todos)
				if err != nil {
					fmt.Println("Error saving todos:", err)
				}
			case "delete":
				fmt.Print("Enter id of the todo to delete: ")
				scanner.Scan()
				input := strings.TrimSpace(scanner.Text())

				id, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("Invalid input please enter a number.")
					continue
				}

				err = utils.DeleteTodos(&todos, id)
				if err != nil {
					fmt.Println("Could not delete: ", err)
				} else {
					utils.SaveTodos(filename, todos)
				}

			case "exit":
				fmt.Println("Bye!")
				return
			}
			break
		}
	}
}
