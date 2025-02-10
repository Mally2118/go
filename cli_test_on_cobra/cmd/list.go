package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Выводит список задач",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := loadTasks()
		if len(tasks) == 0 {
			fmt.Println("Список задач пуст")
			return
		}

		for _, task := range tasks {
			status := "Progress"
			if task.Complete {
				status = "Completed"
			}
			fmt.Printf("%d. %s [%s]\n", task.ID, task.Name, status)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
