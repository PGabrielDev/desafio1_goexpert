package ports

import (
	"github.com/PGabrielDev/desafio1_goexpert/internal/entity"
)

type ITalkerDolarPort interface {
	GetValueDolar() (*entity.USDBRL, error)
}
