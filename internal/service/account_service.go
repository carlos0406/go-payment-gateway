package service

import (
	"github.com/carlos0406/go-gateway/internal/domain"
	"github.com/carlos0406/go-gateway/internal/dto"
)

type AccountService struct {
	accountRepository domain.AccountRepository
}

func NewAccountService(accountRepository domain.AccountRepository) *AccountService {
	return &AccountService{
		accountRepository: accountRepository,
	}
}

func (s *AccountService) CreateAccount(input dto.CreateAccount) (*dto.AccountOutput, error) {
	account := dto.ToDomainAccount(input)
	err := s.accountRepository.ExistsByApiKey(account.APIKey)
	if err == nil {
		return nil, err
	}

	err = s.accountRepository.Save(account)

	if err != nil {
		return nil, err
	}
	output := dto.FromDomainAccount(account)
	return &output, nil
}

func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutput, error) {
	account, err := s.accountRepository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	account.AddBalance(amount)
	err = s.accountRepository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromDomainAccount(account)
	return &output, nil
}

func (s *AccountService) FindByAPIKey(apiKey string) (*dto.AccountOutput, error) {
	account, err := s.accountRepository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	output := dto.FromDomainAccount(account)
	return &output, nil
}

func (s *AccountService) FindByID(id string) (*dto.AccountOutput, error) {
	account, err := s.accountRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	output := dto.FromDomainAccount(account)
	return &output, nil
}

func (s *AccountService) ExistsByApiKey(apiKey string) error {
	return s.accountRepository.ExistsByApiKey(apiKey)
}
