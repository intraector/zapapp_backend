package dict_handlers

import (
	tools "zap/internal/_shared"
	dict_model "zap/internal/dict/domain"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
)

func (h *Handlers) Years() gin.HandlerFunc {

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
			tools.AbortWithErr422(c, err)
			tools.LogErrorWithStack(err, req)
			return
		}

		if req.BrandID == 0 {
			tools.AbortWithErr422(c, "brandID is required")
			return
		}

		if req.ModelID == 0 {
			tools.AbortWithErr422(c, "modelID is required")
			return
		}

		if req.GenID == 0 {
			tools.AbortWithErr422(c, "genID is required")
			return
		}

		list, err := h.Repo.Years(req)
		if err != nil {
			tools.AbortWithErr500(c)
			tools.LogErrorWithStack(err, req)
			return
		}

		c.IndentedJSON(http.StatusOK, list)

	}

	return gin.HandlerFunc(fn)
}
