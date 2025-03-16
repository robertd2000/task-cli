# Task Tracker

Solution for the [task-tracker](https://github.com/robertd2000/task-cli) challenge from [roadmap.sh](https://roadmap.sh/) (without any external libraries).

## How to run

Clone the repository and run the following command:

```bash
git clone https://github.com/robertd2000/task-cli.git
cd backend-projects/task-tracker
```

Run the following command to build and run the project:

```bash
# To add a task
go run . add "Buy groceries"

# To update a task
go run . update 1 "Buy groceries and cook dinner"

# To delete a task
go run . delete 1

# To mark a task as in progress/done/todo
go run . mark-in-progress 1
go run . mark-done 1
go run . mark-todo 1

# To list all tasks
go run . list
go run . list done
go run . list todo
go run . list in-progress
```
