package cmd

import (
	"strconv"

	"todo-list-for-levus/internal/storage"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a task",
	Long:  `Delete a task by its ID.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		return storage.Global.DeleteTask(id)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
