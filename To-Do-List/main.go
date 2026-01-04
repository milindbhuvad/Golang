package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const todoFile = "todo.json"

type Task struct {
	ID          int
	Description string
	Completed   bool
}

func loadTasks() []Task {
	file, err := os.ReadFile(todoFile)
	if err != nil {
		return []Task{}
	}
	var todoList []Task
	_ = json.Unmarshal(file, &todoList)
	return todoList
}

func saveTasks(todoList []Task) {
	data, _ := json.MarshalIndent(todoList, "", "  ")
	_ = os.WriteFile(todoFile, data, 0644)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo [add|list|done|delete]")
		return
	}

	command := os.Args[1]
	todos := loadTasks()

	switch command {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo add \"task name\"")
			return
		}
		todos = append(todos, Task{Description: os.Args[2]})
		saveTasks(todos)
		fmt.Println("Task added ✔")

	case "list":
		if len(todos) == 0 {
			fmt.Println("No tasks yet 🎉")
			return
		}
		for i, t := range todos {
			status := " "
			if t.Completed {
				status = "✔"
			}
			fmt.Printf("%d. [%s] %s\n", i+1, status, t.Description)
		}

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo done <task number>")
			return
		}
		index, _ := strconv.Atoi(os.Args[2])
		todos[index-1].Completed = true
		saveTasks(todos)
		fmt.Println("Task completed ✔")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo delete <task number>")
			return
		}
		index, _ := strconv.Atoi(os.Args[2])
		todos = append(todos[:index-1], todos[index:]...)
		saveTasks(todos)
		fmt.Println("Task deleted 🗑")

	default:
		fmt.Println("Unknown command")
	}
}
