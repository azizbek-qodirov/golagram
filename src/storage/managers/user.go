package managers

import (
	"api-test/src/models"
	"database/sql"
)

type UserManager struct {
	Conn *sql.DB
}

func NewUserManager(conn *sql.DB) *UserManager {
	return &UserManager{Conn: conn}
}

func (um *UserManager) CreateUser(user *models.User) error {
	query := "INSERT INTO users (telegram_id, fullname, joined_date) VALUES ($1, $2, $3)"
	_, err := um.Conn.Exec(query, user.Telegram_id, user.Fullname, user.JoinedDate)
	if err != nil {
		return err
	}
	return nil
}

func (um *UserManager) CheckIfExists(telegram_id int) (bool, error) {
	query := "SELECT COUNT(*) FROM users WHERE telegram_id = $1"
	var count int
	err := um.Conn.QueryRow(query, telegram_id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (um *UserManager) GetUser(telegram_id int) (*models.User, error) {
	query := "SELECT telegram_id, fullname, joined_date FROM users WHERE telegram_id = $1"
	row := um.Conn.QueryRow(query, telegram_id)
	user := &models.User{}
	err := row.Scan(&user.Telegram_id, &user.Fullname, &user.JoinedDate)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (um *UserManager) DeleteUser(telegram_id int) error {
	query := "DELETE FROM users WHERE telegram_id = $1"
	_, err := um.Conn.Exec(query, telegram_id)
	if err != nil {
		return err
	}
	return nil
}
