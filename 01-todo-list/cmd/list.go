/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		listItems(all)
	},
}

var all bool

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "Show all items, including incomplete ones")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listItems(showAll bool) {
	for _, todo := range todos {
		status := " "
		if todo.IsCompleted {
			status = "✓"
		}

		if showAll || !todo.IsCompleted {
			fmt.Printf("%d: %s [%s] - Created at %s\n", todo.ID, todo.Description, status, todo.CreatedAt.Format(time.RFC3339))
		} else if showAll && todo.IsCompleted {
			fmt.Printf("%d: %s [%s] - Created at %s\n", todo.ID, todo.Description, status, todo.CreatedAt.Format(time.RFC3339))
		}
	}
}
