package main

import "fmt"

type Task struct {
	description string
	isComplete  bool
}

func (t *Task) markComplete() {
	t.isComplete = true
}

func addTask(task []Task, description string) []Task {
	newStruct := Task{description: description, isComplete: false}
	return append(task, newStruct)

}

func deleteTask(task []Task, index int) []Task {
	return append(task[:index], task[index+1:]...)
}

func main() {
	arr := []Task{}

	arr = addTask(arr, "Buy milk")
	arr = addTask(arr, "Call mom")
	arr[0].markComplete()
	arr = deleteTask(arr, 1)

	fmt.Println(arr)
}
