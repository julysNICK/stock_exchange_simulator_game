package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomPortfolio(t *testing.T) Portfolio{
 randomPlayer := createRandomPlayer(t)

	portfolio, err := testQueries.CreatePortfolio(context.Background(), randomPlayer.IDPlayer)
	require.NoError(t, err)
	require.NotEmpty(t, portfolio)
	println(portfolio.ID)

	require.NotZero(t, portfolio.ID)
	require.NotZero(t, portfolio.CreatedAt)
	return portfolio
}

func TestCreatePortfolio(t *testing.T) {
	createRandomPortfolio(t)
}

func TestGetPortfolioById(t *testing.T) {
	portfolio1 := createRandomPortfolio(t)
	portfolio2, err := testQueries.GetPortfolioById(context.Background(), portfolio1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, portfolio2)

	require.Equal(t, portfolio1.ID, portfolio2.ID)
	require.Equal(t, portfolio1.PlayerID, portfolio2.PlayerID)
	require.WithinDuration(t, portfolio1.CreatedAt, portfolio2.CreatedAt, time.Second)
}

func TestCountPortfolio(t *testing.T){
	for i := 0; i < 10; i++ {
		createRandomPortfolio(t)
	}
	count, err := testQueries.CountPortfolio(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, count)
}

func TestDeletePortfolio(t *testing.T) {
	portfolio1 := createRandomPortfolio(t)
	portfolio2, err := testQueries.DeletePortfolio(context.Background(), portfolio1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, portfolio2)

	require.Equal(t, portfolio1.ID, portfolio2.ID)
	require.Equal(t, portfolio1.PlayerID, portfolio2.PlayerID)
	require.WithinDuration(t, portfolio1.CreatedAt, portfolio2.CreatedAt, time.Second)

	portfolio3, err := testQueries.GetPortfolioById(context.Background(), portfolio1.ID)
	require.Error(t, err)
	require.Empty(t, portfolio3)
}

func TestGetPortfolioByPlayerId(t *testing.T) {
	portfolio1 := createRandomPortfolio(t)
	portfolio2, err := testQueries.GetPortfolioByPlayerId(context.Background(), portfolio1.PlayerID)
	require.NoError(t, err)
	require.NotEmpty(t, portfolio2)

	require.Equal(t, portfolio1.ID, portfolio2.ID)
	require.Equal(t, portfolio1.PlayerID, portfolio2.PlayerID)
	require.WithinDuration(t, portfolio1.CreatedAt, portfolio2.CreatedAt, time.Second)
}

func TestListPortfolio(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomPortfolio(t)
	}
	arg := ListPortfolioParams{
		Limit:  5,
		Offset: 5,
	}
	portfolio, err := testQueries.ListPortfolio(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, portfolio, 5)
	for _, portfolio := range portfolio {
		require.NotEmpty(t, portfolio)
	}
}

func TestUpdatePortfolio(t *testing.T) {
	portfolio1 := createRandomPortfolio(t)
	arg := UpdatePortfolioParams{
		ID: portfolio1.ID,
		PlayerID: portfolio1.PlayerID,
	}
	portfolio2, err := testQueries.UpdatePortfolio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, portfolio2)

	require.Equal(t, portfolio1.ID, portfolio2.ID)
	require.Equal(t, portfolio1.PlayerID, portfolio2.PlayerID)
	require.WithinDuration(t, portfolio1.CreatedAt, portfolio2.CreatedAt, time.Second)
}