package zap_create

import (
	"database/sql"
	"fmt"

	tools "zap/internal/_shared"
	model "zap/internal/zaps/domain"
)

func create(dictDB *sql.DB, req model.Car) error {

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
		req.BrandID, req.BrandLabel,
		req.ModelID, req.ModelLabel,
		req.GenID, req.GenLabel,
		req.BodyTypeID, req.BodyTypeLabel,
		req.ModID, req.ModLabel,
		req.YearID, req.YearValue,
		req.VinCode, req.VinImage,
		req.Comment,
	)

	tools.Logg(query)

	_, err := dictDB.Exec(query)

	if err != nil {
		return fmt.Errorf("/create\n req: %v, error: %v", req, err)
	}

	return nil
}
