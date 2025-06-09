package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	id        int
	title     string
	completed bool
}

func main() {
	todos := []Todo{
		{id: 1, title: "Make Todo App", completed: false},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nTodo List:")
		for _, t := range todos {
			status := " "
			if t.completed {
				status = "x"
			}
			fmt.Printf("[%s] %d: %s\n", status, t.id, t.title)
		}

		fmt.Print("\nAdd a new todo (or type 'exit' to quit): ")
		scanner.Scan()
		text := strings.TrimSpace(scanner.Text())

		if text == "exit" {
			fmt.Println("Bye!")
			break
		}

		newTodo := Todo{
			id:        len(todos) + 1,
			title:     text,
			completed: false,
		}

		todos = append(todos, newTodo)
	}
}
