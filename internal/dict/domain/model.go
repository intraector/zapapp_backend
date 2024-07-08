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

type IDictRepo interface {
	Brands(req Req) ([]DTO, error)
	Years(req Req) ([]DTO, error)
}
