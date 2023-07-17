package api

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	db "github.com/julysNICK/stock_exchange_simulator_game/db/sqlc"
)

type ActionInfo struct {
	ID          int64  `json:"id"`
	StockID     int64  `json:"stock_id"`
	StockName   string `json:"stock_name"`
	StockSymbol string `json:"stock_symbol"`
	Price       int64  `json:"price"`
	Amount      int64  `json:"amount"`
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	rooms = make(map[string]map[*websocket.Conn]bool)

	roomMux sync.Mutex
)

type GetAllActionsResponse struct {
	Message string      `json:"message"`
	Actions []db.Action `json:"actions"`
}

func (s *Server) HandleGetActions(c *gin.Context) {
	room := c.Param("room")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	roomMux.Lock()

	if _, ok := rooms[room]; !ok {
		rooms[room] = make(map[*websocket.Conn]bool)
	}

	rooms[room][conn] = true

	roomMux.Unlock()

	actions, err := s.store.GetAllActions(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = conn.WriteJSON(GetAllActionsResponse{
		Message: "success",
		Actions: actions,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	go readMessage(room, conn)

}

func readMessage(room string, conn *websocket.Conn) {
	defer func() {
		conn.Close()
		roomMux.Lock()
		delete(rooms[room], conn)
		roomMux.Unlock()
	}()
	for {
		var msg ActionInfo
		err := conn.ReadJSON(&msg)

		if err != nil {
			return
		}

		go broadcastMessage(room, msg)

	}
}

func (s *Server) UpdateActionsCurrentValues(
	c context.Context,
) {

	for {
		time.Sleep(100 * time.Minute)

		actions, err := s.store.GetAllActions(c)

		if err != nil {
			return
		}

		for _, action := range actions {

			_, err := s.store.UpdateAction(c, db.UpdateActionParams{
				ID:           action.ID,
				CurrentValue: sql.NullString{String: "2.00", Valid: true},
			})

			if err != nil {

				return
			}

		}

		actions, err = s.store.GetAllActions(c)

		if err != nil {
			return
		}

		roomMux.Lock()

		for conn := range rooms["1"] {

			err := conn.WriteJSON(GetAllActionsResponse{
				Message: "success update",
				Actions: actions,
			})

			if err != nil {
				conn.Close()
				delete(rooms["1"], conn)
			}

		}

		roomMux.Unlock()

	}
}

func (s *Server) CronCheckPurchaseSchedule(
	c context.Context,
) {

	fmt.Println("CronCheckPurchaseSchedule")

	for {
		fmt.Println("CronCheckPurchaseSchedule for")

		time.Sleep(1 * time.Minute)

		purchaseSchedules, err := s.store.GetAllPurchaseSchedule(c)

		fmt.Println("purchaseSchedules", purchaseSchedules)

		if err != nil {
			fmt.Println("err", err)
			return
		}

		for _, purchaseSchedule := range purchaseSchedules {

			fmt.Println("purchaseSchedule", purchaseSchedule)

			buy, err := s.store.GetBuyById(c, purchaseSchedule.BuyId)
			fmt.Println("buy", buy)
			if err != nil {
				fmt.Println("err", err)
				return
			}

			if buy.Status == "pending" {
				fmt.Println("buy.Status == pending")
				action, err := s.store.GetActionById(c, buy.ActionIDBuy)

				if err != nil {
					return
				}

				convPrice, err := strconv.ParseFloat(action.CurrentValue, 64)

				if err != nil {
					return
				}

				convLimit, err := strconv.ParseFloat(buy.LimitBuy, 64)

				if err != nil {
					return
				}

				fmt.Println("convPrice", convPrice)
				fmt.Println("convLimit", convLimit)

				if convPrice <= convLimit {
					fmt.Println("convPrice <= convLimit")
					_, err := s.store.BuyUpdateTx(c, db.BuyUpdateTxParams{
						IdPurchaseSchedule: purchaseSchedule.ID,
						IdProfile:          buy.ProfileID,
						Status:             "success",
					})

					if err != nil {
						fmt.Println("err", err)
						return
					}
				}

				return

			}
			return
		}

	}

}

func broadcastMessage(room string, msg ActionInfo) {
	roomMux.Lock()
	for conn := range rooms[room] {
		err := conn.WriteJSON(msg)
		if err != nil {
			conn.Close()
			delete(rooms[room], conn)
		}
	}

	roomMux.Unlock()
}

type BuyActionRequest struct {
	StockID    int64  `json:"stock_id"`
	ProfileID  int64  `json:"profile_id"`
	Amount     int32  `json:"amount"`
	LimitPrice string `json:"limit_price"`
}

type BuyActionResponse struct {
	Message string `json:"message"`
}

func (s *Server) HandleBuyAction(c *gin.Context) {
	var req BuyActionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	_, err := s.store.BuyTx(c, db.BuyTxParams{
		ActionIdBuy: req.StockID,
		ProfileId:   req.ProfileID,
		NumberStock: req.Amount,
		Limit:       req.LimitPrice,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, BuyActionResponse{
		Message: "success buy",
	})

	// _, err := s.store.ScheduleBuyTx(c, db.ScheduleBuyTxParams{
	// 	ActionIdBuy: req.StockID,
	// 	ProfileId:   req.ProfileID,
	// 	NumberStock: req.Amount,
	// 	Limit:       req.LimitPrice,
	// })

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})

	// 	return
	// }

	// c.JSON(http.StatusOK, BuyActionResponse{
	// 	Message: "success schedule buy",
	// })
}
