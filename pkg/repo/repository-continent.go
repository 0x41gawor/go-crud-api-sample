package repo

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/0x41gawor/go-crud-api-sample/pkg/model"

	_ "github.com/go-sql-driver/mysql"
)

type RepositoryContinent struct {
	db *sql.DB
}

func NewContinentRepository(db *sql.DB) *RepositoryContinent {
	return &RepositoryContinent{
		db: db,
	}
}

func (r *RepositoryContinent) List() ([]*model.Continent, error) {
	res, err := r.db.Query("SELECT * FROM continents")
	if err != nil {
		return nil, err
	}

	models := []*model.Continent{}

	for res.Next() {
		temp := new(model.Continent)
		err := res.Scan(
			&temp.Id,
			&temp.Name,
			&temp.Population,
			&temp.Gdp,
			&temp.GdpPerCapita,
		)
		if err != nil {
			return nil, err
		}
		models = append(models, temp)
	}

	return models, nil
}

func (r *RepositoryContinent) Create(m *model.Continent) (int64, error) {
	query := fmt.Sprintf(
		"INSERT INTO continents(name, population, gdp, gdp_per_capita) VALUES ('%s', %f, %f, %f);",
		m.Name,
		m.Population,
		m.Gdp,
		m.GdpPerCapita,
	)

	res, err := r.db.Exec(query)

	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (r *RepositoryContinent) Read(id int64) (*model.Continent, error) {
	query := fmt.Sprintf(
		"SELECT * FROM continents WHERE id =%d",
		id,
	)
	res, err := r.db.Query(query)
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

func (r *RepositoryContinent) Update(id int64, m *model.Continent) error {
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

	_, err := r.db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryContinent) Delete(id int64) (bool, error) {
	query := fmt.Sprintf("DELETE FROM continents WHERE id = %d", id)
	res, err := r.db.Exec(query)
	if err != nil {
		return false, err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	var result bool
	if affectedRows != 0 {
		result = true
	} else {
		result = false
	}

	return result, nil
}
