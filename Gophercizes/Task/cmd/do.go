package cmd

import (
	"fmt"
	"golang_practice/Gophercizes/Task/db"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as done",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Cannot parse %s \n", arg)
			} else {
				ids = append(ids, id)
			}
		}

		// This can be done in a better way. Skipping for now.
		for _, id := range ids {
			err := db.RemoveTask(id)
			fmt.Println("Task marked as done.")
			if err != nil {
				fmt.Printf("\nFailed to remove task with id : %d", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
