package usecase

import (
	"context"
	"time"

	"reused-modules-api/domain"
)

type (
	// FindAccountBalanceUseCase input port
	FindAccountBalanceUseCase interface {
		Execute(context.Context, domain.AccountID) (FindAccountBalanceOutput, error)
	}

	// FindAccountBalanceInput input data
	FindAccountBalanceInput struct {
		ID int64 `json:"balance" validate:"gt=0,required"`
	}

	// FindAccountBalancePresenter output port
	FindAccountBalancePresenter interface {
		Output(domain.Money) FindAccountBalanceOutput
	}

	// FindAccountBalanceOutput output data
	FindAccountBalanceOutput struct {
		Balance float64 `json:"balance"`
	}

	findBalanceAccountInteractor struct {
		repo       domain.AccountRepository
		presenter  FindAccountBalancePresenter
		ctxTimeout time.Duration
	}
)

// NewFindBalanceAccountInteractor creates new findBalanceAccountInteractor with its dependencies
func NewFindBalanceAccountInteractor(
	repo domain.AccountRepository,
	presenter FindAccountBalancePresenter,
	t time.Duration,
) FindAccountBalanceUseCase {
	return findBalanceAccountInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

// Execute orchestrates the use case
func (a findBalanceAccountInteractor) Execute(ctx context.Context, ID domain.AccountID) (FindAccountBalanceOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	account, err := a.repo.FindBalance(ctx, ID)
	if err != nil {
		return a.presenter.Output(domain.Money(0)), err
	}

	return a.presenter.Output(account.Balance()), nil
}
