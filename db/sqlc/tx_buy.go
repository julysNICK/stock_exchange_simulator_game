package db

import (
	"context"
	"strconv"
)

type BuyTxParams struct {
	ActionIdBuy int64  `json:"action_id_buy"`
	ProfileId   int64  `json:"profile_id"`
	NumberStock int32  `json:"number_stock"`
	Limit       string `json:"limit"`
	Status      string `json:"status,omitempty"`
}

type BuyTxResult struct {
	Action          Action          `json:"action"`
	Player          Player          `json:"profile"`
	Buy             Buy             `json:"buy omitempty"`
	PortfolioAction PortfolioAction `json:"portfolio_action"`
	Portfolio       Portfolio       `json:"portfolio"`
}

func (store *SQLStore) BuyTx(ctx context.Context, arg BuyTxParams) (BuyTxResult, error) {
	var result BuyTxResult

	err := store.execTx(ctx, func(q *Queries) error {

		var err error

		result.Player, err = q.GetPlayerById(ctx, arg.ProfileId)

		if err != nil {
			return err
		}

		result.Action, err = q.GetActionById(ctx, arg.ActionIdBuy)

		if err != nil {
			return err
		}

		convPlayerCash, err := strconv.ParseInt(result.Player.Cash, 10, 64)
		if err != nil {
			return err
		}

		convLimitStock, err := strconv.ParseInt(arg.Limit, 10, 64)

		if err != nil {
			return err
		}

		convActionCurrentValue, err := strconv.ParseInt(result.Action.CurrentValue, 10, 64)

		if err != nil {
			return err
		}

		if convPlayerCash < convActionCurrentValue {
			return err
		}

		if convLimitStock != 0 {
			result.Buy, err = q.CreateBuy(ctx, CreateBuyParams{
				ActionIDBuy:  arg.ActionIdBuy,
				ProfileID:    arg.ProfileId,
				NumberStocks: arg.NumberStock,
				Limit:        arg.Limit,
			})

			if err != nil {
				return err
			}

			_, _ = q.CreatePurchaseSchedule(ctx, CreatePurchaseScheduleParams{
				BuyId:           arg.ActionIdBuy,
				Stage:           arg.Status,
				CreatedOrderBuy: result.Action.CreatedAt,
			})

			return nil

		} else {
			result.Buy, err = q.CreateBuy(ctx, CreateBuyParams{
				ActionIDBuy:  arg.ActionIdBuy,
				ProfileID:    arg.ProfileId,
				NumberStocks: arg.NumberStock,
				Limit:        arg.Limit,
				Status:       "completed",
			})
			if err != nil {
				return err
			}

			result.Portfolio, err = q.GetPortfolioByPlayerId(ctx, arg.ProfileId)

			if err != nil {
				return err
			}

			result.PortfolioAction, err = q.CreatePortfolioAction(ctx, CreatePortfolioActionParams{
				PortfolioID:   result.Portfolio.ID,
				ActionID:      result.Action.ID,
				Quantity:      arg.NumberStock,
				PurchasePrice: result.Action.CurrentValue,
			})

			if err != nil {
				return err
			}
		}

		return nil
	})

	return result, err

}

type BuyUpdateTxParams struct {
	Status    string `json:"status,omitempty"`
	IdProfile int64  `json:"id_profile"`
}

type BuyUpdateTxResult struct {
	Buy              Buy              `json:"buy omitempty"`
	PurchaseSchedule PurchaseSchedule `json:"purchase_schedule"`
	PortfolioAction  PortfolioAction  `json:"portfolio_action"`
	Portfolio        Portfolio        `json:"portfolio"`
}

func (store *SQLStore) BuyUpdateTx(ctx context.Context, arg BuyUpdateTxParams) (BuyUpdateTxResult, error) {
	var result BuyUpdateTxResult

	err := store.execTx(ctx, func(q *Queries) error {

		var err error

		result.Buy, err = q.GetBuyByBuyIdAndProfileId(ctx, GetBuyByBuyIdAndProfileIdParams{
			ID:        arg.IdProfile,
			ProfileID: arg.IdProfile,
		})

		if err != nil {
			return err
		}

		result.PurchaseSchedule, err = q.UpdatePurchaseSchedule(ctx, UpdatePurchaseScheduleParams{
			BuyId: result.Buy.ID,
			Stage: arg.Status,
		})

		if err != nil {
			return err
		}

		result.Buy, err = q.UpdateBuy(ctx, UpdateBuyParams{
			ID:     result.Buy.ID,
			Status: arg.Status,
		})

		if err != nil {
			return err
		}

		result.Portfolio, err = q.GetPortfolioByPlayerId(ctx, arg.IdProfile)

		if err != nil {
			return err
		}

		result.PortfolioAction, err = q.CreatePortfolioAction(ctx, CreatePortfolioActionParams{
			PortfolioID:   result.Portfolio.ID,
			ActionID:      result.Buy.ActionIDBuy,
			Quantity:      result.Buy.NumberStocks,
			PurchasePrice: "2.00",
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err

}
