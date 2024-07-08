package zap_update

import (
	"database/sql"
	"fmt"

	tools "zap/internal/_shared"
	model "zap/internal/zaps/domain"
)

type Car model.Car

func (car *Car) update(dictDB *sql.DB) error {

	query := fmt.Sprintf(`
		UPDATE zapdb.zaps 
		SET "brand_id" = %d,
			"brand_label" = %q,
			"model_id" = %d,
			"model_label" = %q,
			"gen_id" = %d,
			"gen_label" = %q,
			"body_type_id" = %d,
			"body_type_label" = %q,
			"mod_id" = %d,
			"mod_label" = %q,
			"year_id" = %d,
			"year_value" = %d,
			"vin_code" = %q,
			"vin_image" = %q,
			"comment" = %q 
			WHERE "id" = %d;
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

	tools.Logg(query)

	_, err := dictDB.Exec(query)

	if err != nil {
		return fmt.Errorf("/create\n req: %v, error: %v", car, err)
	}

	return nil
}
