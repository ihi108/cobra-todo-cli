/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	appDefs "github.com/ihi108/cobra-todo-cli/types"
	utils "github.com/ihi108/cobra-todo-cli/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the tasks being tracked",
	Long: `Lists a tracked tasks based on arg. For example:

	task-cli list

Lists all the tasks

	task-cli list done

Lists all done tasks

	task-cli list todo

Lists all tasks to be done

	task-cli list in-progress

Lists all tasks in-progress
`,
	Run: func(cmd *cobra.Command, args []string) {
		var tasks appDefs.Tasks

		bytes := utils.ReadFile(appDefs.JsonFile)
		utils.UnmarshalJSON(bytes, &tasks)
		if len(args) == 0 {
			utils.StatusList(tasks, "")
		} else {
			status := args[0]
			switch status {
			case "todo":
				utils.StatusList(tasks, "todo")
			case "done":
				utils.StatusList(tasks, "done")
			case "in-progress":
				utils.StatusList(tasks, "in-progress")
			default:
				str := "Usage:\n  todo-cli list\n  todo-cli list [done | todo | in-progress]"
				fmt.Println(str)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
