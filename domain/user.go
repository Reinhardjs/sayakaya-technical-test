package domain

import (
	"context"
	"encoding/json"
	"time"
)

type User struct {
	ID             int64     `json:"id" db:"id"`
	Email          string    `json:"email" db:"email"`
	VerifiedStatus bool      `json:"verifiedStatus" db:"verifiedStatus"`
	Birthday       time.Time `json:"birthday" db:"birthday"`
}

func (u *User) UnmarshalJSON(data []byte) error {
	type Alias User
	aux := &struct {
		Birthday string `json:"birthday"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	parsedBirthday, err := time.Parse("2006-01-02", aux.Birthday)
	if err != nil {
		return err
	}

	u.Birthday = parsedBirthday
	return nil
}

func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		*Alias
		Birthday string `json:"birthday"`
	}{
		Birthday: u.Birthday.Format("2006-01-02"),
		Alias:    (*Alias)(u),
	})
}

type UserUsecase interface {
	Fetch(ctx context.Context) ([]User, error)
	FetchByBirthDay(ctx context.Context, birthday time.Time) ([]User, error)
	GetByID(ctx context.Context, id int64) (User, error)
	Update(ctx context.Context, ar *User) error
	Store(context.Context, *User) error
	Delete(ctx context.Context, id int64) error
}

type UserRepository interface {
	Fetch(ctx context.Context) (res []User, err error)
	FetchByBirthDay(ctx context.Context, birthday time.Time) ([]User, error)
	GetByID(ctx context.Context, id int64) (User, error)
	Update(ctx context.Context, ar *User) error
	Store(ctx context.Context, a *User) error
	Delete(ctx context.Context, id int64) error
}
