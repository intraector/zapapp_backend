package dict_handlers

import (
	dict_model "zap/internal/dict/domain"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Router *gin.RouterGroup
	Repo   dict_model.IDictRepo
}

func (h *Handlers) Init() {
	h.Router.GET("/brands", h.Brands())
	h.Router.GET("/models", h.Models())
	h.Router.GET("/generations", h.Generations())
	h.Router.GET("/body_types", h.BodyTypes())
	h.Router.GET("/modifications", h.Modifications())
	h.Router.GET("/years", h.Years())
}
