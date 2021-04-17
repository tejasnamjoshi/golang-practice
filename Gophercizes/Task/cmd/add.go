package cmd

import (
	"fmt"
	"golang_practice/Gophercizes/Task/db"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		id, err := db.CreateTask(task)
		if err != nil {
			panic(err)
		}

		fmt.Println("Task created with id : ", id) // Do not print id in the actual code. This is just for debugging.
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
