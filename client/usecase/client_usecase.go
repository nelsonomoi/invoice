package usecase

import (
	"context"
	"time"

	"github.com/nelsonomoi/invoice/domain"
)


type clientUseCase struct {
	clientRepo domain.ClientRepository
	contextTimeout time.Duration
}


// NewArticleUsecase will create new an clientusecase object representation of domain.ClientUseCase interface
func NewClientUsecase(a domain.ClientRepository, timeout time.Duration) domain.ClientUseCase {
	return &clientUseCase{
		clientRepo:    a,
		contextTimeout: timeout,
	}
}


func (a *clientUseCase) Store(c context.Context, m *domain.Client) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	err = a.clientRepo.Store(ctx, m)
	return
}