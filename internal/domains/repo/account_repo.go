package repo

import "online-shopping/internal/domains/entity"

// AccountRepo 定义了账户存储相关的操作接口
type AccountRepo interface {
	CreateAccount(accountReqEntity *entity.AccountReq) (*entity.Account, error)
	GetAccountByID(id string) (*entity.Account, error)
	GetAccountByUsername(username string) (*entity.Account, error)
}
