-- name: CreateReviewMessage :one
INSERT INTO review_messages (
  review,
  score,
  label,
  review_id
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetReviewMessage :one
SELECT * FROM review_messages
WHERE id = $1 LIMIT 1;

-- name: GetReviewMessages :one
SELECT * FROM review_messages
WHERE id = $1;

-- name: ReviewCount :one
SELECT COUNT(*) AS review_count
FROM review_messages
WHERE review_id = $1;

-- name: UpdateReviewMessage :one
UPDATE review_messages
set review = COALESCE(sqlc.narg(review), review),
  score = COALESCE(sqlc.narg(score), score),
  label = COALESCE(sqlc.narg(label), label),
  updated_at = COALESCE(sqlc.narg(updated_at), updated_at)
WHERE id = sqlc.arg(id)
RETURNING *;