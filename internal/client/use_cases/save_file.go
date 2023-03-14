package usecases

import (
	"context"
	"encoding/json"
	"github.com/PGabrielDev/desafio1_goexpert/internal/client/entity"
	"github.com/PGabrielDev/desafio1_goexpert/internal/client/ports"
	"io"
	"net/http"
	"time"
)

type GetCotacao struct {
	saveFile ports.ISaveFile
}

func NewGetCotacao(saveFile ports.ISaveFile) *GetCotacao {
	return &GetCotacao{saveFile: saveFile}
}

func (g GetCotacao) Execute() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	var c entity.USDBRL

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	err = json.Unmarshal(body, &c)

	if err != nil {
		return err
	}
	err = g.saveFile.SaveFile(c.Bid)
	if err != nil {
		return err
	}

	return nil
}
