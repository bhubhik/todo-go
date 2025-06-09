package main

import "fmt"

type Todo struct {
	id        int
	title     string
	completed bool
}

func main() {
	todos := []Todo{
		{id: 1, title: "Make Todo App", completed: false},
		{id: 2, title: "Learn Go", completed: false},
	}

	fmt.Println("Todo List:")
	for _, t := range todos {
		status := " "
		if t.completed {
			status = "x"
		}
		fmt.Printf("[%s] %d: %s\n", status, t.id, t.title)
	}
}
