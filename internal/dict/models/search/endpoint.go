package car_models

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	tools "zap/internal/_shared"

	"github.com/gin-gonic/gin"
)

func Search(dictDB *sql.DB) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		defer tools.AbortOnPanic(c)

		var req carModelReq
		var err error

		brandIdParam := c.Query("brandID")
		req.brandID, err = strconv.Atoi(brandIdParam)
		if err != nil {
			tools.Loge(fmt.Sprint(err))
			c.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": "markID is required"},
			)
			return
		}

		req.query = c.Query("query")

		limitParam := c.Query("limit")
		if req.limit, err = strconv.Atoi(limitParam); err != nil {
			req.limit = 20
		}

		list, err := search(dictDB, req)
		if err != nil {
			tools.AbortWithErr500(c, err)
			return
		}

		c.IndentedJSON(http.StatusOK, list)

	}

	return gin.HandlerFunc(fn)
}
