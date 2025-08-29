package account

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/ksuid"
)

// server - service - repository - db
type Service interface {
	PostAaccount(ctx context.Context, name string) (*Account, error)
	GetAccountByID(ctx context.Context, id string) (*Account, error)
	GetAccounts(ctx context.Context, skip, take uint64) ([]Account, error)
}

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type accountService struct {
	repository Repository
}

func newService(r Repository) Service {
	return &accountService{repository: r}
}

// implementation of service interface
func (s *accountService) PostAaccount(ctx context.Context, name string) (*Account, error) {
	a := Account{
		ID:   ksuid.New().String(),
		Name: name,
	}
	err := s.repository.PutAccount(ctx, a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}
func (s *accountService) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	return s.repository.GetAccountByID(ctx, id)
}
func (s *accountService) GetAccounts(ctx context.Context, skip, take uint64) ([]Account, error) {
	return s.repository.ListAccounts(ctx, int(skip), int(take))
}

// very simple id generator
func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
