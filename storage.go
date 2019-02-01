package main

type Storage interface {
	Get(id string) (*Task, error)
	GetAll() ([]*Task, error)
	Insert(Task) error
	Update(Task) error
}
