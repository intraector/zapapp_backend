package shared

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theritikchoure/logx"
)

func Loge(value string) {
	logx.Log(value, logx.FGBLACK, logx.BGRED)
}
func Logc(value string) {
	logx.Log(value, logx.FGCYAN, logx.BGBLACK)
}
func Logy(value string) {
	logx.Log(value, logx.FGYELLOW, logx.BGBLACK)
}
func Logg(value string) {
	logx.Log(value, logx.FGGREEN, logx.BGBLACK)
}
func Logm(value string) {
	logx.Log(value, logx.FGMAGENTA, logx.BGBLACK)
}
func Logb(value string) {
	logx.Log(value, logx.FGBLUE, logx.BGBLACK)
}
func Logr(value string) {
	logx.Log(value, logx.FGRED, logx.BGBLACK)
}

func AbortOnPanic(c *gin.Context) {
	if r := recover(); r != nil {
		Loge(fmt.Sprint(r))
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "Internal Server Error"},
		)
	}

}
func AbortWithErr500(c *gin.Context, err error) {
	Loge(fmt.Sprint(err))
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		gin.H{"error": "Internal Server Error"},
	)

}

func LogRequest(c *gin.Context) {
	Logb("\n" + c.Request.URL.Path)

	c.Request.URL.Query()
	Logb("Query: " + c.Request.URL.RawQuery)
	Logb("Headers:")
	for k, v := range c.Request.Header {
		Logb(fmt.Sprintf("    %v: %v", k, v))
	}
	var bodyBytes []byte
	var err error
	body := c.Request.Body
	if body != nil {
		bodyBytes, err = io.ReadAll(body)
		if err != nil {
			fmt.Printf("Body reading error: %v", err)
			return
		}
	}

	if len(bodyBytes) > 0 {
		var prettyJSON bytes.Buffer
		if err = json.Indent(&prettyJSON, bodyBytes, "", "  "); err != nil {
			Loge(fmt.Sprintf("JSON parse error: %v", err))
			return
		}
		Logb(string(prettyJSON.String()))
	} else {
		Logb("Body: No Body Supplied\n")
	}
}
