package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTxBuy(t *testing.T){
	store := NewStoreDB(testDb)

	actionRandom := createRandomAction(t)

	playerRandom := createRandomPlayer(t)

	buy, err := store.BuyTx(context.Background(), BuyTxParams{
		ActionIdBuy: actionRandom.ID,
		ProfileId: playerRandom.IDPlayer,
		NumberStock: 10,
		Limit: "100",
	})

	require.NoError(t, err)

	require.NotEmpty(t, buy)
}