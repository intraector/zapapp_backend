package endpoints

import (
	"fmt"
	"net/http"
	tools "zap/internal/_shared"
	model "zap/internal/zaps/domain"

	"github.com/gin-gonic/gin"
)

func (endpoints *ZapsEndpoints) Update() gin.HandlerFunc {

	fn := func(c *gin.Context) {
		defer tools.AbortOnPanic(c)
		// tools.LogRequest(c)

		var err error
		car := model.Car{}

		if err = c.ShouldBind(&car); err != nil {
			tools.Loge(fmt.Sprint(err))
			c.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": fmt.Sprint(err)},
			)
			return
		}

		if car.ID == 0 {
			c.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": "ID is required"},
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

		err = endpoints.Repo.Update(&car)

		if err != nil {
			tools.AbortWithErr500(c, err)
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"message": "success"})

	}

	return gin.HandlerFunc(fn)
}
