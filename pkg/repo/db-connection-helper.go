package repo

import "database/sql"

type DatabaseConnectionHelper struct {
	DB *sql.DB
}

func NewDatabaseConnectionHelper() (*DatabaseConnectionHelper, error) {
	db, err := sql.Open("mysql", "root:ejek@tcp(127.0.0.1:3306)/go_crud_api_sample_db")

	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DatabaseConnectionHelper{
		DB: db,
	}, nil
}
