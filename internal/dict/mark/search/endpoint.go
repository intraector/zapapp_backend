package car_mark

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/theritikchoure/logx"

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

		list, err := searchInDB(dictDB, query, limit)
		if err != nil {
			logx.Log(query, logx.FGRED, logx.BGBLACK)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}

		c.IndentedJSON(http.StatusOK, list)

	}

	return gin.HandlerFunc(fn)
}
