package car_brands

import (
	"database/sql"
	"fmt"

	tools "zap/internal/_shared"
)

type carBrand struct {
	ID    int
	Label string `json:"label"`
}

func searchInDB(dictDB *sql.DB, label string, limit int) ([]carBrand, error) {
	query := fmt.Sprintf(`
		SELECT id_car_mark, name FROM car_mark
		WHERE name LIKE '%v%%' LIMIT %v`, label, limit)

	rows, err := dictDB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("/marks\n label: %q, error: %v", label, err)
	}
	defer rows.Close()

	var output []carBrand

	for rows.Next() {
		var item carBrand

		if err := rows.Scan(&item.ID, &item.Label); err != nil {
			errS := fmt.Sprintf("/marks\n label: %q, error: %v", label, err)
			tools.Loge(errS)
			continue
		}

		output = append(output, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("cars mark %q: %v", label, err)
	}

	return output, nil
}
