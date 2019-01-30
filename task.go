package main

type Status int

type Task struct {
	id     string
	txt    string
	cat    string
	status Status
}

func NewTask(id string, txt, cat string, status Status) *Task {
	return &Task{
		id:     id,
		txt:    txt,
		cat:    cat,
		status: status,
	}
}

func (t *Task) ID() string       { return t.id }
func (t *Task) Txt() string      { return t.txt }
func (t *Task) Category() string { return t.cat }
func (t *Task) Status() string   { return t.status }

func (t *Task) Start() {
}

func (t *Task) Pause() {
}

func (t *Task) End() {
}

func (t *Task) Abandon() {
}
