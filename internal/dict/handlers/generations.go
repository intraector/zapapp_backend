package dict_handlers

import (
	"fmt"
	"net/http"

	"zap/internal/dict/dict_model"
	tools "zap/internal/tools"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
)

func (h *Handlers) Generations() gin.HandlerFunc {

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
			tools.Logrb(fmt.Sprint(err))
			c.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": "brandID is required"},
			)
			return
		}

		list, err := h.Repo.Generations(req)
		if err != nil {
			tools.AbortWithErr500(c)
			tools.LogErrorWithStack(err, req)
			return
		}

		c.IndentedJSON(http.StatusOK, list)

	}

	return gin.HandlerFunc(fn)
}
