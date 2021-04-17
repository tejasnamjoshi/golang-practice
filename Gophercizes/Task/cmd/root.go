package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "A command line task manager.",
	Long:  "Task is a command line utility tool for managing your tasks.",
}
