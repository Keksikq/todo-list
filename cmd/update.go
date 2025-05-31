package cmd

import (
	"strconv"

	"todo-list-for-levus/internal/storage"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update an existing task",
	Long:  `Update an existing task's title, description, priority, due date, or status.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		task, err := storage.Global.GetTaskByID(id)
		if err != nil {
			return err
		}

		if title, _ := cmd.Flags().GetString("title"); title != "" {
			task.Title = title
		}
		if description, _ := cmd.Flags().GetString("description"); description != "" {
			task.Description = description
		}
		if priority, _ := cmd.Flags().GetInt("priority"); priority != 0 {
			task.Priority = priority
		}
		if dueDate, _ := cmd.Flags().GetString("due"); dueDate != "" {
			task.DueDate = dueDate
		}
		if status, _ := cmd.Flags().GetString("status"); status != "" {
			task.Status = status
		}

		if err := task.Validate(); err != nil {
			return err
		}

		return storage.Global.UpdateTask(id, task)
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("title", "t", "", "New task title")
	updateCmd.Flags().StringP("description", "d", "", "New task description")
	updateCmd.Flags().IntP("priority", "p", 0, "New task priority (1-5)")
	updateCmd.Flags().StringP("due", "u", "", "New due date (YYYY-MM-DD)")
	updateCmd.Flags().StringP("status", "s", "", "New status (pending|done)")
}
