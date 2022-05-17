package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/merdan2000/internal/model"
)

type RepoUser struct {
	db *sqlx.DB
}

func NewRepoUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db: db}
}

func (r RepoUser) CreateUser(user model.Users, password string) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO users (first_name, last_name, email, age, password) VALUES ($1,$2,$3,$4,$5)returning id`

	row := tx.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Age, password)
	err = row.Err()
	if err != nil {
		return 0, err
	}

	var ID int
	err = row.Scan(&ID)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return ID, nil

}

func (r RepoUser) ReadByID(userID int) (model.Users, error) {
	user := model.Users{}
	var password string
	query := `SELECT * FROM users WHERE id= $1`
	err := r.db.QueryRow(query, userID).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Age,
		password,
		&user.Created,
	)
	if err != nil {
		return model.Users{}, err
	}
	return user, nil
}

func (r RepoUser) ReadUser(email, password string) (model.Users, error) {
	user := model.Users{}
	query := `SELECT * FROM users WHERE email = $1 AND password = $2`
	err := r.db.QueryRow(query, email, password).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Age,
		password,
		&user.Created,
	)
	if err != nil {
		return model.Users{}, err
	}
	return user, nil
}

func (r RepoUser) UpdateUser(user model.Users) (model.Users, error) {
	tx, err := r.db.Begin()
	var users model.Users
	query := `UPDATE users SET first_name = $1, last_name = $2, email = $3, age = $4  WHERE id=$5 RETURNING *`

	row := tx.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Age, user.Id)

	err = row.Err()
	if err != nil {
		return model.Users{}, err
	}

	return users, nil
}

func (r RepoUser) DeleteUserByID(userID int, password string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	deleteUser := `UPDATE users SET is_deleted=true WHERE id=$1 and $2`

	err = tx.QueryRow(deleteUser, userID, password).Err()
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
