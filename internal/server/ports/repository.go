package ports

import (
	"github.com/PGabrielDev/desafio1_goexpert/internal/server/entity"
)

type IDolarRepository interface {
	Save(us *entity.USDBRL) error
}
