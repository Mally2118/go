package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Отмечает выполеннные задачи",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Ошибка: вы не ввели ID задачи")
			return
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Ошибка: номер задачи должен быть числом")
			return
		}

		tasks := loadTasks() // from add
		for i, task := range tasks {
			if task.ID == id {
				tasks[i].Complete = true
				saveTasks(tasks)
				fmt.Printf("Задача #%d выполнена!\n", id)
				return
			}
		}
		fmt.Println("Задача не найдена")
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
