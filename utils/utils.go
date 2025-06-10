package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

func AddTodo(text string, todos *[]Todo) error {
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

func DeleteTodos(todos *[]Todo, id int) error {
	newTodos := []Todo{}
	found := false

	for _, t := range *todos {
		if t.ID != id {
			newTodos = append(newTodos, t)
		} else {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("Todo with ID %d not found.", id)
	}

	*todos = newTodos
	fmt.Printf("Todo with ID %d deleted.", id)
	return nil
}

func LoadTodos(filename string) ([]Todo, error) {
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

func PrintTodos(todos []Todo) {
	fmt.Println("\nTodo List:")
	for _, t := range todos {
		status := " "
		if t.Completed {
			status = "x"
		}
		fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Title)
	}
}

func SaveTodos(filname string, todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(filname, data, 0644)
}
