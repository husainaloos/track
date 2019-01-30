# Description of the API

This is not a finished document by any mean.

Each task has:
- text
- category
- status: created, active, paused, done, abandoned

## Creating a task

- Creating a task without a category: `task create my amazing task`
- Creating a task with a category: `task create --category coding my amazing task`

By default, these create the task, but they don't start it. If you want to create and start the task, then you can use
- `task create-and-start my amazing task` or 
- `task create-and-start --category coding my amazing task`

The creation of a task should return an id, and this id should be sequential.

## Listing tasks

- You can list the tasks with `task list`. This should list all tasks that are not done or deleted. List should include id of the task as well as the description.


