package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

func main() {
	todos := []Todo{
		{ID: 1, Title: "Make Todo App", Completed: false},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nTodo List:")
		for _, t := range todos {
			status := " "
			if t.Completed {
				status = "x"
			}
			fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Title)
		}

		var text string

		for {
			fmt.Print("Please add your todo or type exit to quit:")
			scanner.Scan()
			text = strings.TrimSpace(scanner.Text())

			switch text {
			case "exit":
				fmt.Println("Bye!")
				return
			case "":
				fmt.Println("InvalID Input, todo cannot be empty. Please try again.")
				continue
			}
			break
		}

		newTodo := Todo{
			ID:        len(todos) + 1,
			Title:     text,
			Completed: false,
		}

		todos = append(todos, newTodo)
	}
}
