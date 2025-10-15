package main

import "fmt"

func main() {
	task1 := taskInfo{1, "Clean the house", false}
	task2 := taskInfo{2, "Buy milk", false}
	tasks := []taskInfo{task1, task2}
	fmt.Println(tasks)
}

type taskInfo struct {
	taskID int
	taskDescription string
	statusDone bool
}


