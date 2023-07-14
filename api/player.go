package api

import (
	"database/sql"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	db "github.com/julysNICK/stock_exchange_simulator_game/db/sqlc"
)

type CreatePlayerRequest struct {
	Username       string `json:"username" binding:"required"`
	HashedPassword string `json:"hashed_password" binding:"required"`
	Email          string `json:"email" binding:"required"`
}

type CreatePlayerResponse struct {
	Message string            `json:"message"`
	Player  db.PlayerTxResult `json:"player"`
}

func (s *Server) CreatePlayer(ctx *gin.Context) {
	var req CreatePlayerRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	player, err := s.store.PlayerTx(ctx, db.PlayerTxParams{
		UserName:       req.Username,
		HashedPassword: req.HashedPassword,
		Email:          req.Email,
	})

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, CreatePlayerResponse{
		Message: "Player created successfully",
		Player:  player,
	})
}

type UpdatePlayerRequest struct {
	UserName          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Cash              string    `json:"cash"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
}

type UpdatePlayerUriParams struct {
	Id int64 `uri:"id" binding:"required"`
}

type UpdatePlayerResponse struct {
	Message string    `json:"message"`
	Player  db.Player `json:"player"`
}

func (s *Server) UpdatePlayer(ctx *gin.Context) {
	var req UpdatePlayerRequest
	var uri UpdatePlayerUriParams

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	playerUpdate, err := s.store.UpdatePlayer(ctx, db.UpdatePlayerParams{
		IDPlayer: uri.Id,
		Username: sql.NullString{
			String: req.UserName,
			Valid:  req.UserName != "",
		},
		HashedPassword: sql.NullString{
			String: req.HashedPassword,
			Valid:  req.HashedPassword != "",
		},
		FullName: sql.NullString{
			String: req.FullName,
			Valid:  req.FullName != "",
		},
		Cash: sql.NullString{
			String: req.Cash,
			Valid:  req.Cash != "",
		},
		Email: sql.NullString{
			String: req.Email,
			Valid:  req.Email != "",
		},
		PasswordChangedAt: sql.NullTime{
			Time:  req.PasswordChangedAt,
			Valid: req.PasswordChangedAt != time.Time{},
		},
	})

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, UpdatePlayerResponse{
		Message: "Player updated successfully",
		Player:  playerUpdate,
	})

}

type GetPlayersRankingResponse struct {
	Message string        `json:"message"`
	Players []db.Player `json:"players"`
}

type GetPlayersRankingFormParams struct {
	PageId int32 `form:"page_id" binding:"required"`
	Limit  int32 `form:"limit" binding:"required"`
}

func (s *Server) GetPlayersRanking(ctx *gin.Context) {
	var form GetPlayersRankingFormParams

	if err := ctx.ShouldBindQuery(&form); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	players, err := s.store.RankPlayers(ctx, db.RankPlayersParams{
		Limit:  form.Limit,
		Offset: form.PageId * form.Limit,
	})

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, GetPlayersRankingResponse{
		Message: "Players ranking",
		Players: players,
	})
}
	