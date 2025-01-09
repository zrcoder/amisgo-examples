package db

import (
	"database/sql"
	"errors"

	"github.com/zrcoder/amisgo-examples/todo-app/model"
)

const addUser = `
INSERT INTO users (name, password_hash)
VALUES (?, ?)
`

func AddUser(user *model.User) error {
	_, err := db.Exec(addUser, user.Name, user.PasswordHash)
	return err
}

const deleteUser = `DELETE FROM users WHERE id = ?`

func DeleteUser(id int64) error {
	_, err := db.Exec(deleteUser, id)
	return err
}

const updateUser = `
UPDATE users
SET name = ?, password_hash = ?
WHERE id = ?
`

func UpdateUser(user *model.User) error {
	_, err := db.Exec(updateUser, user.Name, user.PasswordHash, user.ID)
	return err
}

const getUser = `
SELECT id, name, password_hash
FROM users
WHERE id = ? LIMIT 1
`

func GetUser(id int64) (*model.User, error) {
	row := db.QueryRow(getUser, id)
	user := &model.User{}
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.PasswordHash,
	)
	if err == sql.ErrNoRows {
		err = errors.New("no user found")
	}
	return user, err
}

const getUserByName = `
SELECT id, name, password_hash
FROM users
WHERE name = ? LIMIT 1
`

func GetUserByName(name string) (*model.User, error) {
	row := db.QueryRow(getUserByName, name)
	user := &model.User{}
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.PasswordHash,
	)
	return user, err
}
