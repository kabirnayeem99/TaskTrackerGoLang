package taskcli

import (
	"errors"
	"github.com/kabirnayeem99/TaskTrackerGoLang/internal/task"
	"log"
	"strconv"
)

func Run(args []string) error {
	ac := len(args)
	if ac < 2 {
		log.Fatalln("Usage: task-cli [method] [task]")
		return errors.New("Didn't pass proper arguements")
	}

	command := task.Command(args[1])

	switch command {
	case task.Add:
		if ac < 3 {
			log.Fatalln("Usage: task-cli add [task_description]")
			return errors.New("Didn't pass proper arguements")
		}
		description := args[2]
		task.AddTask(description)
	case task.Update:
		if ac < 4 {
			log.Fatalln("Usage: task-cli update [task_id] [task_new_description]")
			return errors.New("Didn't pass proper arguements")
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatalln("Usage: task-cli mark-done [task_id], task ID must be a number.")
		}
		description := args[3]
		task.UpdateTask(id, description)
	case task.List:
		task.ListTask()
	case task.MarkDone:
		if ac < 3 {
			log.Fatalln("Usage: task-cli [method] [task]")
			return errors.New("Didn't pass proper arguements")
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatalln("Usage: task-cli mark-done [task_id], task ID must be a number.")
		}
		task.MarkTaskDone(id)
	case task.MarkInProgress:
		if ac < 3 {
			log.Fatalln("Usage: task-cli mark-in-progress [task_id]")
			return errors.New("Didn't pass proper arguements")
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatalln("Usage: task-cli mark-in-progress [task_id], task ID must be a number.")
			return errors.New("Didn't pass proper arguements")
		}
		task.MarkTaskInProgress(id)

	case task.Delete:
	if ac < 3 {
			log.Fatalln("Usage: task-cli [method] [task]")
			return errors.New("Didn't pass proper arguements")
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatalln("Usage: task-cli mark-done [task_id], task ID must be a number.")
		}
		task.DeleteTask(id)	
	default:
		log.Println("Usage: task-cli [method] [task]")
	}

	return nil
}
