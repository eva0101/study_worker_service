package employee

import (
	"errors"
	"sync"
)

type Employee struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Position string `json:"position"`
}

type Num struct {
	ID int `json:"id"`
}

type ListEmployees struct {
	list   []Employee
	mu     sync.Mutex
	IDnext int
}

func (l *ListEmployees) Add(e Employee) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.IDnext++
	e.ID = l.IDnext

	l.list = append(l.list, e)
}

func (l *ListEmployees) Get() []Employee {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.list
}

func (l *ListEmployees) Delete(n Num) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	for i, e := range l.list {
		if e.ID == n.ID {
			l.list = append(l.list[:i], l.list[i+1:]...)
			return nil
		}
	}

	return errors.New("id not found")

}
