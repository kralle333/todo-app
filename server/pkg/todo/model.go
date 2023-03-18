package todo

import (
	"github.com/volatiletech/null/v8"
)

type Task struct {
	ID        int64  `json:"id" db:"id"`
	ListID    int64  `json:"list_id" db:"list_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
type List struct {
	ID        int64     `json:"id" db:"id"`
	Title     string    `json:"title"`
	CreatedAt null.Time `json:"created_at" db:"created_at"`
	UpdatedAt null.Time `json:"updated_at" db:"updated_at"`
}
