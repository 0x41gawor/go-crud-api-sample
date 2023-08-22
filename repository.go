package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DTO interface {
}

type Repository interface {
	Create(model *DTO) (int64, error)
	Read(int) (*DTO, error)
	Update(model *DTO) error
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

func (this *ContinentRepository) List() ([]*Continent, error) {
	rows, err := this.db.Query("select * from continents")
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

func (this *ContinentRepository) Create(model *Continent) (int64, error) {
	query := fmt.Sprintf(
		"INSERT INTO continents(name, population, gdp, gdp_per_capita) VALUES ('%s', %f, %f, %f);",
		model.name,
		model.population,
		model.gdp,
		model.gdpPerCapita,
	)

	res, err := this.db.Exec(query)

	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}
