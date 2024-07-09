package dict_repo

import (
	"fmt"

	tools "zap/internal/_shared"
	dict_model "zap/internal/dict/domain"
)

func (r *Repo) Models(req dict_model.Req) ([]dict_model.DTO, error) {
	query := fmt.Sprintf(`
		SELECT id_car_model, name FROM dictdb.car_model 
		WHERE id_car_mark =%d AND name LIKE '%s%%' 
		LIMIT %d`,
		req.BrandID, req.Query, req.Limit,
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
			tools.LogError("/models", "Error scanning row", req, err)
			continue
		}

		output = append(output, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return output, nil
}
