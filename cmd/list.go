package cmd

import (
	"fmt"
	"sort"
	"strings"

	"todo-list-for-levus/internal/model"
	"todo-list-for-levus/internal/storage"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all tasks with optional filtering by status and sorting by priority or date.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		status, _ := cmd.Flags().GetString("status")
		sortBy, _ := cmd.Flags().GetString("sort")

		tasks := storage.Global.GetTasks()

		if status != "" {
			filtered := make([]*model.Task, 0)
			for _, task := range tasks {
				if task.Status == status {
					filtered = append(filtered, task)
				}
			}
			tasks = filtered
		}

		switch strings.ToLower(sortBy) {
		case "priority":
			sort.Slice(tasks, func(i, j int) bool {
				return tasks[i].Priority > tasks[j].Priority
			})
		case "date":
			sort.Slice(tasks, func(i, j int) bool {
				return tasks[i].DueDate < tasks[j].DueDate
			})
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return nil
		}

		fmt.Printf("%-4s %-30s %-8s %-12s %-8s\n", "ID", "Title", "Priority", "Due Date", "Status")
		fmt.Println(strings.Repeat("-", 70))

		for _, task := range tasks {
			fmt.Printf("%-4d %-30s %-8d %-12s %-8s\n",
				task.ID,
				truncateString(task.Title, 28),
				task.Priority,
				task.DueDate,
				task.Status,
			)
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("status", "s", "", "Filter by status (pending|done)")
	listCmd.Flags().StringP("sort", "o", "", "Sort by (priority|date)")
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
