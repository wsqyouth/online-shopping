package service

import (
	"context"
	"errors"
	"online-shopping/internal/domains/dto"
	"online-shopping/internal/domains/entity"
	repository "online-shopping/internal/domains/repo"
)

type AccountServiceRepo interface {
	CreateAccount(ctx context.Context, createAccontReq *dto.CreateAccountRequest) (account *entity.Account, err error)
	GetAccount(ctx context.Context, getAccountReq *dto.GetAccountRequest) (account *entity.Account, err error)
	GetAccountByID(ctx context.Context, id string) (account *entity.Account, err error)
}

type AccountServiceImpl struct {
	accountRepo repository.AccountRepo
}

func NewAccountService(repo repository.AccountRepo) AccountServiceRepo {
	return &AccountServiceImpl{accountRepo: repo}
}

func (s *AccountServiceImpl) CreateAccount(ctx context.Context, createAccontReq *dto.CreateAccountRequest) (account *entity.Account, err error) {
	if createAccontReq.Username == "" || createAccontReq.Password == "" {
		return nil, errors.New("username or password cannot be empty")
	}

	existingAccount, _ := s.accountRepo.GetAccountByUsername(createAccontReq.Username)
	if existingAccount != nil {
		return nil, errors.New("username already exists")
	}

	// 将 CreateAccountRequest DTO 转换为 Account 实体
	accountReqEntity := entity.NewAccountEntityFromDto(createAccontReq)

	// 调用 repo 的 CreateAccount 方法
	accountEntity, err := s.accountRepo.CreateAccount(accountReqEntity)
	if err != nil {
		return nil, err
	}

	// 返回创建的账户实体
	return accountEntity, nil
}

func (s *AccountServiceImpl) GetAccount(ctx context.Context, getAccountReq *dto.GetAccountRequest) (*entity.Account, error) {
	if getAccountReq.Username == "" {
		return nil, errors.New("username cannot be empty")
	}
	account, err := s.accountRepo.GetAccountByUsername(getAccountReq.Username)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *AccountServiceImpl) GetAccountByID(ctx context.Context, id string) (*entity.Account, error) {
	if id == "" {
		return nil, errors.New("account ID cannot be empty")
	}
	return s.accountRepo.GetAccountByID(id)
}
