package models

import (
	"context"
	"go-test/utils"

	"github.com/jackc/pgx/v5"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func FindUser(username string) (User, error) {
	conn, err := utils.ConnectDB()
	if err != nil {
		return User{}, err
	}
	defer conn.Close()

	query := `SELECT id, username, password FROM users WHERE username = $1`
	rows, err := conn.Query(context.Background(), query, username)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	user, err := pgx.CollectOneRow[User](rows, pgx.RowToStructByName)
	if err != nil {
		return User{}, err
	}

	return user, err
}
