package models

import (
	"database/sql"
	"strings"
)

// UserDao represents the object transaction model
type UserDao struct {
	Connection *sql.DB
}

// Returns a new instance of userDao
func NewUserDao(connection *sql.DB) *UserDao {
	return &UserDao{Connection: connection}
}

func (userDao *UserDao) Register(user User) error {
	query := "INSERT INTO users(name, lastName, cpf, password, avatarURL, UUID, email) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := userDao.Connection.Exec(query, user.Name,
		user.LastName, user.Cpf, user.Password,
		user.AvatarURL, user.UUID, user.Email)
	return err
}

func (userDao *UserDao) FindOne(email string) (*User, error) {
	user := &User{}
	query := "SELECT id, name, email, password, avatarURL, uuid, datastart FROM users WHERE email=$1"
	row := userDao.Connection.QueryRow(query, email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password,
		&user.AvatarURL, &user.UUID, &user.DataStart)
	return user, err
}

func (userDao *UserDao) FindByUUID(uuid string) (*User, error) {
	user := &User{}
	query := "SELECT id, name, email, lastName, avatarURL, uuid, cpf, datastart FROM users WHERE uuid=$1"
	row := userDao.Connection.QueryRow(query, strings.TrimSpace(uuid))
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.LastName,
		&user.AvatarURL, &user.UUID, &user.Cpf, &user.DataStart)
	return user, err
}

func (userDao *UserDao) RemoveUser(id int) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := userDao.Connection.Exec(query, id)
	return err
}

func (userDao *UserDao) UpdateUser(user User) (*User, error) {
	query := `UPDATE users 
		SET name=$2, email=$3, lastName=$4,
		avatarURL=$5, cpf=$6 WHERE id=$1`
	_, err := userDao.Connection.Exec(query, user.ID,
		user.Name, user.Email, user.LastName,
		user.AvatarURL, user.Cpf)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
