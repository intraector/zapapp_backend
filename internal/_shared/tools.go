package shared

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/theritikchoure/logx"
)

// Logs a string with cyan
func Logc(value string) {
	logx.Log(value, logx.FGCYAN, logx.BGBLACK)
}

// Logs a string with yellow
func Logy(value string) {
	logx.Log(value, logx.FGYELLOW, logx.BGBLACK)
}

// Logs a string with green
func Logg(value string) {
	logx.Log(value, logx.FGGREEN, logx.BGBLACK)
}

// Logs a string with magenta
func Logm(value string) {
	logx.Log(value, logx.FGMAGENTA, logx.BGBLACK)
}

// Logs a string with blue
func Logb(value string) {
	logx.Log(value, logx.FGBLUE, logx.BGBLACK)
}

// Logs a string with red background
func Logrb(value string) {
	logx.Log(value, logx.FGBLACK, logx.BGRED)
}

// Logs a string with red
func Logr(value string) {
	logx.Log(value, logx.FGRED, logx.BGBLACK)
}

func LogError(list ...any) {
	Logrb(AnyToStr(list))
}
func LogErrorWithStack(list ...any) {
	b := strings.Builder{}
	b.WriteString(AnyToStr(list))
	b.WriteString("Stacktrace:\n" + string(debug.Stack()))
	Logrb(b.String())
}

func AnyToStr(list ...any) string {
	b := strings.Builder{}
	for i := range len(list) {
		_, isError := list[i].(error)
		if !isError {
			b.WriteString(fmt.Sprintf("%#v", list[i]) + "\n")
		}
	}
	for i := range len(list) {
		v, isError := list[i].(error)
		if isError {
			b.WriteString(fmt.Sprintln(v))
		}
	}
	return b.String()
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
			Logrb(fmt.Sprintf("JSON parse error: %v", err))
			return
		}
		Logb(string(prettyJSON.String()))
	} else {
		Logb("Body: No Body Supplied\n")
	}
}

func AbortWithErr(c *gin.Context, code int, messages ...any) {
	body := make(map[string]any)

	if l := len(messages); l == 0 {
		body["error"] = http.StatusText(code)
	} else if l == 1 {
		body["error"] = fmt.Sprintf("%+v", messages[0])
	} else {
		var output []string
		for i := range l {
			output = append(output, fmt.Sprintf("%+v", messages[i]))
		}
		body["errors"] = output

	}

	c.AbortWithStatusJSON(code, body)

}

func AbortWithErr500(c *gin.Context, messages ...any) {
	AbortWithErr(c, http.StatusInternalServerError, messages)

}

func AbortOnPanic(c *gin.Context, messages ...any) {
	if r := recover(); r != nil {
		AbortWithErr500(c, r, messages)
	}

}

func AbortWithErr422(c *gin.Context, messages ...any) {
	AbortWithErr(c, http.StatusUnprocessableEntity, messages)

}
