package cmd

import (
	"fmt"
	"golang_practice/Gophercizes/Task/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.GetTasks()
		if err != nil {
			panic(err)
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks.")
			return
		}

		fmt.Println("You have the following tasks.")
		for i, task := range tasks {
			fmt.Printf("\n %d. %s", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
