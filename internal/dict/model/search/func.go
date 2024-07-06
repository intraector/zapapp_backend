package car_model

import (
	"database/sql"
	"fmt"

	"github.com/theritikchoure/logx"
)

type carModel struct {
	ID    int
	Label string `json:"label"`
}

func searchInDB(dictDB *sql.DB, markID int, label string, limit int) ([]carModel, error) {

	query := fmt.Sprintf(`
		SELECT id_car_model, name FROM dictdb.car_model 
		WHERE id_car_mark =%v AND name LIKE '%v%%' LIMIT %v`, markID, label, limit)

	rows, err := dictDB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("/models\n markID: %q, label: %q, error: %v", markID, label, err)
	}
	defer rows.Close()

	var output []carModel

	for rows.Next() {
		var item carModel

		if err := rows.Scan(&item.ID, &item.Label); err != nil {
			errS := fmt.Sprintf("/models\n markID: %q, label: %q, error: %v", markID, label, err)
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
