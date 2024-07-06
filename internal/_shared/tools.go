package shared

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theritikchoure/logx"
)

func Log(value string) {
	logx.Log(value, logx.FGBLACK, logx.BGRED)
}

func AbortOnPanic(c *gin.Context) {
	if r := recover(); r != nil {
		Log(fmt.Sprint(r))
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "Internal Server Error"},
		)
	}

}
func AbortWithErr500(c *gin.Context, err error) {
	Log(fmt.Sprint(err))
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		gin.H{"error": "Internal Server Error"},
	)

}
