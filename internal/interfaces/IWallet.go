package interfaces

import (
	"context"
	"ewallet-wallet/internal/models"

	"github.com/gin-gonic/gin"
)

type IWalletRepo interface {
	CreateWallet(ctx context.Context, wallet *models.Wallet) error
	UpdateBalance(ctx context.Context, userID int, amount float64) (models.Wallet, error)
	CreateWalletTrx(ctx context.Context, walletTrx *models.WalletTransaction) error
	GetWalletTransactionByReference(ctx context.Context, reference string) (models.WalletTransaction, error)
}

type IWalletService interface {
	Create(ctx context.Context, wallet *models.Wallet) error
	CreditBalance(ctx context.Context, userID int, req models.TransactionRequest) (models.TransactionResponse, error)
	DebitBalance(ctx context.Context, userID int, req models.TransactionRequest) (models.TransactionResponse, error)
}

type IWalletAPI interface {
	Create(*gin.Context)
	CreditBalance(c *gin.Context)
	DebitBalance(c *gin.Context)
}
