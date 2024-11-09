package account

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"online-shopping/internal/domains/dto"
	accountSer "online-shopping/internal/domains/services/account"
	"online-shopping/internal/infra/memory"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	accountServiceRepo accountSer.AccountServiceRepo
}

func NewAccountHandler() *AccountHandler {
	// 初始化仓储
	accountMemRepo := memory.NewAccountRepositoryMemory()
	// 初始化服务
	accountService := accountSer.NewAccountService(accountMemRepo)
	// 创建处理器
	return &AccountHandler{accountServiceRepo: accountService}
}

func (h *AccountHandler) RegisterRoutes(router *gin.Engine) {
	v1 := router.Group("/v1/accounts")
	{
		v1.POST("", h.CreateAccount)
		v1.GET("", h.GetAccountByID)         // 根据 ID 获取账户信息
		v1.GET("/by-username", h.GetAccount) // 根据用户名获取账户信息
	}
}

// 处理创建账号的请求逻辑
func (h *AccountHandler) CreateAccount(c *gin.Context) {
	ctx := c.Request.Context()

	var request dto.CreateAccountRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	account, err := h.accountServiceRepo.CreateAccount(ctx, &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.CreateAccountResponse{AccountID: account.ID})
}

// 处理获取账号的请求逻辑
func (h *AccountHandler) GetAccountByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Query("id")
	account, err := h.accountServiceRepo.GetAccountByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hasher := md5.New()
	hasher.Write([]byte(account.Password))
	passwordHash := hex.EncodeToString(hasher.Sum(nil))

	c.JSON(http.StatusOK, dto.GetAccountResponse{
		Username: account.Username,
		Password: passwordHash,
	})
}

func (h *AccountHandler) GetAccount(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	account, err := h.accountServiceRepo.GetAccount(c.Request.Context(), &dto.GetAccountRequest{Username: username})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": account.Username,
		"id":       account.ID,
		// Assume password is hashed; never send plain passwords
		// "password": account.Password,
	})
}
