package repositories

import (
	"curso/devbook/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

func (repo users) Create(user models.User) (uint64, error) {
	statement, err := repo.db.Prepare("INSERT INTO users (name, nick, email, password) values (?,?,?,?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	res, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), err
}

func (repo users) FindAllByNameOrNick(nameOrNick string) ([]models.User, error) {
	nameOrNickFilter := fmt.Sprintf("%%%s%%", nameOrNick)

	rows, err := repo.db.Query("SELECT id, name, nick, email, createdAt FROM users WHERE name LIKE ? OR nick LIKE ?", nameOrNickFilter, nameOrNickFilter)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
