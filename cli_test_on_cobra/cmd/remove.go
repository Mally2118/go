package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Удаляет задачу из списка",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Ошибка: введите корректные данные")
			return
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Ошибка: номер задачи должен быть числом")
			return
		}

		tasks := loadTasks()
		newTasks := []Task{}
		found := false

		for _, task := range tasks {
			if task.ID == id {
				found = true
				continue
			}
			newTasks = append(newTasks, task)
		}

		if found {
			saveTasks(newTasks)
			fmt.Printf("Задача #%d удалена!\n", id)
		} else {
			fmt.Println("Задача не найдена")
		}

	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
