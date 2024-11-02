-- name: FetchAllUsers :many
SELECT
    id,
    titletext,
    iscompleted,
    category,
    deadline
FROM
    todo;

-- name: GetTodoById :one
SELECT
    id,
    titletext,
    iscompleted,
    category,
    deadline
FROM
    todo
WHERE
    id = $1;

-- name: DeleteTodoById :exec
DELETE FROM todo
WHERE id = $1;

-- name: CreateTodo :one
INSERT INTO todo(titletext, iscompleted, category, deadline)
    VALUES ($1, $2, $3, $4, $5)
RETURNING
    id;