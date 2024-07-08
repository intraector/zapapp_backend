package handlers

import (
	model "zap/internal/zaps/domain"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Router *gin.RouterGroup
	Repo   model.IZapsRepo
}

func (h *Handlers) Init() {
	h.Router.POST("/create", h.Create())
	h.Router.PUT("/update", h.Update())
}
