package car_body_types

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
		var req carBodyTypeReq
		var err error

		req.query = c.Query("query")

		brandIdParam := c.Query("brandID")
		if req.brandID, err = strconv.Atoi(brandIdParam); err != nil {
			tools.Loge(fmt.Sprint(err))
			c.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": "brandID is required"},
			)
			return
		}

		genIdParam := c.Query("genID")
		if req.genID, err = strconv.Atoi(genIdParam); err != nil {
			tools.Loge(fmt.Sprint(err))
			c.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": "genID is required"},
			)
			return
		}

		limitParam := c.Query("limit")

		if req.limit, err = strconv.Atoi(limitParam); err != nil {
			req.limit = 20
		}

		defer tools.AbortOnPanic(c)

		list, err := searchInDB(dictDB, req)

		if err != nil {
			tools.AbortWithErr500(c, err)
			return
		}

		c.IndentedJSON(http.StatusOK, list)

	}

	return gin.HandlerFunc(fn)
}
