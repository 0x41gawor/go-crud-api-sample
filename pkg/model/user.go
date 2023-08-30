package model

import "fmt"

type User struct {
	Id       int64  `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func NewUser(id int64, login string, password string) *User {
	return &User{
		Id:       id,
		Login:    login,
		Password: password,
	}
}

func (m *User) String() string {
	return fmt.Sprintf("User{%d, %s, %s}", m.Id, m.Login, m.Password)
}
