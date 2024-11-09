package entity

import "online-shopping/internal/domains/dto"

type AccountReq struct {
	Username string
	Password string
}

func NewAccountEntityFromDto(account *dto.CreateAccountRequest) *AccountReq {
	return &AccountReq{
		Username: account.Username,
		Password: account.Password,
	}
}
func (a *AccountReq) ToAccontDTO() *dto.CreateAccountRequest {
	return &dto.CreateAccountRequest{
		Username: a.Username,
		Password: a.Password,
	}
}

// 对应 db 的实体
type Account struct {
	ID       string
	Username string
	Password string
}
