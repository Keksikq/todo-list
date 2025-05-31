package cmd

import (
	"strings"

	"todo-list-for-levus/internal/model"
	"todo-list-for-levus/internal/storage"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add \"[title]\"",
	Short: "Add a new task",
	Long:  `Add a new task with optional description, priority, and due date.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		title := strings.Join(args, " ")
		description, _ := cmd.Flags().GetString("description")
		priority, _ := cmd.Flags().GetInt("priority")
		dueDate, _ := cmd.Flags().GetString("due")

		task := model.NewTask(title, description, priority, dueDate)
		if err := task.Validate(); err != nil {
			return err
		}

		return storage.Global.AddTask(task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("description", "d", "", "Task description")
	addCmd.Flags().IntP("priority", "p", 3, "Task priority (1-5)")
	addCmd.Flags().StringP("due", "t", "", "Due date (YYYY-MM-DD)")
}
