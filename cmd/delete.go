/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	appDefs "github.com/ihi108/cobra-todo-cli/types"
	utils "github.com/ihi108/cobra-todo-cli/utils"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a task from the tracker",
	Long: `Deletes a task with a given ID from the tracker. For example:

	task-cli delete 1

`,
	Run: func(cmd *cobra.Command, args []string) {
		var tasks appDefs.Tasks
		var newTasks appDefs.Tasks
		var found bool

		if len(args) < 1 {
			str := fmt.Sprintf("Usage:\n  todo-cli delete [taskID]")
			fmt.Println(str)
			os.Exit(1)
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		tasks = make(appDefs.Tasks, 0)
		newTasks = make(appDefs.Tasks, 0)
		bytes := utils.ReadFile(appDefs.JsonFile)
		utils.UnmarshalJSON(bytes, &tasks)
		found = false
		for _, task := range tasks {
			if task.Id == id {
				found = true
			} else {
				newTasks = append(newTasks, task)
			}
		}

		if found == false {
			msg := fmt.Sprintf("Task with ID: %v, Not Found\n", id)
			log.Fatal(msg)
		}

		bytes = utils.MarshalJSON(newTasks)
		utils.WriteFile(appDefs.JsonFile, bytes)
		fmt.Printf("Task with ID: %v, deleted successfully\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
