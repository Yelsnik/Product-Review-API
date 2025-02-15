package clients

import (
	"context"
	"review-service/nlp"

	"google.golang.org/grpc"
)

type Client interface {
	Analyze(ctx context.Context, text string) (*nlp.SentimentResponse, error)
}

type ClientgRPC struct {
	NLPClient nlp.SentimentAnalysisClient
}

func NewClientgRPC(nlpconn *grpc.ClientConn) Client {
	return &ClientgRPC{
		NLPClient: nlp.NewSentimentAnalysisClient(nlpconn),
	}
}
