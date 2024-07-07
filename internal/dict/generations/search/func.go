package car_generations

import (
	"database/sql"
	"fmt"

	tools "zap/internal/_shared"
)

type carGenerationReq struct {
	brandID int
	query   string
	limit   int
}

type carGeneration struct {
	ID        int
	Label     *string `json:"label"`
	YearBegin *int    `json:"year_begin"`
	YearEnd   *int    `json:"year_end"`
}

func search(dictDB *sql.DB, req carGenerationReq) ([]carGeneration, error) {
	query := fmt.Sprintf(`
		SELECT id_car_generation, name, year_begin, year_end FROM dictdb.car_generation 
		WHERE id_car_model =%v AND name LIKE '%v%%' LIMIT %v`,
		req.brandID, req.query, req.limit,
	)

	rows, err := dictDB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("/generations\n req: %v, error: %v", req, err)
	}
	defer rows.Close()

	var output []carGeneration
	for rows.Next() {
		var item carGeneration

		if err := rows.Scan(&item.ID, &item.Label, &item.YearBegin, &item.YearEnd); err != nil {
			errS := fmt.Sprintf("/generations\n req: %v, error: %v", req, err)
			tools.Loge(errS)
			continue
		}

		output = append(output, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("/generations\n req: %v, error: %v", req, err)
	}

	return output, nil
}
