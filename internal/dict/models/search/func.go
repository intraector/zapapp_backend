package car_models

import (
	"database/sql"
	"fmt"

	tools "zap/internal/_shared"
)

type carModelReq struct {
	brandID int
	query   string
	limit   int
}

type carModel struct {
	ID    int
	Label string `json:"label"`
}

func search(dictDB *sql.DB, req carModelReq) ([]carModel, error) {

	query := fmt.Sprintf(`
		SELECT id_car_model, name FROM dictdb.car_model 
		WHERE id_car_mark =%v AND name LIKE '%v%%' LIMIT %v`,
		req.brandID, req.query, req.limit)

	rows, err := dictDB.Query(query)

	if err != nil {
		return nil, fmt.Errorf("/models\n req: %v, error: %v", req, err)
	}
	defer rows.Close()

	var output []carModel

	for rows.Next() {
		var item carModel

		if err := rows.Scan(&item.ID, &item.Label); err != nil {
			errS := fmt.Sprintf("/models\n req: %v, error: %v", req, err)
			tools.Loge(errS)
			continue
		}

		output = append(output, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("/models\n req: %v, error: %v", req, err)
	}

	return output, nil
}
