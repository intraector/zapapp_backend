package dict_repo

import (
	"fmt"

	"zap/internal/dict/dict_model"
	tools "zap/internal/tools"
)

func (r *Repo) Brands(req dict_model.Req) ([]dict_model.DTO, error) {
	query := fmt.Sprintf(`
		SELECT id_car_mark, name FROM car_mark
		WHERE name LIKE '%s%%' LIMIT %d`, req.Query, req.Limit,
	)

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var output []dict_model.DTO

	for rows.Next() {
		var item dict_model.DTO

		if err := rows.Scan(&item.ID, &item.Label); err != nil {
			tools.LogError("/brands", "Error scanning row", req, err)
			continue
		}

		output = append(output, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return output, nil
}
