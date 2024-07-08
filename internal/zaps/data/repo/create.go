package repo

import (
	"fmt"
	tools "zap/internal/_shared"
	model "zap/internal/zaps/domain"
)

func (r *ZapsRepo) Create(Car *model.Car) error {

	query := fmt.Sprintf(`
		INSERT INTO zapdb.zaps 
		VALUES (
			%d,
			%d, %q,
			%d, %q,
			%d, %q,
			%d, %q,
			%d, %q,
			%d, %d,
			%q, %q,
			%q
			);
			`,
		0,
		Car.BrandID, Car.BrandLabel,
		Car.ModelID, Car.ModelLabel,
		Car.GenID, Car.GenLabel,
		Car.BodyTypeID, Car.BodyTypeLabel,
		Car.ModID, Car.ModLabel,
		Car.YearID, Car.YearValue,
		Car.VinCode, Car.VinImage,
		Car.Comment,
	)

	tools.Logg(query)

	_, err := r.DB.Exec(query)

	if err != nil {
		return fmt.Errorf("/create\n req: %v, error: %v", Car, err)
	}

	return nil
}
