package car_body_types

import (
	"database/sql"
	"fmt"

	tools "zap/internal/_shared"
)

type carBodyTypeReq struct {
	brandID int
	genID   int
	query   string
	limit   int
}

type carBodyType struct {
	ID    int
	Label *string `json:"label"`
}

func searchInDB(dictDB *sql.DB, req carBodyTypeReq) ([]carBodyType, error) {
	query := fmt.Sprintf(`
		SELECT id_car_serie, name FROM dictdb.car_serie 
		WHERE id_car_model =%v AND id_car_generation = %v AND name LIKE '%v%%' LIMIT %v`,
		req.brandID, req.genID, req.query, req.limit,
	)

	rows, err := dictDB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("/bodyType\n req: %v, error: %v", req, err)
	}
	defer rows.Close()

	var output []carBodyType
	for rows.Next() {
		var item carBodyType

		if err := rows.Scan(&item.ID, &item.Label); err != nil {
			errS := fmt.Sprintf("/bodyType\n req: %v, error: %v", req, err)
			tools.Log(errS)
			continue
		}

		output = append(output, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("/bodyType\n req: %v, error: %v", req, err)
	}

	return output, nil
}
