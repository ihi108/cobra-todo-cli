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

// markInProgressCmd represents the markInProgress command
var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "Mark a task as in-progress",
	Long: `Marks a task with a given ID arg as in-progress.For example:

	task-cli mark-in-progress 1
`,
	Run: func(cmd *cobra.Command, args []string) {
		var tasks, newTasks appDefs.Tasks

		if len(args) < 1 {
			str := fmt.Sprintf("Usage:\n  todo-cli mark-in-progress [taskID]")
			fmt.Println(str)
			os.Exit(1)
		}

		bytes := utils.ReadFile(appDefs.JsonFile)
		utils.UnmarshalJSON(bytes, &tasks)

		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		newTasks = utils.UpdateStatus(tasks, "in-progress", id)
		bytes = utils.MarshalJSON(newTasks)
		utils.WriteFile(appDefs.JsonFile, bytes)
		fmt.Printf("Task with ID: %v, Marked as in-progress\n", id)
	},
}

func init() {
	rootCmd.AddCommand(markInProgressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markInProgressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markInProgressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
