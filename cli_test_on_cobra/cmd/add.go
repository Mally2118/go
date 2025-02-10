package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"encoding/json"
)

const fileName = "tasks.json"

type Task struct {
	ID		 int 	`json:"id"`
	Name	 string `json:"name"`
	Complete bool 	`json:complete`
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Добавить новую задачу",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Ошибка: введите корректные данные")
		}
		taskName := args[0]
		tasks := loadTasks()

		newTask := Task{
			ID:   len(tasks) + 1,
			Name: taskName,
		}
		tasks = append(tasks, newTask)
		saveTasks(tasks)
		fmt.Printf("Добавлена задача #%d: %s\n", newTask.ID, newTask.Name)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func loadTasks() []Task {
	var tasks []Task
	file, err := os.ReadFile(fileName)
	if err == nil {
		json.Unmarshal(file, &tasks)
	}
	return tasks
}

func saveTasks(tasks []Task) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(fileName, data, 0644)
}