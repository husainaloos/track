package main

import (
	"database/sql"
)

type FileStorage struct {
	db sql.DB
}

func NewFileStorage() (*FileStorage, error) {

}

func (s *FileStorage) Get(id string) (*Task, error) {
	rows, err := s.db.Query("SELECT id, txt, cat, status FROM tasks WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var id, txt, cat string
		var status TaskStatus
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		if err := rows.Scan(&txt); err != nil {
			return nil, err
		}
		if err := rows.Scan(&cat); err != nil {
			return nil, err
		}
		if err := rows.Scan(&status); err != nil {
			return nil, err
		}
		return NewTask(id, txt, cat, status), nil
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *FileStorage) GetAll() ([]*Task, error) {
	rows, err := s.db.Query("SELECT id, txt, cat, status FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]*Task, 0)
	for rows.Next() {
		var id, txt, cat string
		var status TaskStatus
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		if err := rows.Scan(&txt); err != nil {
			return nil, err
		}
		if err := rows.Scan(&cat); err != nil {
			return nil, err
		}
		if err := rows.Scan(&status); err != nil {
			return nil, err
		}
		task := NewTask(id, txt, cat, status)
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *FileStorage) Insert(task Task) error {
	_, err := s.db.Exec("INSERT INTO tasks (id, txt, cat, status) VALUES (?, ?, ?, ?)",
		task.ID(), task.Txt(), task.Category(), task.Status())
	return err
}

func (s *FileStorage) Update(task Task) error {
	_, err := s.db.Exec("UPDATE tasks SET id=?, txt=?, cat=?, status=? WHERE id=?",
		task.ID(), task.Txt(), task.Category(), task.Status(), task.ID())
	return err
}
