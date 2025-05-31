package cmd

import (
	"os"
	"path/filepath"

	"todo-list-for-levus/internal/storage"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A CLI application for managing your tasks",
	Long: `Welcome to Todo CLI!

This is a command line application for managing your tasks. Here are some examples:

  todo add "Learn Go" -d "Complete the Go tutorial" -p 4 -t "2024-12-31"
  todo list
  todo list -s pending -o priority
  todo update 1 -t "Master Go" -p 5 -s done
  todo delete 1

Available Commands:
  add    - Add a new task
  list   - List all tasks
  update - Update an existing task
  delete - Delete a task
  help   - Show this help message

Type 'help' for more information about a command.`,
}

func init() {
	initializeStorage()
}

func initializeStorage() {
	var err error
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	storagePath := filepath.Join(homeDir, ".todo", "tasks.json")
	storage.Global, err = storage.NewFileStorage(storagePath)
	if err != nil {
		panic(err)
	}
}
