package model

type Car struct {
	ID            int    `json:"id"`
	BrandID       int    `json:"brandID" binding:"required"`
	BrandLabel    string `json:"brandLabel" binding:"required"`
	ModelID       int    `json:"modelID" binding:"required"`
	ModelLabel    string `json:"modelLabel" binding:"required"`
	GenID         int    `json:"genID" binding:"required"`
	GenLabel      string `json:"genLabel" binding:"required"`
	BodyTypeID    int    `json:"bodyTypeID" binding:"required"`
	BodyTypeLabel string `json:"bodyTypeLabel" binding:"required"`
	ModID         int    `json:"modID" binding:"required"`
	ModLabel      string `json:"modLabel" binding:"required"`
	YearID        int    `json:"yearID" binding:"required"`
	YearValue     int    `json:"year" binding:"required"`
	VinCode       string `json:"vinCode"`
	VinImage      string `json:"vinImage"`
	Comment       string `json:"comment"`
}

type IZapsRepo interface {
	Create(Car *Car) error
	Update(Car *Car) error
}
