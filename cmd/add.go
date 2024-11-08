/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"time"
	"encoding/json"
	"log"
	"os"
	"slices"

	appDefs "github.com/ihi108/cobra-todo-cli/types"
	"github.com/spf13/cobra"
)

// processNewTasks - get all the args, transforms them to Tasks
// and collect them in a slice of Tasks
// id is the count of already existing tasks
func processNewTasks(id int, newTasks *appDefs.Tasks, args []string) {
	timeString := time.Now().Format(time.RFC3339)
	for _, task := range args {
		item := appDefs.Task{
			Id: id, 
			Description: task,
			Status: "todo",
			CreatedAt: timeString,
			UpdatedAt: timeString,
		}
		*newTasks = append(*newTasks, item)
		id++
	}
}

// marshalJSON - converts a list of tasks to JSON
// returns the JSON bytes
func marshalJSON(tasks appDefs.Tasks) []byte {
	bytes, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

// writeFile - writes the given bytes to a specified file
// with permision 0640
func writeFile(file string, bytes []byte) {
	err := os.WriteFile(file, bytes, 0640)
	if err != nil {
		log.Fatal(err)
	}
}

// readFile - reads a file and returns the read bytes
func readFile(file string) []byte {
	bytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

// unmarshalJSON - unmarshals JSON tasks
func unmarshalJSON(bytes []byte, tasks *appDefs.Tasks) {
	err := json.Unmarshal(bytes, tasks)
	if err != nil {
		log.Fatal(err)
	}
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to track",
	Long: `Adds a task to the tracker, specifying the task name
For example:

	task-cli add "Buy groceries"

`,
	Run: func(cmd *cobra.Command, args []string) {
		var tasks appDefs.Tasks
		var newTasks appDefs.Tasks
		var bytes []byte

		tasks = make(appDefs.Tasks, 0)
		newTasks = make(appDefs.Tasks, 0)

		bytes, err := os.ReadFile(appDefs.JsonFile)
		if err != nil {
			// file is empty; no need for reading
			processNewTasks(0, &newTasks, args)
			bytes = marshalJSON(newTasks)
			writeFile(appDefs.JsonFile, bytes)
		} else {
			// read and process file
			bytes = readFile(appDefs.JsonFile)
			unmarshalJSON(bytes, &tasks)
			id := len(tasks)
			processNewTasks(id, &newTasks, args)
			tasks = slices.Concat(tasks, newTasks)
			bytes = marshalJSON(tasks)
			writeFile(appDefs.JsonFile, bytes)
		}

		newTasks.Output()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
