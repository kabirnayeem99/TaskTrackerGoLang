package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
)

const sfName = "tasks.json"

func storageFileExists() (bool, error) {
	_, err := os.Stat(sfName)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, err
	}

	return false, err
}

func createFileIfNotExists() (bool, error) {
	ex, err := storageFileExists()

	if ex {
		return false, nil
	}

	f, err := os.Create(sfName)

	if err == nil {
		defer f.Close()
		return true, nil
	}

	log.Println("Failed to create file", sfName, err)

	return false, err
}

func ReadTasks() []Task {
	fb, err := os.ReadFile(sfName)
	var tasks []Task

	err = json.Unmarshal(fb, &tasks)

	if err != nil {
		return []Task{}
	}

	return tasks
}

func FindTaskById(id int) (Task, error) {
	ts := ReadTasks()

	if len(ts) == 0 {
		return Task{}, errors.New("Tasks list is empty.")
	}

	for _, t := range ts {
		if t.ID == id {
			return t, nil
		}
	}

	return Task{}, errors.New("Task not found.")
}

func DeleteFromStorage(task Task) (int, error) {
	ts := ReadTasks()

	deleteIndex := -1
	for i, t := range ts {
		if t.ID == task.ID {
			deleteIndex = i
			break
		}
	}

	if deleteIndex == -1 {
		return 0, errors.New("task with ID" + fmt.Sprint(task.ID) + " not found")
	}

	ts = append(ts[:deleteIndex], ts[deleteIndex+1:]...)

	if _, err := WriteTasks(ts); err != nil {
		return 0, err
	}

	return task.ID, nil
}

func WriteToStorage(task Task) (int, error) {
	up := task.ID > 0

	if !up {
		_, err := createFileIfNotExists()

		if err != nil {
			return task.ID, err
		}
	}

	ts := ReadTasks()

	if len(ts) > 100 && !up {
		log.Println("Resetting tasks list, you can't add more than 100 tasks.")
		ts = []Task{}
	}

	if !up {
		if len(ts) > 0 {
			task.ID = ts[len(ts)-1].ID + 1
		} else {
			task.ID = 1
		}
		ts = append(ts, task)
	}

	if up {
		for i, t := range ts {
			if t.ID == task.ID {
				ts[i] = task
			}
		}
	}

	if _, err := WriteTasks(ts); err != nil {
		return task.ID, err
	}

	return task.ID, nil
}

func WriteTasks(ts []Task) (bool, error) {
	sort.Slice(ts, func(i, j int) bool {
		return ts[i].ID < ts[j].ID
	})

	tjs, err := json.Marshal(ts)
	if err != nil {
		log.Println(err)
		return false, err
	}

	f, err := os.OpenFile(sfName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println("Failed to open file", err)
		return false, err
	}

	defer f.Close()

	if _, err := f.WriteString(string(tjs)); err != nil {
		log.Println("Failed to write to file", err)
		return false, err
	}

	return true, nil
}
