package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type ReviewTxResult struct {
	Review  Review        `json:"review"`
	Message ReviewMessage `json:"review_message"`
}

type ReviewTxParams struct {
	Review   string    `json:"review"`
	Score    float64   `json:"score"`
	Label    string    `json:"label"`
	ReviewID uuid.UUID `json:"review_id"`
}

func (store *SQLStore) AddReviewTx(ctx context.Context, arg CreateReviewMessageParams) (ReviewTxResult, error) {
	var result ReviewTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// create a review message
		result.Message, err = q.CreateReviewMessage(ctx, CreateReviewMessageParams{
			Review:   arg.Review,
			Score:    arg.Score,
			Label:    arg.Label,
			ReviewID: arg.ReviewID,
		})
		if err != nil {
			return err
		}

		// count review messages
		count, err := q.ReviewCount(ctx, arg.ReviewID)
		if err != nil {
			return err
		}

		// update review
		result.Review, err = q.UpdateReview(ctx, UpdateReviewParams{
			NumOfReviews: pgtype.Float8{
				Float64: float64(count),
				Valid:   true,
			},
			ID: arg.ReviewID,
		})
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}
