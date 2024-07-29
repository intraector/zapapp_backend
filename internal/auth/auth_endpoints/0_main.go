package auth_endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Endpoints struct {
	Router *gin.RouterGroup
	DB     *pgx.Conn
}

func (h *Endpoints) Init() {
	h.Router.GET("/code", h.phoneCode())
	h.Router.PUT("/update", h.Update())
}

func New(
	Router *gin.RouterGroup,
	DB *pgx.Conn,
) Endpoints {
	output := Endpoints{}
	output.Router = Router
	output.DB = DB
	return output
}
