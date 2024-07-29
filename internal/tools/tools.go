package tools

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
	Logr(AnyToStr(list...))
}
func LogErrorWithStack(list ...any) {
	b := strings.Builder{}
	b.WriteString(AnyToStr(list...))
	b.WriteString("Stacktrace:\n" + string(debug.Stack()))
	Logr(b.String())
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

func ReqError(r *http.Request, err error) error {
	return fmt.Errorf(" %s\n%v", ReqToStr(r), err)

}

func LogRequest(r *http.Request, structs ...any) {
	Logg(ReqToStr(r, structs...))
}

func ReqToStr(r *http.Request, structs ...any) string {
	b := strings.Builder{}
	b.WriteString("\n❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆\n")
	b.WriteString("❆ Req: " + r.URL.Path)
	if len(r.URL.RawQuery) > 0 {
		b.WriteString("\n❆ Query: " + r.URL.RawQuery)
	}
	if len(r.Header) > 0 {
		b.WriteString("\n❆ Headers: {")
		for k, v := range r.Header {
			b.WriteString(fmt.Sprintf("\n❆    %v: %v", k, v))
		}
		b.WriteString("\n❆ }")
	}
	var bodyBytes []byte
	var err error
	bodyBytes, err = io.ReadAll(r.Body)
	if err != nil {
		b.WriteString(fmt.Sprintf("❆ Body reading error: %v", err))
		return b.String()
	}
	if len(bodyBytes) > 0 {
		var prettyJSON bytes.Buffer
		if err = json.Indent(&prettyJSON, bodyBytes, "❆ ", "  "); err != nil {
			b.WriteString(fmt.Sprintf("❆ JSON parse error: %v", err))
			return b.String()
		}
		b.WriteString("\n❆ Body: ")
		b.WriteString(string(prettyJSON.String()))
	} else {
		if len(structs) == 0 {
			b.WriteString("\n❆ Body: No Body Supplied")
		}
	}
	for i := range structs {
		structJSON, _ := json.MarshalIndent(structs[i], "❆ ", "  ")
		b.WriteString("\n❆ Struct: " + string(structJSON))
	}
	b.WriteString("\n❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆ ❆\n")
	return b.String()
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

func Success(c *gin.Context, data ...any) {
	b := strings.Builder{}
	if len(data) == 0 {
		b.WriteString("Success")
	} else {
		for i := range len(data) {
			b.WriteString(fmt.Sprintf("%v", data[i]))
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": b.String()})
}

func AbortWithErr500(c *gin.Context, messages ...any) {
	AbortWithErr(c, http.StatusInternalServerError, messages...)

}

func AbortOnPanic(c *gin.Context, messages ...any) {
	if r := recover(); r != nil {
		LogErrorWithStack(r)
		AbortWithErr500(c, messages...)
	}

}

func AbortWithErr422(c *gin.Context, messages ...any) {
	AbortWithErr(c, http.StatusUnprocessableEntity, messages...)

}
func AbortWithErr404(c *gin.Context, messages ...any) {
	AbortWithErr(c, http.StatusNotFound, messages...)

}
