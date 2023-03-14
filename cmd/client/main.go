package main

import (
	"github.com/PGabrielDev/desafio1_goexpert/infra"
	usecases "github.com/PGabrielDev/desafio1_goexpert/internal/client/use_cases"
)

func main() {
	saveFile := infra.NewSaveFile()
	useCase := usecases.NewGetCotacao(saveFile)
	useCase.Execute()
}
