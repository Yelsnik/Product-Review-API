// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: review.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createReview = `-- name: CreateReview :one
INSERT INTO reviews (
  rating, product_id, num_of_reviews
) VALUES (
  $1, $2, $3
) RETURNING id, rating, product_id, num_of_reviews
`

type CreateReviewParams struct {
	Rating       string  `json:"rating"`
	ProductID    string  `json:"product_id"`
	NumOfReviews float64 `json:"num_of_reviews"`
}

func (q *Queries) CreateReview(ctx context.Context, arg CreateReviewParams) (Review, error) {
	row := q.db.QueryRow(ctx, createReview, arg.Rating, arg.ProductID, arg.NumOfReviews)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.Rating,
		&i.ProductID,
		&i.NumOfReviews,
	)
	return i, err
}

const getReview = `-- name: GetReview :one
SELECT id, rating, product_id, num_of_reviews FROM reviews
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetReview(ctx context.Context, id uuid.UUID) (Review, error) {
	row := q.db.QueryRow(ctx, getReview, id)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.Rating,
		&i.ProductID,
		&i.NumOfReviews,
	)
	return i, err
}

const updateReview = `-- name: UpdateReview :one
UPDATE reviews
set rating = COALESCE($1, rating),
  num_of_reviews = COALESCE($2, num_of_reviews)
WHERE id = $3
RETURNING id, rating, product_id, num_of_reviews
`

type UpdateReviewParams struct {
	Rating       pgtype.Text   `json:"rating"`
	NumOfReviews pgtype.Float8 `json:"num_of_reviews"`
	ID           uuid.UUID     `json:"id"`
}

func (q *Queries) UpdateReview(ctx context.Context, arg UpdateReviewParams) (Review, error) {
	row := q.db.QueryRow(ctx, updateReview, arg.Rating, arg.NumOfReviews, arg.ID)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.Rating,
		&i.ProductID,
		&i.NumOfReviews,
	)
	return i, err
}
