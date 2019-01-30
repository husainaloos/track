# Description of the API

This is not a finished document by any mean.

Each task has:
- id
- text
- category
- status: created, active, paused, done, abandoned

## Creating a task

- Creating a task without a category: `track create my amazing task`
- Creating a task with a category: `track create coding my amazing task category:coding`

By default, these create the task, but are not active. If you want to create and activate the task, then you can use
- `track create my amazing task active:true` or 
- `track create coding my amazing task category:coding active:true`

The creation of a task should return an id, and this id should be sequential.

## Listing tasks

- You can list the tasks with `track list`. This should list all tasks that are not done or deleted. List should include id of the task as well as the description.


## Task life-cycle

- create a task: `track create clean office`
- start a task by id: `track start 12`
- pause a task by id: `track pause 12`
- mark a task as done: `track done 12`
- abandon a task: `track abandon 12`

