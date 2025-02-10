-- name: CreateReview :one
INSERT INTO reviews (
  rating, product_id, num_of_reviews
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetReview :one
SELECT * FROM reviews
WHERE id = $1 LIMIT 1;

-- name: UpdateReview :one
UPDATE reviews
set rating = COALESCE(sqlc.narg(rating), rating),
  num_of_reviews = COALESCE(sqlc.narg(num_of_reviews), num_of_reviews)
WHERE id = sqlc.arg(id)
RETURNING *;