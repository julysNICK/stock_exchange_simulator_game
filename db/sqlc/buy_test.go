package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomBuy(t *testing.T) Buy {

	action := createRandomAction(t)
	profile := createRandomPlayer(t)

	arg := CreateBuyParams{
		ActionIDBuy:  action.ID,
		ProfileID:    profile.IDPlayer,
		NumberStocks: 1,
		LimitBuy:     "100.00",
	}

	buy, err := testQueries.CreateBuy(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, buy)

	require.Equal(t, arg.ActionIDBuy, buy.ActionIDBuy)
	require.Equal(t, arg.ProfileID, buy.ProfileID)
	require.Equal(t, arg.NumberStocks, buy.NumberStocks)
	require.Equal(t, arg.LimitBuy, buy.LimitBuy)

	return buy

}

func createRandomBuyFix(t *testing.T) Buy {

	action := createRandomAction(t)

	arg := CreateBuyParams{
		ActionIDBuy:  action.ID,
		ProfileID:    1,
		NumberStocks: 1,
		LimitBuy:     "100.00",
	}

	buy, err := testQueries.CreateBuy(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, buy)

	require.Equal(t, arg.ActionIDBuy, buy.ActionIDBuy)
	require.Equal(t, arg.ProfileID, buy.ProfileID)
	require.Equal(t, arg.NumberStocks, buy.NumberStocks)
	require.Equal(t, arg.LimitBuy, buy.LimitBuy)

	return buy

}

func TestCreateBuy(t *testing.T) {
	createRandomBuy(t)
}

func TestCountBuy(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomBuy(t)
	}

	count, err := testQueries.CountBuy(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, count)
}

func TestDeleteBuy(t *testing.T) {
	buy1 := createRandomBuy(t)
	buyDeleted, err := testQueries.DeleteBuy(context.Background(), buy1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, buyDeleted)
	require.Equal(t, buy1.ID, buyDeleted.ID)
	require.Equal(t, buy1.ActionIDBuy, buyDeleted.ActionIDBuy)
	require.Equal(t, buy1.ProfileID, buyDeleted.ProfileID)
	require.Equal(t, buy1.NumberStocks, buyDeleted.NumberStocks)
}

func TestGetBuyById(t *testing.T) {
	buy1 := createRandomBuy(t)
	buyGet, err := testQueries.GetBuyById(context.Background(), buy1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, buyGet)
	require.Equal(t, buy1.ID, buyGet.ID)
	require.Equal(t, buy1.ActionIDBuy, buyGet.ActionIDBuy)
	require.Equal(t, buy1.ProfileID, buyGet.ProfileID)
	require.Equal(t, buy1.NumberStocks, buyGet.NumberStocks)
}

func TestGetBuyByProfileId(t *testing.T) {
	buy1 := createRandomBuy(t)
	buyGet, err := testQueries.GetBuyByProfile_id(context.Background(), buy1.ProfileID)
	require.NoError(t, err)
	require.NotEmpty(t, buyGet)

}

func TestListBuy(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomBuy(t)
	}

	arg := ListBuyParams{
		Limit:  5,
		Offset: 5,
	}

	buys, err := testQueries.ListBuy(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, buys, int(arg.Limit))
}

// func TestListBuyByProfile_id(t *testing.T){

// 	for i := 0; i < 10; i++ {
// 		createRandomBuyFix(t)
// 	}

// 	arg := ListBuyByProfile_idParams{
// 		ProfileID:  1,
// 		Limit:  5,
// 		Offset: 5,
// 	}

// 	buys, err := testQueries.ListBuyByProfile_id(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.Len(t, buys, int(arg.Limit))
// }
