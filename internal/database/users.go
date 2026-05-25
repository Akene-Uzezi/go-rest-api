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

func (m *UserModel) GetAll(c context.Context) ([]*User, error) {
	ctx, cancel := context.WithTimeout(c, 3*time.Second)
	defer cancel()
	query := "SELECT * FROM users"
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Email, &user.Name, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
	
}