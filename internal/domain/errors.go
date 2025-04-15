package domain

import "errors"

var (
	ErrAccountNotFound    = errors.New("account not found")
	ErrDuplicatedApiKey   = errors.New("duplicated api key")
	ErrUnauthorizedAccess = errors.New("unauthorized")
	ErrInvoiceNotFound    = errors.New("invoice not found")
	ErrInvalidAmount      = errors.New("invalid amount")
	ErrInvalidStatus      = errors.New("invalid status")
)
