package repo

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/0x41gawor/go-crud-api-sample/pkg/model"
)

type CountryRepository struct {
	db *sql.DB
}

func NewCountryRepository(db *sql.DB) *CountryRepository {
	return &CountryRepository{
		db: db,
	}
}

func (this *CountryRepository) List() ([]*model.Country, error) {
	rows, err := this.db.Query("SELECT * FROM countries")
	if err != nil {
		return nil, err
	}

	models := []*model.Country{}

	top5CitiesStr := new(string)

	for rows.Next() {
		temp := new(model.Country)
		err := rows.Scan(
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
		models = append(models, temp)
	}

	return models, nil
}

func (this *CountryRepository) Create(m *model.Country) (int64, error) {
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
