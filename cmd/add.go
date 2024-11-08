/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"slices"

	appDefs "github.com/ihi108/cobra-todo-cli/types"
	utils "github.com/ihi108/cobra-todo-cli/utils"
	"github.com/spf13/cobra"
)

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
			utils.ProcessNewTasks(0, &newTasks, args)
			bytes = utils.MarshalJSON(newTasks)
			utils.WriteFile(appDefs.JsonFile, bytes)
		} else {
			// read and process file
			bytes = utils.ReadFile(appDefs.JsonFile)
			utils.UnmarshalJSON(bytes, &tasks)
			id := len(tasks)
			utils.ProcessNewTasks(id, &newTasks, args)
			tasks = slices.Concat(tasks, newTasks)
			bytes = utils.MarshalJSON(tasks)
			utils.WriteFile(appDefs.JsonFile, bytes)
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
