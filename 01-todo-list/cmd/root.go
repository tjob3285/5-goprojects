/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todolist",
	Short: "Add, Create, Complete, Delete, and List your ToDo List",
	Long: `A CLI application that allows a user to add, complete, delete, and list their
	tasks.
	
	The list will be saved as a common csv for now, maybe JSON or SQLite in the future.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todolist.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	var err error
	todos, err = LoadTodos()
	if err == nil {
		nextID = len(todos)
	} else {
		nextID = 0
	}
}

type Todo struct {
	ID          int
	Description string
	CreatedAt   time.Time
	IsCompleted bool
}

const csvFile = "todolist.csv"

// Load from csv
func LoadTodos() ([]Todo, error) {
	file, err := os.Open(csvFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var todos []Todo
	for _, record := range records[1:] {
		id, _ := strconv.Atoi(record[0])
		description := record[1]
		createdAt, _ := time.Parse(time.RFC3339, record[2])
		isCompleted, _ := strconv.ParseBool(record[3])

		todos = append(todos, Todo{ID: id, Description: description, CreatedAt: createdAt, IsCompleted: isCompleted})
	}
	return todos, nil
}

// SaveTodos saves todos to the CSV file
func SaveTodos(todos []Todo) error {
	file, err := os.Create(csvFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Description", "CreatedAt", "IsCompleted"})
	for _, todo := range todos {
		writer.Write([]string{
			strconv.Itoa(todo.ID),
			todo.Description,
			todo.CreatedAt.Format(time.RFC3339),
			strconv.FormatBool(todo.IsCompleted),
		})
	}
	return nil
}
