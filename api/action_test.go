package api

import (
	"testing"

	"github.com/gin-gonic/gin"
	db "github.com/julysNICK/stock_exchange_simulator_game/db/sqlc"
	"github.com/julysNICK/stock_exchange_simulator_game/util"
)

func TestServer_HandleBuyAction(t *testing.T) {
	type fields struct {
		config util.Config
		store  db.StoreDB
		router *gin.Engine
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				config: tt.fields.config,
				store:  tt.fields.store,
				router: tt.fields.router,
			}
			s.HandleBuyAction(tt.args.c)
		})
	}
}
