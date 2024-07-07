package zaps_create

import (
	"database/sql"
	"fmt"
	"net/http"
	tools "zap/internal/_shared"

	"github.com/gin-gonic/gin"
)

type Car struct {
	BrandID       int    `json:"brandID" binding:"required"`
	BrandLabel    string `json:"brandLabel" binding:"required"`
	ModelID       int    `json:"modelID" binding:"required"`
	ModelLabel    string `json:"modelLabel" binding:"required"`
	GenID         int    `json:"genID" binding:"required"`
	GenLabel      string `json:"genLabel" binding:"required"`
	BodyTypeID    int    `json:"bodyTypeID" binding:"required"`
	BodyTypeLabel string `json:"bodyTypeLabel" binding:"required"`
	ModID         int    `json:"modID" binding:"required"`
	ModLabel      string `json:"modLabel" binding:"required"`
	YearID        int    `json:"yearID" binding:"required"`
	YearValue     int    `json:"year" binding:"required"`
	VinCode       string `json:"vinCode"`
	VinImage      string `json:"vinImage"`
	Comment       string `json:"comment"`
}

func Create(db *sql.DB) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		defer tools.AbortOnPanic(c)
		// tools.LogRequest(c)

		var err error
		car := Car{}

		if err = c.ShouldBind(&car); err != nil {
			tools.Loge(fmt.Sprint(err))
			c.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": fmt.Sprint(err)},
			)
			return
		}

		if car.VinCode == "" && car.VinImage == "" {
			c.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": "Either vinCode or vinImage is required"},
			)
			return
		}

		err = create(db, car)

		if err != nil {
			tools.AbortWithErr500(c, err)
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"message": "success"})

	}

	return gin.HandlerFunc(fn)
}
