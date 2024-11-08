// functions useful to all the commands
package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	appDefs "github.com/ihi108/cobra-todo-cli/types"
)

// processNewTasks - get all the args, transforms them to Tasks
// and collect them in a slice of Tasks
// id is the count of already existing tasks
func ProcessNewTasks(id int, newTasks *appDefs.Tasks, args []string) {
	timeString := time.Now().Format(time.RFC3339)
	for _, task := range args {
		item := appDefs.Task{
			Id:          id,
			Description: task,
			Status:      "todo",
			CreatedAt:   timeString,
			UpdatedAt:   timeString,
		}
		*newTasks = append(*newTasks, item)
		id++
	}
}

// MarshalJSON - converts a list of tasks to JSON
// returns the JSON bytes
func MarshalJSON(tasks appDefs.Tasks) []byte {
	bytes, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

// WriteFile - writes the given bytes to a specified file
// with permission 0640
func WriteFile(file string, bytes []byte) {
	err := os.WriteFile(file, bytes, 0640)
	if err != nil {
		log.Fatal(err)
	}
}

// ReadFile - reads a file and returns the read bytes
func ReadFile(file string) []byte {
	bytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

// UnmarshalJSON - unmarshals JSON Tasks
func UnmarshalJSON(bytes []byte, tasks *appDefs.Tasks) {
	err := json.Unmarshal(bytes, tasks)
	if err != nil {
		log.Fatal(err)
	}
}

// UpdateStatus - updates the status of a task as done or in-progress
func UpdateStatus(tasks appDefs.Tasks, status string, id int) appDefs.Tasks {
	var found bool
	var newTasks appDefs.Tasks

	newTasks = make(appDefs.Tasks, 0)

	found = false
	for _, task := range tasks {
		if task.Id == id {
			found = true
			task.Status = status
		}
		newTasks = append(newTasks, task)
	}

	if found == false {
		str := fmt.Sprintf("Task with ID: %v, Not Found", id)
		fmt.Println(str)
		os.Exit(1)
	}

	return newTasks
}
