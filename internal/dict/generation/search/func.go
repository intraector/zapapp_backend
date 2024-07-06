package car_generation

import (
	"database/sql"
	"fmt"

	"github.com/theritikchoure/logx"
)

type carGeneration struct {
	ID        int
	Label     *string `json:"label"`
	YearBegin *int    `json:"year_begin"`
	YearEnd   *int    `json:"year_end"`
}

func searchInDB(dictDB *sql.DB, modelID int, label string, limit int) ([]carGeneration, error) {
	query := fmt.Sprintf(`
	SELECT id_car_generation, name, year_begin, year_end FROM dictdb.car_generation 
	WHERE id_car_model =%v AND name LIKE '%v%%' LIMIT %v`, modelID, label, limit)

	rows, err := dictDB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("/generations\n modelID: %q, label: %q, error: %v", modelID, label, err)
	}
	defer rows.Close()

	var output []carGeneration
	for rows.Next() {
		var item carGeneration

		if err := rows.Scan(&item.ID, &item.Label, &item.YearBegin, &item.YearEnd); err != nil {
			errS := fmt.Sprintf("/generations\n modelID: %q, label: %q, error: %v", modelID, label, err)
			logx.Log(errS, logx.FGRED, logx.BGBLACK)
			continue
		}

		output = append(output, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("cars generations %q: %v", label, err)
	}

	return output, nil
}
