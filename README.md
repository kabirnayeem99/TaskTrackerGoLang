# Task Tracker CLI

Task tracker is a project used to track and manage your tasks. In this task, you will build a simple command line interface (CLI) to track what you need to do, what you have done, and what you are currently working on. This project will help you practice your programming skills, including working with the filesystem, handling user inputs, and building a simple CLI application.

This is a [task](https://roadmap.sh/projects/task-tracker) from Roadmap.sh.

# Features

The application runs from the command line, accepts user actions and inputs as arguments, and stores the tasks in a JSON file (tasks.json). The user can:

- Add, Update, and Delete tasks

- Mark a task as in progress or done

- List all tasks

- List all tasks that are done

- List all tasks that are not done

- List all tasks that are in progress

## Technical Details

Here are some constraints to guide the implementation:

- Uses Go Programming Language to complete this task.

- Use positional arguments in command line to accept user inputs.

- Uses a JSON file to store the tasks in the current directory.

- The JSON file should be created if it does not exist.

- Uses the native file system module of your programming language to interact with the JSON file.

- Doesn't use any external libraries or frameworks to build this project.

- Hndles errors and edge cases gracefully.

# Usage

The list of commands and their usage is given below:

```sh
# Adding a new task

task-cli add "Buy groceries"

# Output: Task added successfully (ID: 1)

# Updating and deleting tasks

task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1

# Marking a task as in progress or done

task-cli mark-in-progress 1
task-cli mark-done 1

# Listing all tasks

task-cli list

# Listing tasks by status

task-cli list done
task-cli list todo
task-cli list in-progress
```

## Getting Started

### 1. Set Up Your Go Environment

- Install **Go** (latest stable version recommended)
  [https://go.dev/doc/install](https://go.dev/doc/install)

- Verify installation:

  ```sh
  go version
  ```

- Install a code editor (VS Code recommended)
  [https://code.visualstudio.com/](https://code.visualstudio.com/)
  Go extension: [https://marketplace.visualstudio.com/items?itemName=golang.Go](https://marketplace.visualstudio.com/items?itemName=golang.Go)

Official Go documentation:

- [https://go.dev/doc/](https://go.dev/doc/)
- [https://go.dev/doc/tutorial/getting-started](https://go.dev/doc/tutorial/getting-started)

---

### 2. Project Build

- Navigate to project directory:

  ```sh
  cd TaskTrackerGoLang
  ```

- Build

  ```sh
  go build -o task-cli
  ```
