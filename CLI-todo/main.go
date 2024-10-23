package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

const (
	Todo       = "todo"
	InProgress = "in-progress"
	Done       = "done"
)

var tasks []Task

/*
func findTaskByID(id string) (int, *Task) {
	for idx, task := range tasks {
		if task.ID == id {
			return idx, &task
		}
	}
	return -1, nil
}
*/

func loadData(filename string) ([]Task, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	return tasks, nil
}

func saveTasksToFile(filename string, tasks []Task) error {
	jsonData, err := json.MarshalIndent(tasks, "", "   ")
	if err != nil {
		return fmt.Errorf("error marshalling tasks to JSON: %w", err)
	}
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating a file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("error writing to a file: %w", err)
	}
	return nil
}

func main() {

	filename := "tasks.json"
	var err error
	tasks, err = loadData(filename)
	if err != nil {
		fmt.Println("Error loading tasks from a JSON file.")
		return
	}

	rootCmd := &cobra.Command{
		Use:   "root",
		Short: "CLI todo application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("You're running your CLI app!")
		},
	}

	deleteTasks := &cobra.Command{
		Use:   "delete",
		Short: "Provide an ID to delete a task.",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				fmt.Println("Please provide an ID to delete a task.")
				return
			}

			id := args[0]
			taskFound := false

			for idx, task := range tasks {
				if task.ID == id {
					tasks = append(tasks[:idx], tasks[idx+1:]...)
					taskFound = true
					fmt.Printf("Task with ID %s was deleted successfully.\n", id)
					break
				}
			}

			if taskFound {
				err := saveTasksToFile(filename, tasks)
				if err != nil {
					fmt.Println("Error writing to the JSON file.")
				} else {
					fmt.Println("Task was successfully deleted from the JSON file.")
				}
			} else {
				fmt.Println("Couldn't find the macthing ID")
			}
		},
	}

	AddTasks := &cobra.Command{
		Use:   "add",
		Short: "Add a task.",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) < 1 {
				fmt.Println("Error: Please provide a task description.")
				return
			}

			newTask := Task{ID: strconv.Itoa(len(tasks) + 1), Description: args[0], Status: "todo"}
			tasks = append(tasks, newTask)

			err := saveTasksToFile(filename, tasks)
			if err != nil {
				fmt.Println("Error writing to the JSON file.")
				return

			} else {
				fmt.Println("Task was successfully added to the JSON file.")
			}

			fmt.Println("Task added successfully!")

		},
	}

	ListTasks := &cobra.Command{
		Use:   "list",
		Short: "List all the tasks.",
		Run: func(cmd *cobra.Command, args []string) {
			tasks, err := loadData(filename)
			if err != nil {
				fmt.Println("Error loading tasks", err)
				return
			}

			if len(tasks) == 0 {
				fmt.Println("No tasks yet.")
				return
			}

			fmt.Println("Tasks:")
			for _, task := range tasks {
				fmt.Printf("ID: %s, Description: %s, Status: %s\n", task.ID, task.Description, task.Status)
			}
		},
	}

	updateTasks := &cobra.Command{
		Use:   "update",
		Short: "Provide an ID to update a task.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				fmt.Println("Invalid amount of parameters.")
				return
			}
			id := args[0]
			descr := args[1]
			taskFound := false

			for idx, task := range tasks {
				if task.ID == id {
					tasks[idx].Description = descr
					taskFound = true
					fmt.Println("Task updated successfully.")
					break
				}
			}

			if taskFound {
				err := saveTasksToFile(filename, tasks)
				if err != nil {
					fmt.Println("Couldn't apply changes to the JSON file")
				} else {
					fmt.Println("Changes made successfully!")
				}
			} else {
				fmt.Printf("A task with id %s doesn't exist.\n", id)
			}
		},
	}

	mark := &cobra.Command{
		Use:   "mark",
		Short: "Mark something as done or in-progress",
	}

	markDone := &cobra.Command{
		Use:   "done",
		Short: "Provide an ID to mark a task as done.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Please provide an ID to mark.")
				return
			}
			id := args[0]
			taskFound := false

			for idx, task := range tasks {
				if task.ID == id {
					tasks[idx].Status = Done
					taskFound = true
					fmt.Println("Task marked as done successfully.")
					break
				}
			}

			if taskFound {
				err := saveTasksToFile(filename, tasks)
				if err != nil {
					fmt.Println("Couldn't apply changes to the JSON file.")
				} else {
					fmt.Println("Updates made.")
				}
			} else {
				fmt.Println("No matching ID.")
			}
		},
	}

	markInProgress := &cobra.Command{
		Use:   "in-progress",
		Short: "Provide an ID to mark a task as in-progress",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Please provide an ID to mark.")
				return
			}
			id := args[0]
			taskFound := false

			for idx, task := range tasks {
				if task.ID == id {
					tasks[idx].Status = InProgress
					taskFound = true
					fmt.Println("Task marked as in-progress successfully.")
					break
				}
			}

			if taskFound {
				err := saveTasksToFile(filename, tasks)
				if err != nil {
					fmt.Println("Couldn't apply changes to the JSON file.")
				} else {
					fmt.Println("Updates made.")
				}
			} else {
				fmt.Println("No matching ID.")
			}
		},
	}

	mark.AddCommand(markInProgress)
	mark.AddCommand(markDone)
	rootCmd.AddCommand(AddTasks)
	rootCmd.AddCommand(deleteTasks)
	rootCmd.AddCommand(ListTasks)
	rootCmd.AddCommand(updateTasks)
	rootCmd.AddCommand(mark)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
