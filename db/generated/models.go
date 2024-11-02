// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package generated

import (
	"database/sql"
)

type Todo struct {
	ID          int32          `json:"id"`
	Title       string         `json:"title"`
	Text        string         `json:"text"`
	Iscompleted bool           `json:"iscompleted"`
	Category    sql.NullString `json:"category"`
	Deadline    sql.NullTime   `json:"deadline"`
}
