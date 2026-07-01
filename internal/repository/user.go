package repository

import (
	"github.com/hanzjefferson/go-cleanarch/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	SQL *sqlx.DB
}

func NewUserRepo(
	sql *sqlx.DB,
) *UserRepo {
	return &UserRepo{
		SQL: sql,
	}
}

func (repo *UserRepo) FindByUsername(
	entity *entity.User,
	username string,
) error {
	query := `
		SELECT * FROM users
		WHERE username = ?
	`

	err := repo.SQL.Get(entity, query, username)
	return err
}

func (repo *UserRepo) HasExist(
	username string,
) (bool, error) {
	var exists bool
	err := repo.SQL.Get(
		&exists,
		"SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)",
		username,
	)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (repo *UserRepo) Create(
	entity *entity.User,
) error {
	query := `
		INSERT INTO users (email, username, password)
		VALUES (?, ?, ?)
	`

	result, err := repo.SQL.Exec(query, entity.Email, entity.Username, entity.Password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	entity.ID = uint(id)
	return nil
}
