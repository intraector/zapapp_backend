package car_model

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/theritikchoure/logx"

	"github.com/gin-gonic/gin"
)

func Search(dictDB *sql.DB) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		queryParam := c.Query("query")
		markParam := c.Query("markID")

		limitParam := c.Query("limit")
		var limit int
		limit, err := strconv.Atoi(limitParam)
		if err != nil {
			limit = 20
		}

		markID, err := strconv.Atoi(markParam)
		if err != nil {
			logx.Log(fmt.Sprint(err), logx.FGRED, logx.BGBLACK)
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "markID is required"})
			return
		}

		list, err := searchInDB(dictDB, markID, queryParam, limit)
		if err != nil {
			logx.Log(fmt.Sprint(err), logx.FGRED, logx.BGBLACK)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		c.IndentedJSON(http.StatusOK, list)

	}

	return gin.HandlerFunc(fn)
}
