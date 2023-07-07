package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/golang/mock/mockgen/model"
	"github.com/julysNICK/stock_exchange_simulator_game/api"
	db "github.com/julysNICK/stock_exchange_simulator_game/db/sqlc"
	"github.com/julysNICK/stock_exchange_simulator_game/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		fmt.Println("Error loading config file ", err.Error())
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		fmt.Println("Error connecting to database ", err.Error())
	}

	store := db.NewStoreDB(conn)

	server, err := api.NewServer(config, store)

	if err != nil {
		fmt.Println("Error creating server ", err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go server.UpdateActionsCurrentValues(ctx)

	err = server.Start(config.HTTP_SERVER_ADDRESS)

	if err != nil {
		fmt.Println("Error starting server ", err.Error())
	}

}
