package dto

type CreateAccountRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateAccountResponse struct {
	AccountID string `json:"account_id"`
}

type GetAccountRequest struct {
	AccountID string `json:"account_id"`
	Username  string `json:"username" binding:"required"`
}
type GetAccountResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
