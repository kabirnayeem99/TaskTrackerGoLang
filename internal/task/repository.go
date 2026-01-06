package task

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func AddTask(description string) {
	if len(description) < 1 {
		log.Println("Please, provide a proper task description.")
		return
	}

	t := time.Now()

	task := Task{
		Description: description,
		Status:      ToDo,
		CreatedAt:   t,
		UpdatedAt:   t,
	}

	taskID, err := WriteToStorage(task)

	if err != nil {
		log.Println("Failed to add task ", err)
		return
	}

	if taskID > 0 {
		log.Printf("Task added successfully (ID: %v)\n", taskID)
	}
}

func UpdateTask(id int, description string) {
	t, err := FindTaskById(id)

	if err != nil {
		log.Fatalf("Failed to mark task as done, because, %v", err)
	}

	t.Description = description
	t.UpdatedAt = time.Now()
	WriteToStorage(t)
}

func ListTask() {
	tasks := ReadTasks()

	fmt.Printf(
		"%-5s %-15s %-30s %-20s %-20s\n",
		"ID", "STATUS", "DESCRIPTION", "CREATED AT", "UPDATED AT",
	)
	fmt.Println(strings.Repeat("-", 95))

	for _, t := range tasks {
		fmt.Printf(
			"%-5d %-15s %-30s %-20s %-20s\n",
			t.ID,
			Truncate(string(t.Status), 13),
			Truncate(t.Description, 30),
			t.CreatedAt.Format("2006-01-02 15:04"),
			t.UpdatedAt.Format("2006-01-02 15:04"),
		)
	}
}

func UpdateTaskStatus(id int, status TaskStatus) {
	t, err := FindTaskById(id)

	if err != nil {
		log.Fatalf("Failed to mark task as done, because, %v", err)
	}

	t.Status = status
	t.UpdatedAt = time.Now()
	WriteToStorage(t)
}

func MarkTaskDone(id int) {
	UpdateTaskStatus(id, Done)
}

func MarkTaskInProgress(id int) {
	UpdateTaskStatus(id, InProgress)
}

func DeleteTask(id int) {
	t, err := FindTaskById(id)

	if err != nil {
		log.Fatalf("Failed to mark task as done, because, %v", err)
	}
	DeleteFromStorage(t)
}
