package main

import (
	"bufio"
	"encoding/json"
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
	const filename = "todos.json"

	todos, err := loadTodos(filename)
	if err != nil {
		fmt.Println("Error loading todos:", err)
		return
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
			fmt.Print("\nType 'add' to add todo, 'delete' to delete or 'exit' to quit.")
			scanner.Scan()
			cmd := strings.TrimSpace(scanner.Text())

			switch cmd {
			case "add":
				fmt.Print("Enter your todo:")
				scanner.Scan()
				text = strings.TrimSpace(scanner.Text())
				if err := addTodo(text, &todos); err != nil {
					fmt.Println("Error adding todo:", err)
				}
				err = saveTodos(filename, todos)
				if err != nil {
					fmt.Println("Error saving todos:", err)
				}
			case "exit":
				fmt.Println("Bye!")
				return
			}
			break
		}
	}
}

func addTodo(text string, todos *[]Todo) error {
	if text == "" {
		return fmt.Errorf("Invalid. Empty Todo.")
	}

	newTodo := Todo{
		ID:        len(*todos) + 1,
		Title:     text,
		Completed: false,
	}

	*todos = append(*todos, newTodo)
	fmt.Println("Todo Added!")
	return nil
}

func loadTodos(filename string) ([]Todo, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil
		}
		return nil, err
	}

	var todos []Todo
	err = json.Unmarshal(data, &todos)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func saveTodos(filname string, todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(filname, data, 0644)
}
