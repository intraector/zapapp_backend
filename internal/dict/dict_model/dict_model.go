package dict_model

type Req struct {
	Query      string `json:"query"`
	Limit      int    `json:"limit"`
	BrandID    int    `json:"brandID"`
	ModelID    int    `json:"modelID"`
	BodyTypeID int    `json:"bodyTypeID"`
	GenID      int    `json:"genID"`
}

type DTO struct {
	ID    *int
	Label *string `json:"label"`
}
type DTOWithYears struct {
	DTO
	YearBegin *int `json:"year_begin"`
	YearEnd   *int `json:"year_end"`
}

type IDictRepo interface {
	Brands(req Req) ([]DTO, error)
	Models(req Req) ([]DTO, error)
	Generations(req Req) ([]DTOWithYears, error)
	BodyTypes(req Req) ([]DTO, error)
	Modifications(req Req) ([]DTO, error)
	Years(req Req) ([]DTO, error)
}
