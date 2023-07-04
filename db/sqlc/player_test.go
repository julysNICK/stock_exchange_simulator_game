package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/julysNICK/stock_exchange_simulator_game/util"
	"github.com/stretchr/testify/require"
)

func createRandomPlayer(t *testing.T) Player {
	arg := CreatePlayerParams{
		Username: sql.NullString{
			String: util.RandomNamePlayer(),
			Valid:  true,
		},
		HashedPassword: util.RandomPasswordPlayer(),
		FullName:       util.RandomFullNamePlayer(),
		Cash:           util.RandomCashPlayer(),
		Email:          util.RandomEmailPlayer(),
	}

	player, err := testQueries.CreatePlayer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, player)

	require.Equal(t, arg.Username, player.Username)
	require.Equal(t, arg.HashedPassword, player.HashedPassword)
	require.Equal(t, arg.FullName, player.FullName)
	require.Equal(t, arg.Cash, player.Cash)
	require.Equal(t, arg.Email, player.Email)

	require.NotZero(t, player.IDPlayer)
	require.NotZero(t, player.CreatedAt)

	return player
}

func TestCreatePlayer(t *testing.T) {
	createRandomPlayer(t)
}

func TestGetPlayer(t *testing.T) {
	player1 := createRandomPlayer(t)
	player2, err := testQueries.GetPlayerById(context.Background(), player1.IDPlayer)
	require.NoError(t, err)
	require.NotEmpty(t, player2)

	require.Equal(t, player1.Username, player2.Username)
	require.Equal(t, player1.HashedPassword, player2.HashedPassword)
	require.Equal(t, player1.FullName, player2.FullName)
	require.Equal(t, player1.Cash, player2.Cash)
	require.Equal(t, player1.Email, player2.Email)

	require.Equal(t, player1.IDPlayer, player2.IDPlayer)
	require.Equal(t, player1.CreatedAt, player2.CreatedAt)
}

func TestUpdatePlayer(t *testing.T) {
	player1 := createRandomPlayer(t)

	arg := UpdatePlayerParams{
		IDPlayer: player1.IDPlayer,
		Username: sql.NullString{
			String: util.RandomNamePlayer(),
			Valid:  true,
		},
		Cash: sql.NullString{
			String: util.RandomCashPlayer(),
			Valid:  true,
		},
	}

	player2, err := testQueries.UpdatePlayer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, player2)

	require.NotEqual(t, player1.Username, player2.Username)
	require.Equal(t, player1.HashedPassword, player2.HashedPassword)
	require.Equal(t, player1.FullName, player2.FullName)
	require.NotEqual(t, arg.Cash, player2.Cash)

	require.Equal(t, player1.IDPlayer, player2.IDPlayer)
	require.Equal(t, player1.CreatedAt, player2.CreatedAt)
}

func TestDeletePlayer(t *testing.T) {
	player1 := createRandomPlayer(t)

	err := testQueries.DeletePlayer(context.Background(), player1.IDPlayer)
	require.NoError(t, err)

	player2, err := testQueries.GetPlayerById(context.Background(), player1.IDPlayer)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, player2)
}

func TestListPlayers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomPlayer(t)
	}

	arg := ListPlayersParams{
		Limit:  5,
		Offset: 5,
	}

	players, err := testQueries.ListPlayers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, players, 5)

	for _, player := range players {
		require.NotEmpty(t, player)
	}
}

func TestCountPlayer(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomPlayer(t)
	}

	count, err := testQueries.CountPlayers(context.Background())
	require.NoError(t, err)
	require.NotZero(t, count)
}

func TestGetPlayerByEmail(t *testing.T) {
	player1 := createRandomPlayer(t)

	player2, err := testQueries.GetPlayerByEmail(context.Background(), player1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, player2)

	require.Equal(t, player1.Username, player2.Username)
	require.Equal(t, player1.HashedPassword, player2.HashedPassword)
	require.Equal(t, player1.FullName, player2.FullName)
	require.Equal(t, player1.Cash, player2.Cash)
	require.Equal(t, player1.Email, player2.Email)

	require.Equal(t, player1.IDPlayer, player2.IDPlayer)
	require.Equal(t, player1.CreatedAt, player2.CreatedAt)
}

func TestGetPlayerById(t *testing.T) {
	player1 := createRandomPlayer(t)

	player2, err := testQueries.GetPlayerById(context.Background(), player1.IDPlayer)
	require.NoError(t, err)
	require.NotEmpty(t, player2)

	require.Equal(t, player1.Username, player2.Username)
	require.Equal(t, player1.HashedPassword, player2.HashedPassword)
	require.Equal(t, player1.FullName, player2.FullName)
	require.Equal(t, player1.Cash, player2.Cash)
	require.Equal(t, player1.Email, player2.Email)

	require.Equal(t, player1.IDPlayer, player2.IDPlayer)
	require.Equal(t, player1.CreatedAt, player2.CreatedAt)
}

func TestGetPlayerByUsername(t *testing.T) {
	player1 := createRandomPlayer(t)

	player2, err := testQueries.GetPlayerByUsername(context.Background(), player1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, player2)

	require.Equal(t, player1.Username, player2.Username)
	require.Equal(t, player1.HashedPassword, player2.HashedPassword)
	require.Equal(t, player1.FullName, player2.FullName)
	require.Equal(t, player1.Cash, player2.Cash)
	require.Equal(t, player1.Email, player2.Email)

	require.Equal(t, player1.IDPlayer, player2.IDPlayer)
	require.Equal(t, player1.CreatedAt, player2.CreatedAt)
}

func TestRankPlayers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomPlayer(t)
	}

	players, err := testQueries.RankPlayers(context.Background(), RankPlayersParams{
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Len(t, players, 10)

	for _, player := range players {
		require.NotEmpty(t, player)
	}
}
