package dict_repo

import (
	"fmt"

	tools "zap/internal/_shared"
	dict_model "zap/internal/dict/domain"
)

func (r *Repo) Generations(req dict_model.Req) ([]dict_model.DTO, error) {
	query := fmt.Sprintf(`
		SELECT id_car_generation, name, year_begin, year_end FROM dictdb.car_generation 
		WHERE id_car_model =%d AND name LIKE '%s%%' LIMIT %d`,
		req.BrandID, req.Query, req.Limit,
	)

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("/generations\n req: %v, error: %v", req, err)
	}
	defer rows.Close()

	var output []dict_model.DTO

	for rows.Next() {
		var item dict_model.DTO

		if err := rows.Scan(&item.ID, &item.Label, &item.YearBegin, &item.YearEnd); err != nil {
			errS := fmt.Sprintf("/generations\n req: %v, error: %v", req, err)
			tools.Loge(errS)
			continue
		}

		output = append(output, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("/generations\n req: %v, error: %v", req, err)
	}

	return output, nil
}
