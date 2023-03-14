package main

import (
	"github.com/PGabrielDev/desafio1_goexpert/infra"
	"github.com/PGabrielDev/desafio1_goexpert/internal/client/use_cases"
)

func main() {
	saveFile := infra.NewSaveFile()
	useCase := use_cases.NewGetCotacao(saveFile)

	useCase.Execute()

}

