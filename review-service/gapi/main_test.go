package gapi

import (
	"review-service/clients"
	db "review-service/db/sqlc"
	"review-service/helpers"
	"review-service/leaderboard"
	"review-service/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store, helpers helpers.Helpers, client clients.Client, leaderboard leaderboard.Leaderboard) *Server {
	config := util.Config{}

	server, err := NewServer(config, store, helpers, client, leaderboard)
	require.NoError(t, err)

	return server
}
