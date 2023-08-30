package repo

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/0x41gawor/go-crud-api-sample/pkg/model"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (this *UserRepository) Create(m *model.User) (int64, error) {

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)

	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf(
		"INSERT INTO users(login, password) VALUES ('%s', '%s');",
		m.Login,
		string(encryptedPassword),
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

func (this *UserRepository) FindByLogin(login string) (*model.User, error) {
	query := fmt.Sprintf(
		"SELECT * FROM users WHERE login = '%s'",
		login,
	)

	res, err := this.db.Query(query)
	defer res.Close()

	if err != nil {
		return nil, err
	}

	model := new(model.User)

	if res.Next() {
		err := res.Scan(
			&model.Id,
			&model.Login,
			&model.Password,
		)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("No item with given id")
	}

	return model, nil
}
