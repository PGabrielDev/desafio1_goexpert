package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/PGabrielDev/desafio1_goexpert/internal/server/entity"
	_ "modernc.org/sqlite"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "infra/database/dolar.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

type DolarRepository struct {
	DB *sql.DB
}

func NewDolarRepository() (*DolarRepository, error) {
	db, err := NewDB()
	if err != nil {
		return nil, err
	}
	return &DolarRepository{DB: db}, nil
}

func (dr DolarRepository) Save(us *entity.USDBRL) error {
	stmt, err := dr.DB.Prepare("INSERT INTO dolar_brl_value (bid) VALUES (?)")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer stmt.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()
	_, err = stmt.ExecContext(ctx, us.Bid)
	if err != nil {
		return err
	}
	return nil

}
