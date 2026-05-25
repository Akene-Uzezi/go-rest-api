package database

import (
	"context"
	"database/sql"
	"time"
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"-"`
}

func (m *UserModel) Insert(user *User, ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	
	query := "INSERT INTO users (email, name, password) VALUES ($1, $2, $3) RETURNING id"

	return  m.DB.QueryRowContext(ctx, query, user.Email, user.Name, user.Password).Scan(&user.Id)
}

func (m *UserModel) Delete(id int, ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := "DELETE FROM users WHERE id = $1"
	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}