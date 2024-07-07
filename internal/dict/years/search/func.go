package car_years

import (
	"database/sql"
	"fmt"

	tools "zap/internal/_shared"
)

type carYearsReq struct {
	brandID int
	modelID int
	genID   int
	limit   int
}

type carYear struct {
	ID    int
	Label *string `json:"label"`
}

func searchInDB(dictDB *sql.DB, req carYearsReq) ([]carYear, error) {
	query := fmt.Sprintf(`
		SELECT id, year FROM dictdb.year 
		WHERE  id_car_make =%v AND id_car_model =%v AND id_car_generation = %v 
		LIMIT %v`,
		req.brandID, req.modelID, req.genID, req.limit,
	)

	rows, err := dictDB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("/years\n req: %v, error: %v", req, err)
	}
	defer rows.Close()

	var output []carYear
	for rows.Next() {
		var item carYear

		if err := rows.Scan(&item.ID, &item.Label); err != nil {
			errS := fmt.Sprintf("/years\n req: %v, error: %v", req, err)
			tools.Loge(errS)
			continue
		}

		output = append(output, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("/years\n req: %v, error: %v", req, err)
	}

	return output, nil
}
