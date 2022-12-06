// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: chronology.sql

package sqlc

import (
	"context"
)

const createChronology = `-- name: CreateChronology :one
INSERT INTO chronology (title)
VALUES ($1) RETURNING id
`

func (q *Queries) CreateChronology(ctx context.Context, title string) (string, error) {
	row := q.db.QueryRow(ctx, createChronology, title)
	var id string
	err := row.Scan(&id)
	return id, err
}

const deleteChronology = `-- name: DeleteChronology :exec
DELETE FROM chronology
WHERE id = $1
`

func (q *Queries) DeleteChronology(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteChronology, id)
	return err
}

const getAllChronology = `-- name: GetAllChronology :many
SELECT id, title, created_at, updated_at
FROM chronology
WHERE chronology.title ilike '%' || $1::VARCHAR || '%'
ORDER BY created_at DESC OFFSET $2
LIMIT $3
`

type GetAllChronologyParams struct {
	Search string `json:"search"`
	Offset int32  `json:"offset"`
	Limit  int32  `json:"limit"`
}

func (q *Queries) GetAllChronology(ctx context.Context, arg GetAllChronologyParams) ([]Chronology, error) {
	rows, err := q.db.Query(ctx, getAllChronology, arg.Search, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Chronology
	for rows.Next() {
		var i Chronology
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateChronology = `-- name: UpdateChronology :exec
UPDATE chronology
SET title = $1,
    updated_at = now()
WHERE id = $2
`

type UpdateChronologyParams struct {
	Title string `json:"title"`
	ID    string `json:"id"`
}

func (q *Queries) UpdateChronology(ctx context.Context, arg UpdateChronologyParams) error {
	_, err := q.db.Exec(ctx, updateChronology, arg.Title, arg.ID)
	return err
}