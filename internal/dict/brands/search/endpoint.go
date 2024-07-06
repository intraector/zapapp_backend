package car_brands

import (
	"database/sql"
	"net/http"
	"strconv"

	tools "zap/internal/_shared"

	"github.com/gin-gonic/gin"
)

func Search(dictDB *sql.DB) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		query := c.Query("query")

		limitParam := c.Query("limit")
		var limit int
		limit, err := strconv.Atoi(limitParam)
		if err != nil {
			limit = 20
		}

		defer tools.AbortOnPanic(c)

		list, err := searchInDB(dictDB, query, limit)
		if err != nil {
			tools.AbortWithErr500(c, err)
		}

		c.IndentedJSON(http.StatusOK, list)

	}

	return gin.HandlerFunc(fn)
}
