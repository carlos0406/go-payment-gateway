package domain

type AccountRepository interface {
	Save(account *Account) error
	FindByID(id string) (*Account, error)
	UpdateBalance(account *Account) error
	FindByAPIKey(apiKey string) (*Account, error)
	ExistsByApiKey(apiKey string) error
}

type InvoiceRepository interface {
	Save(invoice *Invoice) error
	FindByID(id string) (*Invoice, error)
	FindByAccountID(accountID string, pagination *Pagination) (*PaginationOutput[Invoice], error)
	UpdateStatus(invoice *Invoice) error
}
