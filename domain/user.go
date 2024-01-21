package domain

import (
	"context"
	"time"
)

type User struct {
	ID             int64     `json:"id" db:"id"`
	Email          string    `json:"title" db:"email"`
	VerifiedStatus string    `json:"content" db:"verifiedStatus"`
	Birthday       time.Time `json:"birthday" db:"birthday"`
}

type UserUsecase interface {
	Fetch(ctx context.Context, cursor string) ([]User, string, error)
	GetByID(ctx context.Context, id int64) (User, error)
	Update(ctx context.Context, ar *User) error
	Store(context.Context, *User) error
	Delete(ctx context.Context, id int64) error
}
