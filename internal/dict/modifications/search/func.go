package car_modifications

import (
	"database/sql"
	"fmt"

	tools "zap/internal/_shared"
)

type carModificationsReq struct {
	brandID    int
	bodyTypeID int
	query      string
	limit      int
}

type carModification struct {
	ID    int
	Label *string `json:"label"`
}

func searchInDB(dictDB *sql.DB, req carModificationsReq) ([]carModification, error) {
	query := fmt.Sprintf(`
		SELECT id_car_modification, name FROM dictdb.car_modification 
		WHERE id_car_model =%v AND id_car_serie = %v AND name LIKE '%v%%' LIMIT %v`,
		req.brandID, req.bodyTypeID, req.query, req.limit,
	)

	rows, err := dictDB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("/modifications\n req: %v, error: %v", req, err)
	}
	defer rows.Close()

	var output []carModification
	for rows.Next() {
		var item carModification

		if err := rows.Scan(&item.ID, &item.Label); err != nil {
			errS := fmt.Sprintf("/modifications\n req: %v, error: %v", req, err)
			tools.Log(errS)
			continue
		}

		output = append(output, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("/modifications\n req: %v, error: %v", req, err)
	}

	return output, nil
}
