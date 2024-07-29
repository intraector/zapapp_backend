package zap_handlers

import (
	"encoding/json"
	"fmt"
	"strings"
	"zap/internal/tools"
	"zap/internal/zaps/zap_model"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) Update() gin.HandlerFunc {

	fn := func(c *gin.Context) {
		defer tools.AbortOnPanic(c)
		// tools.LogRequest(c.Request)

		var err error
		car := zap_model.Zap{}

		err = json.NewDecoder(c.Request.Body).Decode(&car)
		if err != nil {
			tools.AbortWithErr422(c, err)
			return
		}

		errorStr := strings.Builder{}
		if car.ID == 0 {
			errorStr.WriteString("ID is required\n")
		}

		if car.VinCode == "" && car.VinImage == "" {
			errorStr.WriteString("Either vinCode or vinImage is required\n")
		}

		if errorStr.Len() > 0 {
			tools.AbortWithErr422(c, errorStr.String())
			return
		}

		query := fmt.Sprintf(`
		UPDATE zaps 
		SET brand_id = %d,
			brand_label = '%s',
			model_id = %d,
			model_label = '%s',
			gen_id = %d,
			gen_label = '%s',
			body_type_id = %d,
			body_type_label = '%s',
			mod_id = %d,
			mod_label = '%s',
			year_id = %d,
			year = %d,
			vin_code = '%s',
			vin_image = '%s',
			comment = '%s' 
			WHERE ID = %d;
			`,

			car.BrandID, car.BrandLabel,
			car.ModelID, car.ModelLabel,
			car.GenID, car.GenLabel,
			car.BodyTypeID, car.BodyTypeLabel,
			car.ModID, car.ModLabel,
			car.YearID, car.YearValue,
			car.VinCode, car.VinImage,
			car.Comment,
			car.ID,
		)

		tag, err := h.DB.Exec(c, query)
		if err != nil {
			tools.LogError(err)
			tools.AbortWithErr500(c)
			return
		}
		if tag.RowsAffected() == 0 {
			tools.AbortWithErr404(c)
			return
		}

		tools.Success(c)
	}

	return gin.HandlerFunc(fn)
}
