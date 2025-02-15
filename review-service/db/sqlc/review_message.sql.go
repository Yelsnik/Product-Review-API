// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: review_message.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createReviewMessage = `-- name: CreateReviewMessage :one
INSERT INTO review_messages (
  review,
  score,
  label,
  review_id
) VALUES (
  $1, $2, $3, $4
) RETURNING id, review, score, label, review_id, updated_at, created_at
`

type CreateReviewMessageParams struct {
	Review   string    `json:"review"`
	Score    float64   `json:"score"`
	Label    string    `json:"label"`
	ReviewID uuid.UUID `json:"review_id"`
}

func (q *Queries) CreateReviewMessage(ctx context.Context, arg CreateReviewMessageParams) (ReviewMessage, error) {
	row := q.db.QueryRow(ctx, createReviewMessage,
		arg.Review,
		arg.Score,
		arg.Label,
		arg.ReviewID,
	)
	var i ReviewMessage
	err := row.Scan(
		&i.ID,
		&i.Review,
		&i.Score,
		&i.Label,
		&i.ReviewID,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getReviewMessage = `-- name: GetReviewMessage :one
SELECT id, review, score, label, review_id, updated_at, created_at FROM review_messages
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetReviewMessage(ctx context.Context, id uuid.UUID) (ReviewMessage, error) {
	row := q.db.QueryRow(ctx, getReviewMessage, id)
	var i ReviewMessage
	err := row.Scan(
		&i.ID,
		&i.Review,
		&i.Score,
		&i.Label,
		&i.ReviewID,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getReviewMessages = `-- name: GetReviewMessages :many
SELECT id, review, score, label, review_id, updated_at, created_at FROM review_messages
WHERE id = $1
`

func (q *Queries) GetReviewMessages(ctx context.Context, id uuid.UUID) ([]ReviewMessage, error) {
	rows, err := q.db.Query(ctx, getReviewMessages, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ReviewMessage
	for rows.Next() {
		var i ReviewMessage
		if err := rows.Scan(
			&i.ID,
			&i.Review,
			&i.Score,
			&i.Label,
			&i.ReviewID,
			&i.UpdatedAt,
			&i.CreatedAt,
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

const getReviewMessagesByReview = `-- name: GetReviewMessagesByReview :many
SELECT id, review, score, label, review_id, updated_at, created_at FROM review_messages
WHERE review_id = $1
`

func (q *Queries) GetReviewMessagesByReview(ctx context.Context, reviewID uuid.UUID) ([]ReviewMessage, error) {
	rows, err := q.db.Query(ctx, getReviewMessagesByReview, reviewID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ReviewMessage
	for rows.Next() {
		var i ReviewMessage
		if err := rows.Scan(
			&i.ID,
			&i.Review,
			&i.Score,
			&i.Label,
			&i.ReviewID,
			&i.UpdatedAt,
			&i.CreatedAt,
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

const reviewCount = `-- name: ReviewCount :one
SELECT COUNT(*) AS review_count
FROM review_messages
WHERE review_id = $1
`

func (q *Queries) ReviewCount(ctx context.Context, reviewID uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, reviewCount, reviewID)
	var review_count int64
	err := row.Scan(&review_count)
	return review_count, err
}

const updateReviewMessage = `-- name: UpdateReviewMessage :one
UPDATE review_messages
set review = COALESCE($1, review),
  score = COALESCE($2, score),
  label = COALESCE($3, label),
  updated_at = COALESCE($4, updated_at)
WHERE id = $5
RETURNING id, review, score, label, review_id, updated_at, created_at
`

type UpdateReviewMessageParams struct {
	Review    pgtype.Text        `json:"review"`
	Score     pgtype.Float8      `json:"score"`
	Label     pgtype.Text        `json:"label"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	ID        uuid.UUID          `json:"id"`
}

func (q *Queries) UpdateReviewMessage(ctx context.Context, arg UpdateReviewMessageParams) (ReviewMessage, error) {
	row := q.db.QueryRow(ctx, updateReviewMessage,
		arg.Review,
		arg.Score,
		arg.Label,
		arg.UpdatedAt,
		arg.ID,
	)
	var i ReviewMessage
	err := row.Scan(
		&i.ID,
		&i.Review,
		&i.Score,
		&i.Label,
		&i.ReviewID,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}
