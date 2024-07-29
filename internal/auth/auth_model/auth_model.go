package auth_model

type AuthCode struct {
	Account   string `json:"account" binding:"required"`
	Code      string `json:"code" binding:"required"`
	CreatedAt int    `json:"created_at"`
}
