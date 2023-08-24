package repo

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/0x41gawor/go-crud-api-sample/pkg/model"

	_ "github.com/go-sql-driver/mysql"
)

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

func (this *ContinentRepository) List() ([]*model.Continent, error) {
	rows, err := this.db.Query("SELECT * FROM continents")
	if err != nil {
		return nil, err
	}

	continents := []*model.Continent{}

	for rows.Next() {
		temp := new(model.Continent)
		err := rows.Scan(
			&temp.Id,
			&temp.Name,
			&temp.Population,
			&temp.Gdp,
			&temp.GdpPerCapita,
		)
		if err != nil {
			return nil, err
		}
		continents = append(continents, temp)
	}

	return continents, nil
}

func (this *ContinentRepository) Create(m *model.Continent) (int64, error) {
	query := fmt.Sprintf(
		"INSERT INTO continents(name, population, gdp, gdp_per_capita) VALUES ('%s', %f, %f, %f);",
		m.Name,
		m.Population,
		m.Gdp,
		m.GdpPerCapita,
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

func (this *ContinentRepository) Read(id int64) (*model.Continent, error) {
	query := fmt.Sprintf(
		"SELECT * FROM continents WHERE id =%d",
		id,
	)
	res, err := this.db.Query(query)
	defer res.Close()

	if err != nil {
		return nil, err
	}

	model := new(model.Continent)

	if res.Next() {
		err = res.Scan(&model.Id, &model.Name, &model.Population, &model.Gdp, &model.GdpPerCapita)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("No item with given id")
	}

	return model, nil
}

func (this *ContinentRepository) Update(id int64, m *model.Continent) error {
	query := fmt.Sprintf(
		`UPDATE continents 
		SET 
			name = '%s',
			population = %f,
			gdp = %f,
			gdp_per_capita = %f
		WHERE
		id = %d`,
		m.Name,
		m.Population,
		m.Gdp,
		m.GdpPerCapita,
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
