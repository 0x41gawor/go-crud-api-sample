package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DTO interface {
}

type Repository interface {
	// niech create zwraca id stworzonego elementu, a nie caly element, od tego masz read: atomowosc operacji ;)
	Create(model *DTO) (int64, error)
	Read(id int64) (*DTO, error)
	Update(id int64, model *DTO) error
	Delete(id int64) error
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
	rows, err := this.db.Query("SELECT * FROM continents")
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

func (this *ContinentRepository) Read(id int64) (*Continent, error) {
	query := fmt.Sprintf(
		"SELECT * FROM continents WHERE id =%d",
		id,
	)
	res, err := this.db.Query(query)
	defer res.Close()

	if err != nil {
		return nil, err
	}

	model := new(Continent)

	if res.Next() {
		err = res.Scan(&model.id, &model.name, &model.population, &model.gdp, &model.gdpPerCapita)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("No item with given id")
	}

	return model, nil
}

func (this *ContinentRepository) Update(id int64, model *Continent) error {
	query := fmt.Sprintf(
		`UPDATE continents 
		SET 
			name = '%s',
			population = %f,
			gdp = %f,
			gdp_per_capita = %f
		WHERE
		id = %d`,
		model.name,
		model.population,
		model.gdp,
		model.gdpPerCapita,
		id,
	)

	_, err := this.db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func (this *ContinentRepository) Delete(id int64) error {
	query := fmt.Sprintf("DELETE FROM continents WHERE id = %d", id)
	_, err := this.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
