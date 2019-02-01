package main

type Storage interface {
	Get(id string) (*Task, error)
	GetAll() ([]*Task, error)
	Insert(Task) error
	Update(Task) error
}

type FileStorage struct {
}

func (s *FileStorage) Get(id string) (*Task, error) {
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
