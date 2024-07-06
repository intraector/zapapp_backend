package car_mark

import (
	"database/sql"
	"fmt"

	"github.com/theritikchoure/logx"
)

type carMark struct {
	ID    int
	Label string `json:"label"`
}

func searchInDB(dictDB *sql.DB, label string, limit int) ([]carMark, error) {
	query := fmt.Sprintf(`
		SELECT id_car_mark, name FROM car_mark
		WHERE name LIKE '%v%%' LIMIT %v`, label, limit)

	rows, err := dictDB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("/marks\n label: %q, error: %v", label, err)
	}
	defer rows.Close()

	var output []carMark

	for rows.Next() {
		var item carMark

		if err := rows.Scan(&item.ID, &item.Label); err != nil {
			errS := fmt.Sprintf("/marks\n label: %q, error: %v", label, err)
			logx.Log(errS, logx.FGRED, logx.BGBLACK)
			continue
		}

		output = append(output, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("cars mark %q: %v", label, err)
	}

	return output, nil
}
