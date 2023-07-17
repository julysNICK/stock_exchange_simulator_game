package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/julysNICK/stock_exchange_simulator_game/db/sqlc"
)

type ListPortfolioRequest struct {
 IdUser int64 `json:"id_user"`
}


type ListPortfolioResponse struct {
Portfolio db.Portfolio `json:"portfolio"`
ActionsInPortfolio []db.PortfolioAction `json:"ActionsInPortfolio"`
}


func (s *Server) ListPortfolio(ctx *gin.Context){

	var req ListPortfolioRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	portfolio, err := s.store.GetPortfolioByPlayerId(ctx, req.IdUser)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	portfolioAction, err := s.store.ListPortfolioActionsByPortfolio_id(ctx, db.ListPortfolioActionsByPortfolio_idParams{
		PortfolioID: portfolio.ID,
		Limit: 100,
		Offset: 0,
	})

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, ListPortfolioResponse{
		Portfolio: portfolio,
		ActionsInPortfolio: portfolioAction,
	})


}