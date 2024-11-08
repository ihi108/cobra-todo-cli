/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	appDefs "github.com/ihi108/cobra-todo-cli/types"
	utils "github.com/ihi108/cobra-todo-cli/utils"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a task",
	Long: `Updates the contents of a task with ID. For example:

	task-cli udpate 1 "Buy groceries and cook dinner"
`,
	Run: func(cmd *cobra.Command, args []string) {
		var tasks appDefs.Tasks
		var newTasks appDefs.Tasks
		var found bool

		tasks = make(appDefs.Tasks, 0)
		newTasks = make(appDefs.Tasks, 0)
		found = false

		if len(args) < 2 {
			str := fmt.Sprintf("Usage:\n  todo-cli update [taskID] [taskDescription]")
			fmt.Println(str)
			os.Exit(1)
		}

		id, task := args[0], args[1]

		bytes := utils.ReadFile(appDefs.JsonFile)
		utils.UnmarshalJSON(bytes, &tasks)
		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
		}

		// update task
		for _, item := range tasks {
			if item.Id == idInt {
				item.Description = task
				found = true

				timeStr := time.Now().Format(time.RFC3339)
				item.UpdatedAt = timeStr
			}
			newTasks = append(newTasks, item)
		}
		if found == false {
			msg := fmt.Sprintf("Task with ID: %v, Not Found\n", id)
			log.Fatal(msg)
		}

		bytes = utils.MarshalJSON(newTasks)
		utils.WriteFile(appDefs.JsonFile, bytes)
		tasks.UpdateOutput(idInt)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
