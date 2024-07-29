package auth_endpoints

import (
	"fmt"
	"strings"
	"zap/internal/auth/auth_model"
	"zap/internal/tools"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"

	"github.com/dchest/uniuri"
)

func (h *Endpoints) phoneCode() gin.HandlerFunc {

	fn := func(c *gin.Context) {
		defer tools.AbortOnPanic(c)
		// tools.LogRequest(c.Request)

		var err error
		req := auth_model.AuthCode{}

		err = schema.NewDecoder().Decode(
			&req,
			c.Request.URL.Query(),
		)

		if err != nil {
			tools.AbortWithErr422(c, err)
			tools.LogErrorWithStack(err, req)
			return
		}

		errorStr := strings.Builder{}

		if req.Account == "" {
			errorStr.WriteString("Account is required\n")
		}

		if errorStr.Len() > 0 {
			tools.AbortWithErr422(c, errorStr.String())
			return
		}

		var StdChars = []byte("0123456789")
		code := uniuri.NewLenChars(6, StdChars)

		query := fmt.Sprintf(`
		INSERT
		INTO auth_codes (account, code, created_at)
		VALUES (%s, '%s', CURRENT_TIMESTAMP);
			`,
			req.Account, code,
		)

		_, err = h.DB.Exec(c, query)
		if err != nil {
			tools.LogError(err)
			tools.AbortWithErr500(c)
			return
		}

		tools.Success(c)
	}

	return gin.HandlerFunc(fn)
}
