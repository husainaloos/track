package main

import "testing"

func TestStart(t *testing.T) {
	tests := []struct {
		name    string
		task    *Task
		wantErr bool
		err     error
	}{
		{
			"should start task if task has not started",
			NewTask("id", "txt", "cat", TaskStatusCreated),
			false,
			nil,
		},
		{
			"should do nothing if task is already started",
			NewTask("id", "txt", "cat", TaskStatusStarted),
			false,
			nil,
		},
		{
			"should start task if it is paused",
			NewTask("id", "txt", "cat", TaskStatusPaused),
			false,
			nil,
		},
		{
			"should fail to start task if task is done",
			NewTask("id", "txt", "cat", TaskStatusDone),
			true,
			ErrTaskAlreadyDone,
		},
		{
			"should fail to start task if task is abandoned",
			NewTask("id", "txt", "cat", TaskStatusAbandoned),
			true,
			ErrTaskAlreadyAbandoned,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.task.Start()
			if (err == nil) == test.wantErr {
				t.Fatalf("got err=%v, but expected=%v", err, test.err)
			}
			if err != nil {
				return
			}
			status := test.task.Status()
			if status != TaskStatusStarted {
				t.Fatalf("got status=%v, but expected=%v", status, TaskStatusStarted)
			}
		})
	}
}

func TestPause(t *testing.T) {
	tests := []struct {
		name    string
		task    *Task
		wantErr bool
		err     error
	}{
		{
			"should fail to pause task if task has not started",
			NewTask("id", "txt", "cat", TaskStatusCreated),
			true,
			ErrTaskHasNotStarted,
		},
		{
			"should pause task if it is started",
			NewTask("id", "txt", "cat", TaskStatusStarted),
			false,
			nil,
		},
		{
			"should do nothing if task is already paused",
			NewTask("id", "txt", "cat", TaskStatusPaused),
			false,
			nil,
		},
		{
			"should fail to pause task if it is done",
			NewTask("id", "txt", "cat", TaskStatusDone),
			true,
			ErrTaskAlreadyDone,
		},
		{
			"should fail to pause task if task is abandoned",
			NewTask("id", "txt", "cat", TaskStatusAbandoned),
			true,
			ErrTaskAlreadyAbandoned,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.task.Pause()
			if (err == nil) == test.wantErr {
				t.Fatalf("got err=%v, but expected=%v", err, test.err)
			}
			if err != nil {
				return
			}
			status := test.task.Status()
			if status != TaskStatusPaused {
				t.Fatalf("got status=%v, but expected=%v", status, TaskStatusPaused)
			}
		})
	}
}

func TestFinish(t *testing.T) {
	tests := []struct {
		name    string
		task    *Task
		wantErr bool
		err     error
	}{
		{
			"should fail to finish task if it was only created",
			NewTask("id", "txt", "cat", TaskStatusCreated),
			true,
			ErrTaskHasNotStarted,
		},
		{
			"should finish task if it is started",
			NewTask("id", "txt", "cat", TaskStatusStarted),
			false,
			nil,
		},
		{
			"should finish task if it is paused",
			NewTask("id", "txt", "cat", TaskStatusPaused),
			false,
			nil,
		},
		{
			"should do nothing if task already finished",
			NewTask("id", "txt", "cat", TaskStatusDone),
			false,
			nil,
		},
		{
			"should fail to finish task if task is abandoned",
			NewTask("id", "txt", "cat", TaskStatusAbandoned),
			true,
			ErrTaskAlreadyAbandoned,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.task.Finish()
			if (err == nil) == test.wantErr {
				t.Fatalf("got err=%v, but expected=%v", err, test.err)
			}
			if err != nil {
				return
			}
			status := test.task.Status()
			if status != TaskStatusDone {
				t.Fatalf("got status=%v, but expected=%v", status, TaskStatusDone)
			}
		})
	}
}

func TestAbandon(t *testing.T) {
	tests := []struct {
		name    string
		task    *Task
		wantErr bool
		err     error
	}{
		{
			"should abandon task if task has not started",
			NewTask("id", "txt", "cat", TaskStatusCreated),
			false,
			nil,
		},
		{
			"should abandon task if task is started",
			NewTask("id", "txt", "cat", TaskStatusStarted),
			false,
			nil,
		},
		{
			"should abandon task if task is paused",
			NewTask("id", "txt", "cat", TaskStatusPaused),
			false,
			nil,
		},
		{
			"should do nothing if task is already abandoned",
			NewTask("id", "txt", "cat", TaskStatusAbandoned),
			false,
			nil,
		},
		{
			"should fail to abandon task if task is done",
			NewTask("id", "txt", "cat", TaskStatusDone),
			true,
			ErrTaskAlreadyDone,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.task.Abandon()
			if (err == nil) == test.wantErr {
				t.Fatalf("got err=%v, but expected=%v", err, test.err)
			}
			if err != nil {
				return
			}
			status := test.task.Status()
			if status != TaskStatusAbandoned {
				t.Fatalf("got status=%v, but expected=%v", status, TaskStatusAbandoned)
			}
		})
	}
}
