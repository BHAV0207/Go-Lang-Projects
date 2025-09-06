package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Description string `json:"description"`
	IsComplete  bool   `json:"isComplete"`
}

func (t *Task) markComplete() {
	t.IsComplete = true
}

func addTask(task []Task, description string) []Task {
	newStruct := Task{Description: description, IsComplete: false}
	return append(task, newStruct)

}

func deleteTask(task []Task, index int) []Task {
	return append(task[:index], task[index+1:]...)
}

func printTasks(tasks []Task) {
	for i, t := range tasks {
		if t.IsComplete {
			fmt.Printf("%d. [x] %s\n", i, t.Description)
		} else {
			fmt.Printf("%d. [ ] %s\n", i, t.Description)
		}
	}
}

func main() {
	arr := []Task{}

	fmt.Println("the current data")
	fileData, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("No previous tasks found, starting fresh")
	} else {
		err = json.Unmarshal(fileData, &arr)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
		}
	}

	printTasks(arr)

	for {
		var a string
		fmt.Scanln(&a)
		if strings.ToLower(a) == "add" {
			var b string
			fmt.Scanln(&b)
			arr = addTask(arr, b)
		} else if strings.ToLower(a) == "delete" {
			var b int
			fmt.Scanln(&b)
			if b >= len(arr) {
				fmt.Println("cannot delete element which does not exist")
			} else {
				arr = deleteTask(arr, b)
			}
		} else if strings.ToLower(a) == "complete" {
			var b int
			fmt.Scanln(&b)
			arr[b].markComplete()

		} else if strings.ToLower(a) == "quit" {
			break
		}
	}

	printTasks(arr)

	data, err := json.MarshalIndent(arr, "", " ")
	if err != nil {
		fmt.Println("Error Marshelling", err)
		return
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error saving", err)
		return
	}

	fmt.Println("Tasks saved to tasks.json")

}
