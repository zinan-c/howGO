package repository

import (
	"context"
	"howgo/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	FindByID(ctx context.Context, id int) (*model.User, error)
}

// type userRepo struct {
//     db *sqlx.DB
// }

// func NewUserRepository(db *sqlx.DB) UserRepository {
//     return &userRepo{db: db}
// }
