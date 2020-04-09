package cmd

import (
	"fmt"
	"strconv"

	"github.com/jasonlift/gophercises/task/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks the specific task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument: ", arg)
			} else {
				ids = append(ids, id)
			}
		}
		for _, id := range ids {
			db.DeleteTask(id)
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
