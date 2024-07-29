package handlers

import (
	"fmt"
	"strings"
	tools "zap/internal/_shared"
	model "zap/internal/zaps/domain"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) Create() gin.HandlerFunc {

	fn := func(c *gin.Context) {
		defer tools.AbortOnPanic(c)
		// tools.LogRequest(c.Request)

		var err error
		car := model.Car{}

		err = c.ShouldBind(&car)
		if err != nil {
			tools.AbortWithErr422(c, err)
			return
		}

		errorStr := strings.Builder{}

		if car.VinCode == "" && car.VinImage == "" {
			errorStr.WriteString("Either vinCode or vinImage is required\n")
		}

		if errorStr.Len() > 0 {
			tools.AbortWithErr422(c, errorStr.String())
			return
		}

		query := fmt.Sprintf(`
		INSERT
		INTO zaps (
			ID,
			brand_id, brand_label,
			model_id, model_label,
			gen_id, gen_label,
			body_type_id, body_type_label,
			mod_id, mod_label,
			year_id, year,
			vin_code, vin_image,
			comment
		)
		VALUES (
			DEFAULT,
			%d, '%s',
			%d, '%s',
			%d, '%s',
			%d, '%s',
			%d, '%s',
			%d, %d,
			'%s', '%s',
			'%s'
			);
			`,
			car.BrandID, car.BrandLabel,
			car.ModelID, car.ModelLabel,
			car.GenID, car.GenLabel,
			car.BodyTypeID, car.BodyTypeLabel,
			car.ModID, car.ModLabel,
			car.YearID, car.YearValue,
			car.VinCode, car.VinImage,
			car.Comment,
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
