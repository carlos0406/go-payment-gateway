package dto

import (
	"time"

	"github.com/carlos0406/go-gateway/internal/domain"
)

type CreateAccount struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type AccountOutput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Balance   float64   `json:"balance"`
	APIKey    string    `json:"api_key"`
}

func ToDomainAccount(input CreateAccount) *domain.Account {
	return domain.NewAccount(
		input.Name,
		input.Email,
	)
}

func FromDomainAccount(account *domain.Account) AccountOutput {
	return AccountOutput{
		ID:        account.ID,
		Name:      account.Name,
		Email:     account.Email,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
		Balance:   account.Balance,
		APIKey:    account.APIKey,
	}
}
