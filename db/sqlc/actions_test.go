package db

import (
	"context"
	"database/sql"
	"strconv"
	"testing"
	"time"

	"github.com/julysNICK/stock_exchange_simulator_game/util"
	"github.com/stretchr/testify/require"
)


func createRandomAction(t *testing.T) Action {
	var bid string = util.RandomBidAction()

	var ask string = util.RandomAskAction()

	println("bid: " + bid)
	println("ask: " + ask)

	convBid, err := strconv.ParseInt(bid, 10, 64)
	require.NoError(t, err)

	convAsk, err := strconv.ParseInt(ask, 10, 64)
	require.NoError(t, err)


	var spread string = util.RandomSpreadAction(convBid, convAsk)

	arg := CreateActionParams{
		Name: util.RandomNameAction(),
		IDActions: sql.NullInt32{
			Int32: util.RandomIDAction(),
			Valid: true,
		},

		Isin: util.RandomIsinAction(),
		Wkn: util.RandomWknAction(),
		CurrentValue: util.RandomCurrentValueAction(),
		Bid: bid + ".00",
		Ask: ask + ".00",
		Spread: spread + ".00",
		TimeOfLastRefresh: time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC),
		ChangePercentage: util.RandomChangePercentageAction(
			util.RandomINT32String(
				1,
				100,
			),
			util.RandomINT32String(
				1,100,
			),
		) + ".00",
		ChangeAbsolute: util.RandomChangeAbsoluteAction(
			util.RandomINT32String(
				1,
				100,
			),
			util.RandomINT32String(
				1,100,
			),
		) + ".00",
		Peak24h: util.RandomPeak24hAction() + ".00",
		Peak7d: util.RandomPeak7dAction() + ".00",
		Peak30d: util.RandomPeak30dAction() + ".00",
		Low24h: util.RandomLow24hAction() + ".00",
		Low7d: util.RandomLow7dAction() + ".00",
		Low30d: util.RandomLow30dAction() + ".00",
	}

	action, err := testQueries.CreateAction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, action)

	require.Equal(t, arg.Name, action.Name)
	require.Equal(t, arg.IDActions, action.IDActions)
	require.Equal(t, arg.Isin, action.Isin)
	require.Equal(t, arg.Wkn, action.Wkn)
	require.Equal(t, arg.CurrentValue, action.CurrentValue)
	require.Equal(t, arg.Bid, action.Bid)

	require.Equal(t, arg.Ask, action.Ask)

	require.Equal(t, arg.Spread, action.Spread)

	require.WithinDuration(t, arg.TimeOfLastRefresh, action.TimeOfLastRefresh, time.Second)
	require.Equal(t, arg.ChangePercentage, action.ChangePercentage)
	require.Equal(t, arg.ChangeAbsolute, action.ChangeAbsolute)
	require.Equal(t, arg.Peak24h, action.Peak24h )
	require.Equal(t, arg.Peak7d, action.Peak7d)
	require.Equal(t, arg.Peak30d, action.Peak30d)
	require.Equal(t, arg.Low24h, action.Low24h)
	require.Equal(t, arg.Low7d, action.Low7d)
	require.Equal(t, arg.Low30d, action.Low30d)
	
	return action
}

func TestCreateAction(t *testing.T) {
	createRandomAction(t)
}

func TestGetActionById(t *testing.T){
	action1 := createRandomAction(t)
	action2, err := testQueries.GetActionById(context.Background(), action1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, action2)

	require.Equal(t, action1.Name, action2.Name)
	require.Equal(t, action1.IDActions, action2.IDActions)
	require.Equal(t, action1.Isin, action2.Isin)
	require.Equal(t, action1.Wkn, action2.Wkn)
	require.Equal(t, action1.CurrentValue, action2.CurrentValue)
	require.Equal(t, action1.Bid, action2.Bid)

	require.Equal(t, action1.Ask, action2.Ask)

	require.Equal(t, action1.Spread, action2.Spread)

	require.WithinDuration(t, action1.TimeOfLastRefresh, action2.TimeOfLastRefresh, time.Second)
	require.Equal(t, action1.ChangePercentage, action2.ChangePercentage)
	require.Equal(t, action1.ChangeAbsolute, action2.ChangeAbsolute)
	require.Equal(t, action1.Peak24h, action2.Peak24h)
	require.Equal(t, action1.Peak7d, action2.Peak7d)
	require.Equal(t, action1.Peak30d, action2.Peak30d)
	require.Equal(t, action1.Low24h, action2.Low24h)
	require.Equal(t, action1.Low7d, action2.Low7d)
	require.Equal(t, action1.Low30d, action2.Low30d)

}


func TestCountAction(t *testing.T){
	for i := 0; i < 10; i++ {
		createRandomAction(t)
	}

	count, err := testQueries.CountActions(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, count)
}

func TestDelete(t *testing.T){
	action1 := createRandomAction(t)
	err := testQueries.DeleteAction(context.Background(), action1.ID)
	require.NoError(t, err)

	action2, err := testQueries.GetActionById(context.Background(), action1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, action2)
}

func TestGetActionByName(t *testing.T) {
	action1 := createRandomAction(t)
	action2, err := testQueries.GetActionByName(context.Background(), action1.Name)
	require.NoError(t, err)
	require.NotEmpty(t, action2)

	require.Equal(t, action1.Name, action2.Name)
	require.Equal(t, action1.IDActions, action2.IDActions)
	require.Equal(t, action1.Isin, action2.Isin)
	require.Equal(t, action1.Wkn, action2.Wkn)
	require.Equal(t, action1.CurrentValue, action2.CurrentValue)
	require.Equal(t, action1.Bid, action2.Bid)

	require.Equal(t, action1.Ask, action2.Ask)

	require.Equal(t, action1.Spread, action2.Spread)

	require.Equal(t, action1.TimeOfLastRefresh, action2.TimeOfLastRefresh)
	require.Equal(t, action1.ChangePercentage, action2.ChangePercentage)
	require.Equal(t, action1.ChangeAbsolute, action2.ChangeAbsolute)
	require.Equal(t, action1.Peak24h, action2.Peak24h)
	require.Equal(t, action1.Peak7d, action2.Peak7d)
	require.Equal(t, action1.Peak30d, action2.Peak30d)
	require.Equal(t, action1.Low24h, action2.Low24h)
	require.Equal(t, action1.Low7d, action2.Low7d)
	require.Equal(t, action1.Low30d, action2.Low30d)
	
}

func TestListActions(t *testing.T){
	for i := 0; i < 10; i++ {
		createRandomAction(t)
	}

	actions, err := testQueries.ListActions(context.Background(), ListActionsParams{
		Limit:  5,
		Offset: 5,
	})
	require.NoError(t, err)
	require.NotEmpty(t, actions)
	require.Len(t, actions, 5)

	for _, action := range actions {
		require.NotEmpty(t, action)
	}
}

func TestUpdateAction(t *testing.T){
	actions1 := createRandomAction(t)

	arg := UpdateActionParams{
		ID: actions1.ID,
		Name: sql.NullString{
			String: "Test",
			Valid: true,
		},
	}

	action2, err := testQueries.UpdateAction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, action2)

	require.NotEqual(t, actions1.Name, action2.Name)

	require.Equal(t, actions1.IDActions, action2.IDActions)
	require.Equal(t, actions1.Isin, action2.Isin)
	require.Equal(t, actions1.Wkn, action2.Wkn)
	require.Equal(t, actions1.CurrentValue, action2.CurrentValue)
	require.Equal(t, actions1.Bid, action2.Bid)

}



