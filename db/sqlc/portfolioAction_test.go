package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomPortfolioAction(
	t *testing.T,
) PortfolioAction{
	randomPortfolio := createRandomPortfolio(t)
	randomAction := createRandomAction(t)
	arg := CreatePortfolioActionParams{
		PortfolioID: randomPortfolio.ID,
		ActionID: randomAction.ID,
		PlayerID: randomPortfolio.PlayerID,
		PurchasePrice: "1.00",
		Quantity: 1,
	}

	portfolioAction, err := testQueries.CreatePortfolioAction(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, portfolioAction)

	require.Equal(t, arg.PortfolioID, portfolioAction.PortfolioID)

	require.Equal(t, arg.ActionID, portfolioAction.ActionID)

	require.Equal(t, arg.PlayerID, portfolioAction.PlayerID)

	require.Equal(t, arg.PurchasePrice, portfolioAction.PurchasePrice)

	require.Equal(t, arg.Quantity, portfolioAction.Quantity)

	require.NotZero(t, portfolioAction.ID)

	require.NotZero(t, portfolioAction.CreatedAt)

	return portfolioAction
}

func TestCreatePortfolioAction(t *testing.T) {
	createRandomPortfolioAction(t)
}

func TestCountPortfolioActions(t *testing.T){

	for i := 0; i < 10; i++ {
		createRandomPortfolioAction(t)
	}

	count, err := testQueries.CountPortfolioActions(context.Background())

	require.NoError(t, err)

	require.NotEmpty(t, count)

}


func TestDeletePortfolioAction(t *testing.T) {
	portfolioAction1 := createRandomPortfolioAction(t)
	portfolioAction2, err := testQueries.DeletePortfolioAction(context.Background(), portfolioAction1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, portfolioAction2)

	require.Equal(t, portfolioAction1.ID, portfolioAction2.ID)

	require.Equal(t, portfolioAction1.PortfolioID, portfolioAction2.PortfolioID)

	require.Equal(t, portfolioAction1.ActionID, portfolioAction2.ActionID)

	require.Equal(t, portfolioAction1.PlayerID, portfolioAction2.PlayerID)

	require.Equal(t, portfolioAction1.PurchasePrice, portfolioAction2.PurchasePrice)

	require.Equal(t, portfolioAction1.Quantity, portfolioAction2.Quantity)

	require.WithinDuration(t, portfolioAction1.CreatedAt, portfolioAction2.CreatedAt, time.Second)

	portfolioAction3, err := testQueries.GetPortfolioActionById(context.Background(), portfolioAction1.ID)
	require.Error(t, err)
	require.Empty(t, portfolioAction3)
}


func TestGetPortfolioActionByAction_id(t *testing.T){
	portfolioAction1 := createRandomPortfolioAction(t)
	portfolioAction2, err := testQueries.GetPortfolioActionByAction_id(context.Background(), portfolioAction1.ActionID)
	require.NoError(t, err)
	require.NotEmpty(t, portfolioAction2)

	require.NotEmpty(t, portfolioAction2)
}

func TestGetPortfolioActionById(t *testing.T){
	portfolioAction1 := createRandomPortfolioAction(t)
	portfolioAction2, err := testQueries.GetPortfolioActionById(context.Background(), portfolioAction1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, portfolioAction2)

	require.NotEmpty(t, portfolioAction2)
}

func TestGetPortfolioActionByPlayer_id(t *testing.T){
	portfolioAction1 := createRandomPortfolioAction(t)
	portfolioAction2, err := testQueries.GetPortfolioActionByPlayer_id(context.Background(), portfolioAction1.PlayerID)
	require.NoError(t, err)
	require.NotEmpty(t, portfolioAction2)

	require.NotEmpty(t, portfolioAction2)
}

func TestGetPortfolioActionByPortfolio_id(t *testing.T){
	portfolioAction1 := createRandomPortfolioAction(t)
	portfolioAction2, err := testQueries.GetPortfolioActionByPortfolio_id(context.Background(), portfolioAction1.PortfolioID)
	require.NoError(t, err)
	require.NotEmpty(t, portfolioAction2)

	require.NotEmpty(t, portfolioAction2)
}

func TestListPortfolioActions(t *testing.T){
	for i := 0; i < 10; i++ {
		createRandomPortfolioAction(t)
	}

	

	portfolioActions, err := testQueries.ListPortfolioActions(context.Background(), ListPortfolioActionsParams{
		Limit:  5,
		Offset: 5,
	})

	require.NoError(t, err)

	require.Len(t, portfolioActions, 5)

	for _, portfolioAction := range portfolioActions {
		require.NotEmpty(t, portfolioAction)
	}
}

func TestUpdatePortfolioAction(t *testing.T){
	portfolioAction1 := createRandomPortfolioAction(t)
	
	errs := make(chan error)

	for i := 0; i < 10; i++ {
		go func() {
			portfolioAction2, err := testQueries.UpdatePortfolioAction(context.Background(), UpdatePortfolioActionParams{
				ID: portfolioAction1.ID,
				PortfolioID: portfolioAction1.PortfolioID,
				ActionID: portfolioAction1.ActionID,
				PlayerID: portfolioAction1.PlayerID,
				PurchasePrice: portfolioAction1.PurchasePrice,
				Quantity: portfolioAction1.Quantity,
			})
			errs <- err
			require.NotEmpty(t, portfolioAction2)
		}()
	}

	for i := 0; i < 10; i++ {
		err := <- errs
		require.NoError(t, err)
	}

	portfolioAction3, err := testQueries.GetPortfolioActionById(context.Background(), portfolioAction1.ID)

	require.NoError(t, err)

	require.NotEmpty(t, portfolioAction3)

	require.Equal(t, portfolioAction1.ID, portfolioAction3.ID)

	require.Equal(t, portfolioAction1.PortfolioID, portfolioAction3.PortfolioID)

	require.Equal(t, portfolioAction1.ActionID, portfolioAction3.ActionID)

	require.Equal(t, portfolioAction1.PlayerID, portfolioAction3.PlayerID)

}