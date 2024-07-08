package dict_repo

import (
	"fmt"

	tools "zap/internal/_shared"
	dict_model "zap/internal/dict/domain"
)

func (r *Repo) Years(req dict_model.Req) ([]dict_model.DTO, error) {
	query := fmt.Sprintf(`
		SELECT id, year FROM dictdb.year 
		WHERE  id_car_make =%d AND id_car_model =%d AND id_car_generation = %d 
		LIMIT %d`,
		req.BrandID, req.ModelID, req.GenID, req.Limit,
	)

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("/years\n req: %v, error: %v", req, err)
	}
	defer rows.Close()

	var output []dict_model.DTO
	for rows.Next() {
		var item dict_model.DTO

		if err := rows.Scan(&item.ID, &item.Label); err != nil {
			errS := fmt.Sprintf("/years\n req: %v, error: %v", req, err)
			tools.Loge(errS)
			continue
		}

		output = append(output, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("/years\n req: %v, error: %v", req, err)
	}

	return output, nil
}
