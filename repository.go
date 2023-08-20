package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DTO interface {
}

type Repository interface {
	Create(*DTO) error
	Read(int) (*DTO, error)
	Update(*DTO) error
	Delete(int) error
	List() ([]*DTO, error)
}

type ContinentRepository struct {
	db *sql.DB
}

func NewContinentRepository() (*ContinentRepository, error) {
	db, err := sql.Open("mysql", "root:ejek@tcp(127.0.0.1:3306)/go_crud_api_sample_db")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &ContinentRepository{
		db: db,
	}, nil
}

func (r *ContinentRepository) List() ([]*Continent, error) {
	rows, err := r.db.Query("select * from continents")
	if err != nil {
		return nil, err
	}

	continents := []*Continent{}

	for rows.Next() {
		temp := new(Continent)
		err := rows.Scan(
			&temp.id,
			&temp.name,
			&temp.population,
			&temp.gdp,
			&temp.gdpPerCapita,
		)
		if err != nil {
			return nil, err
		}
		continents = append(continents, temp)
	}

	return continents, nil
}
