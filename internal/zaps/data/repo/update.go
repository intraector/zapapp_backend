package repo

import (
	"fmt"
	tools "zap/internal/_shared"
	model "zap/internal/zaps/domain"
)

func (r *ZapsRepo) Update(Car *model.Car) error {

	query := fmt.Sprintf(`
		UPDATE zapdb.zaps 
		SET brand_id = %d,
			brand_label = %q,
			model_id = %d,
			model_label = %q,
			gen_id = %d,
			gen_label = %q,
			body_type_id = %d,
			body_type_label = %q,
			mod_id = %d,
			mod_label = %q,
			year_id = %d,
			year = %d,
			vin_code = %q,
			vin_image = %q,
			comment = %q 
			WHERE ID = %d;
			`,

		Car.BrandID, Car.BrandLabel,
		Car.ModelID, Car.ModelLabel,
		Car.GenID, Car.GenLabel,
		Car.BodyTypeID, Car.BodyTypeLabel,
		Car.ModID, Car.ModLabel,
		Car.YearID, Car.YearValue,
		Car.VinCode, Car.VinImage,
		Car.Comment,
		Car.ID,
	)

	fmt.Println(query)

	tools.Logg(query)

	_, err := r.DB.Exec(query)

	if err != nil {
		return fmt.Errorf("/create\n req: %v, error: %v", Car, err)
	}

	return nil
}
