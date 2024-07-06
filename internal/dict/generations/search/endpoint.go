package car_generations

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

		var req carGenerationReq
		var err error

		req.query = c.Query("query")

		brandParam := c.Query("brandID")
		if req.brandID, err = strconv.Atoi(brandParam); err != nil {
			tools.Log(fmt.Sprint(err))
			c.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": "brandID is required"},
			)
			return
		}

		limitParam := c.Query("limit")
		req.limit, err = strconv.Atoi(limitParam)
		if err != nil {
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
