package db

type Task struct {
	Key   int
	Value string
}

type DbOperator interface {
	createTask(task string) (int, error)
	readAllTasks() ([]Task, error)
	deleteTask(id int) error
}
