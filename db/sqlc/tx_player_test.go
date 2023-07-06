package db

import (
	"context"
	"testing"

	"github.com/julysNICK/stock_exchange_simulator_game/util"
	"github.com/stretchr/testify/require"
)

func TestTxUser(t *testing.T){
	store := NewStoreDB(testDb)



	user, err := store.PlayerTx(context.Background(), PlayerTxParams{
			UserName: util.RandomNamePlayer(),
			HashedPassword: util.RandomPasswordPlayer(),
			FullName: util.RandomString(10),
			Cash: util.RandomCashPlayer(),
			Email: util.RandomEmailPlayer(),
	})

	require.NoError(t, err)

	require.NotEmpty(t, user)
}