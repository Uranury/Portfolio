# CLI Todo Application

A simple command-line interface (CLI) application for managing tasks. This project allows you to add, delete, list, and update tasks, with a focus on providing a user-friendly experience.

## Features

- **Add Tasks:** Easily add new tasks with descriptions.
- **Delete Tasks:** Remove tasks by their unique ID.
- **List Tasks:** View all existing tasks with their status.
- **Update Tasks:** Modify the description of existing tasks.
- **Mark Tasks:** Change the status of tasks between "todo," "in-progress," and "done."

## Installation

1. Clone the repository to your local machine:
   ```bash
   git clone https://github.com/your-username/your-repo-name.git
   ```
2. Change to the project directory:
   ```bash
   cd your-repo-name
   ```

3. Install the required packages (if any). Ensure you have Go installed on your machine.

## Usage

Run the application using the following command:

```bash
go run main.go [command] [arguments]
```

### Commands

- **Add a task:**
  ```bash
  go run main.go add "Task description"
  ```

- **Delete a task:**
  ```bash
  go run main.go delete <task_id>
  ```

- **List all tasks:**
  ```bash
  go run main.go list
  ```

- **Update a task:**
  ```bash
  go run main.go update <task_id> "New task description"
  ```

- **Mark a task as done:**
  ```bash
  go run main.go mark done <task_id>
  ```

- **Mark a task as in-progress:**
  ```bash
  go run main.go mark in-progress <task_id>
  ```

## Data Persistence

Tasks are saved in a JSON file (`tasks.json`). The application will automatically load tasks from this file upon startup and save changes whenever tasks are added, updated, or deleted.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Cobra](https://github.com/spf13/cobra) for creating a powerful CLI application framework in Go.
- Inspiration from various task management tools.

```

### Tips for Customization
- **Replace placeholders:** Make sure to replace `your-username` and `your-repo-name` with your actual GitHub username and the name of the repository.
- **Additional features:** If you add more features in the future, update the `Features` section accordingly.
- **Add screenshots:** If you want to make your README more visually appealing, consider adding screenshots or examples of how the commands work.
