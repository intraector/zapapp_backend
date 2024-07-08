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
	h.Router.GET("/years", h.Years())
}
