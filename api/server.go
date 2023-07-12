package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/julysNICK/stock_exchange_simulator_game/db/sqlc"
	"github.com/julysNICK/stock_exchange_simulator_game/util"
)

type Server struct {
	config util.Config
	store  db.StoreDB
	router *gin.Engine
}

func NewServer(config util.Config, store db.StoreDB) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/game/actions/:room", server.HandleGetActions)
	router.POST("/game/actions/buy", server.HandleBuyAction)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
