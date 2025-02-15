package clients

import (
	"context"
	"review-service/nlp"
)

func (client *ClientgRPC) Analyze(ctx context.Context, text string) (*nlp.SentimentResponse, error) {
	in := &nlp.SentimentRequest{
		Text: text,
	}

	return client.NLPClient.Analyze(ctx, in)
}
