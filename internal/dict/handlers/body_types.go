package dict_handlers

import (
	"net/http"

	tools "zap/internal/_shared"
	dict_model "zap/internal/dict/domain"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
)

func (h *Handlers) BodyTypes() gin.HandlerFunc {

	fn := func(c *gin.Context) {
		defer tools.AbortOnPanic(c)

		var err error
		req := dict_model.Req{Limit: 20}

		err = schema.NewDecoder().Decode(
			&req,
			c.Request.URL.Query(),
		)

		if err != nil {
			tools.AbortWithErr422(c, err)
			tools.LogErrorWithStack(err, req)
			return
		}

		if req.BrandID == 0 {
			tools.AbortWithErr422(c, "BrandID is required")
			return
		}

		if req.GenID == 0 {
			tools.AbortWithErr422(c, "genID is required")
			return
		}

		list, err := h.Repo.BodyTypes(req)
		if err != nil {
			tools.AbortWithErr500(c, "Blah blah")
			tools.LogErrorWithStack(err, req)
			return
		}

		c.IndentedJSON(http.StatusOK, list)

	}

	return gin.HandlerFunc(fn)
}
