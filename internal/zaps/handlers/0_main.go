package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Handlers struct {
	Router *gin.RouterGroup
	DB     *pgx.Conn
}

func (h *Handlers) Init() {
	h.Router.POST("/create", h.Create())
	h.Router.PUT("/update", h.Update())
}

func New(
	Router *gin.RouterGroup,
	DB *pgx.Conn,
) Handlers {
	output := Handlers{}
	output.Router = Router
	output.DB = DB
	return output
}
