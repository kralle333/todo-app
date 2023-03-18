package todo

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Component struct {
	repo repo
}

func NewTodoComponent(db *sqlx.DB) Component {

	return Component{
		repo: *newRepo(db),
	}
}

func (t *Component) PrepareDB() error {
	return t.repo.Prepare()
}

func (t *Component) CreateTodoList(ctx context.Context, title string) (List, error) {
	return t.repo.CreateTodoList(ctx, title)
}

func (t *Component) GetTodoList(ctx context.Context, id int64) (List, error) {
	return t.repo.GetTodoList(ctx, id)
}

func (t *Component) GetAllTodoLists(ctx context.Context) ([]List, error) {
	return t.repo.GetAllTodoLists(ctx)
}

func (t *Component) DeleteTodoList(ctx context.Context, id int64) error {
	return t.repo.DeleteTodoList(ctx, id)
}

func (t *Component) GetAllTasks(ctx context.Context, todoListID int64) ([]Task, error) {
	return t.repo.SelectTasks(ctx, todoListID)
}

func (t *Component) GetTask(ctx context.Context, id int64) (Task, error) {
	return t.repo.GetTask(ctx, id)
}

func (t *Component) AddTask(ctx context.Context, id int64, title string) (Task, error) {
	taskID, err := t.repo.InsertTask(ctx, Task{
		ListID: id,
		Title:  title,
	})
	if err != nil {
		return Task{}, err
	}
	return Task{
		ID:        taskID,
		ListID:    id,
		Title:     title,
		Completed: false,
	}, nil
}
func (t *Component) MarkTaskCompleted(ctx context.Context, taskID int64) error {
	return t.repo.MarkTaskCompleted(ctx, taskID)
}

func (t *Component) DeleteTask(ctx context.Context, taskID int64) error {
	return t.repo.DeleteTask(ctx, taskID)
}
