package main

import (
	"database/sql"
)

type FileStorage struct {
	db sql.DB
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

	return nil, nil
}

func (s *FileStorage) GetAll() ([]*Task, error) {
	return nil, nil
}

func (s *FileStorage) Insert(Task) error {
	return nil
}

func (s *FileStorage) Update(Task) error {
	return nil
}
