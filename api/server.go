package api

import (
	db "github.com/Coluding/udemy_backend/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	// Go lets us use the pointer to a struct to directly affect the fields of a struct, in other langugaes we would
	// have to maually derefeernece it --> server is a pointer
	// we also can directly acess struct methods, see above
	server.router = router

	return server

}

func (server *Server) Start(adress string) error {
	return server.router.Run(adress)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
