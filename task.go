package main

import "errors"

var (
	ErrTaskAlreadyDone      error = errors.New("task already done")
	ErrTaskAlreadyAbandoned       = errors.New("task already abandoned")
	ErrTaskHasNotStarted          = errors.New("task has not started")
)

type TaskStatus int

const (
	TaskStatusCreated TaskStatus = iota
	TaskStatusStarted
	TaskStatusPaused
	TaskStatusDone
	TaskStatusAbandoned
)

type Task struct {
	id     string
	txt    string
	cat    string
	status TaskStatus
}

func NewTask(id string, txt, cat string, status TaskStatus) *Task {
	return &Task{
		id:     id,
		txt:    txt,
		cat:    cat,
		status: status,
	}
}

func (t *Task) ID() string         { return t.id }
func (t *Task) Txt() string        { return t.txt }
func (t *Task) Category() string   { return t.cat }
func (t *Task) Status() TaskStatus { return t.status }

func (t *Task) Start() error {
	if t.status == TaskStatusDone {
		return ErrTaskAlreadyDone
	} else if t.status == TaskStatusAbandoned {
		return ErrTaskAlreadyAbandoned
	}

	t.status = TaskStatusStarted
	return nil
}

func (t *Task) Pause() error {
	if t.status == TaskStatusCreated {
		return ErrTaskHasNotStarted
	} else if t.status == TaskStatusDone {
		return ErrTaskAlreadyDone
	} else if t.status == TaskStatusAbandoned {
		return ErrTaskAlreadyAbandoned
	}

	t.status = TaskStatusPaused
	return nil
}

func (t *Task) Finish() error {
	if t.status == TaskStatusCreated {
		return ErrTaskHasNotStarted
	} else if t.status == TaskStatusAbandoned {
		return ErrTaskAlreadyAbandoned
	}

	t.status = TaskStatusDone
	return nil
}

func (t *Task) Abandon() error {
	if t.status == TaskStatusDone {
		return ErrTaskAlreadyDone
	}

	t.status = TaskStatusAbandoned
	return nil
}
