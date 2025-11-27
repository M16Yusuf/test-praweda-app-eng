package repository

import (
	"context"
	"test-praweda-app-eng/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) FindOrCreate(ctx context.Context, username string) (model.User, error) {

	var user model.User

	// Check existing
	sqlGet := `SELECT id, username, created_at FROM users WHERE username = $1`
	err := r.db.QueryRow(ctx, sqlGet, username).Scan(&user.ID, &user.Username, &user.CreatedAt)
	if err == nil { // user exists
		return user, nil
	}

	// Create new user
	sqlCreate := `INSERT INTO users ( username) VALUES ($1) RETURNING id, username, created_at`
	err = r.db.QueryRow(ctx, sqlCreate, username).Scan(&user.ID, &user.Username, &user.CreatedAt)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
