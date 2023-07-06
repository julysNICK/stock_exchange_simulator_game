package db

import "context"

type BuyTxParams struct {
	ActionIdBuy int64 `json:"action_id_buy"`
	ProfileId int64 `json:"profile_id"`
	NumberStock int32 `json:"number_stock"`
	Limit string `json:"limit"`
}

type BuyTxResult struct {
	Action Action `json:"action"`
	Player Player `json:"profile"`
	Buy Buy `json:"number"`
}

func (store *SQLStore) BuyTx(ctx context.Context, arg BuyTxParams) (BuyTxResult, error){
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


		result.Buy, err = q.CreateBuy(ctx, CreateBuyParams{
			ActionIDBuy: arg.ActionIdBuy,
			ProfileID: arg.ProfileId,
			NumberStocks: arg.NumberStock,
			Limit: arg.Limit,
		})

		if err != nil {
			return err
		}


		return nil
	})

	return result, err

}