package usecases

import (
	"github.com/PGabrielDev/desafio1_goexpert/internal/server/entity"
	"github.com/PGabrielDev/desafio1_goexpert/internal/server/ports"
)

type GetValueDolarUseCase struct {
	TalkerDolar     ports.ITalkerDolarPort
	DolarRepository ports.IDolarRepository
}

func NewGetValueDolarUseCase(talkerDolar ports.ITalkerDolarPort, dolarRepository ports.IDolarRepository) *GetValueDolarUseCase {
	return &GetValueDolarUseCase{talkerDolar, dolarRepository}
}

func (u GetValueDolarUseCase) Execute() (*entity.USDBRL, error) {
	dolarValue, err := u.TalkerDolar.GetValueDolar()
	if err != nil {
		return nil, err
	}
	err = u.DolarRepository.Save(dolarValue)
	if err != nil {
		return nil, err
	}
	return dolarValue, nil
}
