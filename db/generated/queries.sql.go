// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package generated

import (
	"context"
	"database/sql"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todo (title, text, iscompleted, category, deadline)
    VALUES ($1, $2, $3, $4, $5)
RETURNING
    id
`

type CreateTodoParams struct {
	Title       string         `json:"title"`
	Text        string         `json:"text"`
	Iscompleted bool           `json:"iscompleted"`
	Category    sql.NullString `json:"category"`
	Deadline    sql.NullTime   `json:"deadline"`
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createTodo,
		arg.Title,
		arg.Text,
		arg.Iscompleted,
		arg.Category,
		arg.Deadline,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const deleteTodoById = `-- name: DeleteTodoById :exec
DELETE FROM todo
WHERE id = $1
`

func (q *Queries) DeleteTodoById(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTodoById, id)
	return err
}

const fetchAllTodos = `-- name: FetchAllTodos :many
SELECT
    id,
    title,
    text,
    iscompleted,
    category,
    deadline
FROM
    todo
`

func (q *Queries) FetchAllTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, fetchAllTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Text,
			&i.Iscompleted,
			&i.Category,
			&i.Deadline,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTodoById = `-- name: GetTodoById :one
SELECT
    id,
    title,
    text,
    iscompleted,
    category,
    deadline
FROM
    todo
WHERE
    id = $1
`

func (q *Queries) GetTodoById(ctx context.Context, id int64) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodoById, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Text,
		&i.Iscompleted,
		&i.Category,
		&i.Deadline,
	)
	return i, err
}
