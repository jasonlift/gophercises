package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/jasonlift/gophercises/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Lists a",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		id, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong: ", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Task (%d)%s is created.\n", id, task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
