package entity

import "errors"

type USDBRL struct {
	Bid string `json:"bid"`
}

func (u USDBRL) Validate() error {
	if u.Bid == "" {
		return errors.New("Valor em dolar nao pode estar vazio")
	}
	return nil
}
