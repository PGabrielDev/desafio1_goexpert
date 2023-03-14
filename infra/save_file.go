package infra

import (
	"fmt"
	"os"
)

type SaveFile struct {
}

func (s *SaveFile) SaveFile(dolar string) error {

	f, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	_, err = f.Write([]byte("Dólar:{" + dolar + "}"))

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func NewSaveFile() *SaveFile {
	return &SaveFile{}
}
