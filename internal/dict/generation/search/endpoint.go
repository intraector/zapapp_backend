package car_generation

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
		query := c.Query("query")

		model := c.Query("modelID")
		modelID, err := strconv.Atoi(model)
		if err != nil {
			logx.Log(fmt.Sprint(err), logx.FGRED, logx.BGBLACK)
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "modelID is required"})
			return
		}

		limitParam := c.Query("limit")
		var limit int
		limit, err = strconv.Atoi(limitParam)
		if err != nil {
			limit = 20
		}

		list, err := searchInDB(dictDB, modelID, query, limit)

		if err != nil {
			logx.Log(fmt.Sprint(err), logx.FGRED, logx.BGBLACK)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		c.IndentedJSON(http.StatusOK, list)

	}

	return gin.HandlerFunc(fn)
}
