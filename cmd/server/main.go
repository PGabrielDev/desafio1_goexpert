package main

import (
	"encoding/json"
	"fmt"
	"github.com/PGabrielDev/desafio1_goexpert/infra/database"
	"github.com/PGabrielDev/desafio1_goexpert/infra/talkers"
	"github.com/PGabrielDev/desafio1_goexpert/internal/use_cases"
	"net/http"
)

func main() {

	talker := talkers.NewTalkerGetDolar()
	repository, err := database.NewDolarRepository()
	if err != nil {
		panic(err)
	}
	usecase := usecases.NewGetValueDolarUseCase(talker, repository)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		us, err := usecase.Execute()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err.Error())
			return
		}
		writer.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(writer).Encode(us); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err.Error())
			return
		}
	})
	http.ListenAndServe(":8000", nil)

}
