package talkers

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/PGabrielDev/desafio1_goexpert/internal/entity"
)

type TalkerGetDolar struct {
}

func NewTalkerGetDolar() *TalkerGetDolar {
	return &TalkerGetDolar{}
}

func (t TalkerGetDolar) GetValueDolar() (*entity.USDBRL, error) {
	cliente := http.Client{Timeout: time.Millisecond * 200}
	req, err := http.NewRequest("GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}
	res, err := cliente.Do(req)
	if err != nil {
		return nil, err
	}
	var usdbrl entity.USDBRL
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	var retorno map[string]interface{}
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &retorno); err != nil {
		return nil, err
	}

	if valor, ok := retorno["USDBRL"].(map[string]interface{}); ok {
		usdbrl.Bid = valor["bid"].(string)
	}

	return &usdbrl, nil
}
