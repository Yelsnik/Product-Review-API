package leaderboard

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Leaderboard interface{
	UpdateLeaderBoard(ctx context.Context, productId string, score float64) error
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
	err := l.redis.ZIncrBy(ctx, "product_leaderboard", score, productId).Err()
	
	return err
}


