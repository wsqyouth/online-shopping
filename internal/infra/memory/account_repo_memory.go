package memory

import (
	"errors"
	"online-shopping/internal/domains/entity"
	"online-shopping/internal/domains/repo"
	"sync"

	"github.com/google/uuid"
)

type accountRepoMemory struct {
	accounts map[string]*entity.Account
	mutex    sync.RWMutex
}

// NewAccountRepositoryMemory 创建一个新的内存账户仓储
func NewAccountRepositoryMemory() repo.AccountRepo {
	return &accountRepoMemory{
		accounts: make(map[string]*entity.Account),
	}
}

func (r *accountRepoMemory) CreateAccount(accountReq *entity.AccountReq) (*entity.Account, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.accounts[accountReq.Username]; exists {
		return nil, errors.New("username already exists")
	}

	accountEntity := &entity.Account{
		ID:       generateID(),
		Username: accountReq.Username,
		Password: accountReq.Password,
	}
	r.accounts[accountEntity.ID] = accountEntity
	return accountEntity, nil
}

func (r *accountRepoMemory) GetAccountByID(id string) (*entity.Account, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	account, exists := r.accounts[id]
	if !exists {
		return nil, errors.New("account not found")
	}

	return account, nil
}

func (r *accountRepoMemory) GetAccountByUsername(username string) (*entity.Account, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, account := range r.accounts {
		if account.Username == username {
			return account, nil
		}
	}

	return nil, errors.New("username has not regsisterd")
}

// 假设有一个用于生成唯一ID的函数
func generateID() string {
	return uuid.New().String()
}
