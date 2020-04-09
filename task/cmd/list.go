package cmd

import (
	"fmt"
	"os"

	"github.com/jasonlift/gophercises/task/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the task items",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ReadAllTasks()
		if err != nil {
			fmt.Println("Something went wrong: ", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete! Why not take a vocation? 🤗")
			return
		}
		fmt.Println("You have the following tasks:")
		for _, task := range tasks {
			fmt.Printf("%d. %s\n", task.Key, task.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
