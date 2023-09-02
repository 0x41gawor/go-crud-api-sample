package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/0x41gawor/go-crud-api-sample/pkg/model"
)

type RepositoryCountry struct {
	db *sql.DB
}

func NewCountryRepository(db *sql.DB) *RepositoryCountry {
	return &RepositoryCountry{
		db: db,
	}
}

func (r *RepositoryCountry) List() ([]*model.Country, error) {
	res, err := r.db.Query("SELECT * FROM countries")
	if err != nil {
		return nil, err
	}

	models := []*model.Country{}

	for res.Next() {
		temp := new(model.Country)

		top5CitiesStr := new(string)
		err := res.Scan(
			&temp.Id,
			&temp.Name,
			&temp.ContinentId,
			&temp.Population,
			&temp.Gdp,
			&temp.GdpPerCapita,
			&top5CitiesStr,
		)
		if err != nil {
			return nil, err
		}
		temp.Top5Cities = strings.Split(*top5CitiesStr, ",")
		temp.Top5Cities = removeSpacesFromStringSlice(temp.Top5Cities)
		models = append(models, temp)
	}

	return models, nil
}

func (r *RepositoryCountry) Create(m *model.Country) (int64, error) {
	top5citiesStr := strings.Join(m.Top5Cities, ",")

	query := fmt.Sprintf(
		"INSERT INTO countries(name, continent_id, population, gdp, gdp_per_capita, top5cities) VALUES ('%s', %d, %f, %f, %f, '%s');",
		m.Name,
		m.ContinentId,
		m.Population,
		m.Gdp,
		m.GdpPerCapita,
		top5citiesStr,
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

func (r *RepositoryCountry) Read(id int64) (*model.Country, error) {
	query := fmt.Sprintf(
		"SELECT * FROM countries WHERE id =%d",
		id,
	)
	res, err := r.db.Query(query)
	defer res.Close()

	if err != nil {
		return nil, err
	}

	model := new(model.Country)

	if res.Next() {
		top5CitiesStr := new(string)
		err := res.Scan(
			&model.Id,
			&model.Name,
			&model.ContinentId,
			&model.Population,
			&model.Gdp,
			&model.GdpPerCapita,
			&top5CitiesStr,
		)
		if err != nil {
			return nil, err
		}
		model.Top5Cities = strings.Split(*top5CitiesStr, ",")
		model.Top5Cities = removeSpacesFromStringSlice(model.Top5Cities)
	} else {
		return nil, errors.New("No item with given id")
	}

	return model, nil
}

func (r *RepositoryCountry) Update(id int64, m *model.Country) error {
	query := fmt.Sprintf(
		`UPDATE countries
		SET
			name = '%s',
			continent_id = '%d',
			population = %f,
			gdp = %f,
			gdp_per_capita = %f,
			top5cities = '%s'
		WHERE	
		id = %d`,
		m.Name,
		m.ContinentId,
		m.Population,
		m.Gdp,
		m.GdpPerCapita,
		strings.Join(m.Top5Cities, ","),
		id,
	)

	_, err := r.db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryCountry) Delete(id int64) (bool, error) {
	query := fmt.Sprintf("DELETE FROM countries WHERE id = %d", id)
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
