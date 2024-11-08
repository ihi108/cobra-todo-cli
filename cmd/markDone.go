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

// markDoneCmd represents the markDone command
var markDoneCmd = &cobra.Command{
	Use:   "mark-done",
	Short: "Marks a task as done",
	Long: `Marks a task with a given ID as done.For example:
	
	task-cli mark-done 1
`,
	Run: func(cmd *cobra.Command, args []string) {
		var tasks, newTasks appDefs.Tasks
		var found bool

		if len(args) < 1 {
			str := fmt.Sprintf("Usage:\n  todo-cli mark-done [taskID]")
			fmt.Println(str)
			os.Exit(1)
		}

		bytes := utils.ReadFile(appDefs.JsonFile)
		utils.UnmarshalJSON(bytes, &tasks)

		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		found = false
		for _, task := range tasks {
			if task.Id == id {
				found = true
				task.Status = "done"
			}
			newTasks = append(newTasks, task)
		}

		if found == false {
			str := fmt.Sprintf("Task with ID: %v, Not Found", id)
			fmt.Println(str)
			os.Exit(1)
		}

		bytes = utils.MarshalJSON(newTasks)
		utils.WriteFile(appDefs.JsonFile, bytes)

		fmt.Printf("Task with ID: %v, Marked as done\n", id)

	},
}

func init() {
	rootCmd.AddCommand(markDoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markDoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markDoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
