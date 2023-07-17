// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CountActions(ctx context.Context) (int64, error)
	CountBuy(ctx context.Context) (int64, error)
	CountPlayers(ctx context.Context) (int64, error)
	CountPortfolio(ctx context.Context) (int64, error)
	CountPortfolioActions(ctx context.Context) (int64, error)
	CreateAction(ctx context.Context, arg CreateActionParams) (Action, error)
	CreateBuy(ctx context.Context, arg CreateBuyParams) (Buy, error)
	CreatePlayer(ctx context.Context, arg CreatePlayerParams) (Player, error)
	CreatePortfolio(ctx context.Context, playerID int64) (Portfolio, error)
	CreatePortfolioAction(ctx context.Context, arg CreatePortfolioActionParams) (PortfolioAction, error)
	CreatePurchaseSchedule(ctx context.Context, arg CreatePurchaseScheduleParams) (PurchaseSchedule, error)
	DeleteAction(ctx context.Context, id int64) error
	DeleteBuy(ctx context.Context, id int64) (Buy, error)
	DeletePlayer(ctx context.Context, idPlayer int64) error
	DeletePortfolio(ctx context.Context, id int64) (Portfolio, error)
	DeletePortfolioAction(ctx context.Context, id int64) (PortfolioAction, error)
	GetActionById(ctx context.Context, id int64) (Action, error)
	GetActionByName(ctx context.Context, name string) (Action, error)
	GetAllActions(ctx context.Context) ([]Action, error)
	GetAllPurchaseSchedule(ctx context.Context) ([]PurchaseSchedule, error)
	GetBuyByActionId(ctx context.Context, actionIDBuy int64) ([]Buy, error)
	GetBuyByBuyIdAndProfileId(ctx context.Context, arg GetBuyByBuyIdAndProfileIdParams) (Buy, error)
	GetBuyById(ctx context.Context, id int64) (Buy, error)
	GetBuyByProfile_id(ctx context.Context, profileID int64) ([]Buy, error)
	GetPlayerByEmail(ctx context.Context, email string) (Player, error)
	GetPlayerById(ctx context.Context, idPlayer int64) (Player, error)
	GetPlayerByUsername(ctx context.Context, username sql.NullString) (Player, error)
	GetPortfolioActionByAction_id(ctx context.Context, actionID int64) ([]PortfolioAction, error)
	GetPortfolioActionById(ctx context.Context, id int64) (PortfolioAction, error)
	GetPortfolioActionByPlayer_id(ctx context.Context, playerID int64) ([]PortfolioAction, error)
	GetPortfolioActionByPortfolio_id(ctx context.Context, portfolioID int64) ([]PortfolioAction, error)
	GetPortfolioById(ctx context.Context, id int64) (Portfolio, error)
	GetPortfolioByPlayerId(ctx context.Context, playerID int64) (Portfolio, error)
	GetPurchaseScheduleByBuyId(ctx context.Context, buyid int64) ([]PurchaseSchedule, error)
	GetPurchaseScheduleById(ctx context.Context, id int64) (PurchaseSchedule, error)
	ListActions(ctx context.Context, arg ListActionsParams) ([]Action, error)
	ListBuy(ctx context.Context, arg ListBuyParams) ([]Buy, error)
	ListBuyByProfile_id(ctx context.Context, arg ListBuyByProfile_idParams) ([]Buy, error)
	ListPlayers(ctx context.Context, arg ListPlayersParams) ([]Player, error)
	ListPortfolio(ctx context.Context, arg ListPortfolioParams) ([]Portfolio, error)
	ListPortfolioActions(ctx context.Context, arg ListPortfolioActionsParams) ([]PortfolioAction, error)
	ListPortfolioActionsByPortfolio_id(ctx context.Context, arg ListPortfolioActionsByPortfolio_idParams) ([]PortfolioAction, error)
	RankPlayers(ctx context.Context, arg RankPlayersParams) ([]Player, error)
	UpdateAction(ctx context.Context, arg UpdateActionParams) (Action, error)
	UpdateBuy(ctx context.Context, arg UpdateBuyParams) (Buy, error)
	UpdatePlayer(ctx context.Context, arg UpdatePlayerParams) (Player, error)
	UpdatePortfolio(ctx context.Context, arg UpdatePortfolioParams) (Portfolio, error)
	UpdatePortfolioAction(ctx context.Context, arg UpdatePortfolioActionParams) (PortfolioAction, error)
	UpdatePurchaseSchedule(ctx context.Context, arg UpdatePurchaseScheduleParams) (PurchaseSchedule, error)
}

var _ Querier = (*Queries)(nil)
