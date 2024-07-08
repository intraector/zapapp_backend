package dict_handlers

import (
	"fmt"
	"net/http"

	tools "zap/internal/_shared"
	dict_model "zap/internal/dict/domain"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
)

func (h *Handlers) Modifications() gin.HandlerFunc {

	fn := func(c *gin.Context) {
		defer tools.AbortOnPanic(c)
		// tools.LogRequest(c)

		var err error
		req := dict_model.Req{Limit: 20}

		err = schema.NewDecoder().Decode(
			&req,
			c.Request.URL.Query(),
		)

		if err != nil {
			tools.Loge(fmt.Sprint(err))
			c.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": fmt.Sprint(err)},
			)
			return
		}

		if req.BrandID == 0 {
			c.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": "brandID is required"},
			)
			return
		}

		if req.BodyTypeID == 0 {
			c.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": "bodyTypeID is required"},
			)
			return
		}

		list, err := h.Repo.Modifications(req)
		if err != nil {
			tools.AbortWithErr500(c, err)
			return
		}

		c.IndentedJSON(http.StatusOK, list)

	}

	return gin.HandlerFunc(fn)
}
