package todo

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

var getTodo = `
	SELECT * 
	FROM todo_list
	WHERE id = $1
`
var selectTodos = `
	SELECT *
	FROM todo_list
`

var insertTodo = `
	INSERT INTO todo_list(title)
	VALUES ($1)
	RETURNING id, title, created_at, updated_at
`

var deleteTodo = `
	DELETE FROM todo_list 
	WHERE id = $1
`

var selectTasks = `
	SELECT 
	    task.id, task.list_id, task.title, task.completed
	FROM task
	WHERE task.list_id = $1
`

var insertTask = `
	INSERT into task
	(title,list_id)
	VALUES (:title,:list_id)
`
var deleteTask = `
	DELETE FROM task
	WHERE id = $1
`

var markTaskCompleted = `
	UPDATE task
	SET completed=true
	WHERE id = $1
	`

type repo struct {
	db          *sqlx.DB
	insertTodo  *sqlx.Stmt
	deleteTodo  *sqlx.Stmt
	selectTodos *sqlx.Stmt
	getTodo     *sqlx.Stmt

	insertTask        *sqlx.NamedStmt
	deleteTask        *sqlx.Stmt
	selectTasks       *sqlx.Stmt
	markTaskCompleted *sqlx.Stmt
}

func newRepo(db *sqlx.DB) *repo {
	return &repo{
		db: db,
	}
}

func (r *repo) Prepare() error {
	var err error
	r.insertTodo, err = r.db.Preparex(insertTodo)
	if err != nil {
		return err
	}
	r.getTodo, err = r.db.Preparex(getTodo)
	if err != nil {
		return err
	}
	r.selectTodos, err = r.db.Preparex(selectTodos)
	if err != nil {
		return err
	}
	r.deleteTodo, err = r.db.Preparex(deleteTodo)
	if err != nil {
		return err
	}
	r.insertTask, err = r.db.PrepareNamed(insertTask)
	if err != nil {
		return err
	}
	r.selectTasks, err = r.db.Preparex(selectTasks)
	if err != nil {
		return err
	}
	r.deleteTask, err = r.db.Preparex(deleteTask)
	if err != nil {
		return err
	}
	r.markTaskCompleted, err = r.db.Preparex(markTaskCompleted)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) CreateTodoList(ctx context.Context, title string) (List, error) {
	var todoList List
	err := r.insertTodo.GetContext(ctx, &todoList, title)
	if err != nil {
		return List{}, err
	}
	return todoList, nil
}

func (r *repo) GetTodoList(ctx context.Context, todoID int64) (List, error) {
	var list List
	err := r.getTodo.GetContext(ctx, &list, todoID)
	if err != nil {
		return List{}, err
	}
	return list, nil
}
func (r *repo) GetAllTodoLists(ctx context.Context) ([]List, error) {
	var lists []List
	err := r.selectTodos.SelectContext(ctx, &lists)
	if err != nil {
		return nil, err
	}
	return lists, nil
}

func (r *repo) InsertTodo(ctx context.Context, list List) (int64, error) {
	var id int64
	err := r.insertTodo.GetContext(ctx, &id, list)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repo) DeleteTodoList(ctx context.Context, id int64) error {
	_, err := r.deleteTodo.ExecContext(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) GetTask(ctx context.Context, taskID int64) (Task, error) {
	var task Task
	err := r.getTodo.GetContext(ctx, &task, taskID)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (r *repo) InsertTask(ctx context.Context, task Task) (int64, error) {
	var id int64
	err := r.insertTask.GetContext(ctx, &id, task)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	return id, nil
}

func (r *repo) DeleteTask(ctx context.Context, taskID int64) error {
	_, err := r.deleteTask.ExecContext(ctx, taskID)
	return err
}

func (r *repo) SelectTasks(ctx context.Context, todoID int64) ([]Task, error) {
	var tasks []Task
	err := r.selectTasks.SelectContext(ctx, &tasks, todoID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *repo) MarkTaskCompleted(ctx context.Context, id int64) error {
	_, err := r.markTaskCompleted.ExecContext(ctx, id)
	return err
}
