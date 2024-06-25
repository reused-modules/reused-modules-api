package presenter

import (
	"reused-modules-api/domain"
	"reused-modules-api/usecase"
)

type findAccountBalancePresenter struct{}

func NewFindAccountBalancePresenter() usecase.FindAccountBalancePresenter {
	return findAccountBalancePresenter{}
}

func (a findAccountBalancePresenter) Output(balance domain.Money) usecase.FindAccountBalanceOutput {
	return usecase.FindAccountBalanceOutput{Balance: balance.Float64()}
}
