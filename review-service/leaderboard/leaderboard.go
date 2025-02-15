package leaderboard

import (
	"context"

	"github.com/redis/go-redis/v9"
)

const key = "product_leaderboard"

type LeaderboardEntry struct {
	ProductId string
	score     float64
}

type Leaderboard interface {
	UpdateLeaderBoard(ctx context.Context, productId string, score float64) error
	GetTopProducts(ctx context.Context, limit int64) ([]LeaderboardEntry, error)
}

type LeaderboardClient struct {
	redis *redis.Client
}

func NewLeaderBoardClient(redis *redis.Client) Leaderboard {
	return &LeaderboardClient{
		redis: redis,
	}
}

func (l *LeaderboardClient) UpdateLeaderBoard(ctx context.Context, productId string, score float64) error {
	err := l.redis.ZIncrBy(ctx, key, score, productId).Err()

	return err
}

func (l *LeaderboardClient) GetTopProducts(ctx context.Context, limit int64) ([]LeaderboardEntry, error) {
	results, err := l.redis.ZRevRangeWithScores(ctx, key, 0, limit-1).Result()

	var leaderboard []LeaderboardEntry
	for _, v := range results {
		data := LeaderboardEntry{
			ProductId: v.Member.(string),
			score:     v.Score,
		}

		leaderboard = append(leaderboard, data)
	}

	return leaderboard, err
}
