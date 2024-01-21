package usecase

import (
	"context"
	"time"

	"github.com/reinhardjs/sayakaya/domain"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

// NewUserUsecase will create new an userUsecase object representation of domain.UserUsecase interface
func NewUserUsecase(a domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       a,
		contextTimeout: timeout,
	}
}

func (a *userUsecase) Fetch(c context.Context) (res []domain.User, nextCursor string, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.userRepo.Fetch(ctx)
	if err != nil {
		return nil, "", err
	}

	return
}

func (a *userUsecase) GetByID(c context.Context, id int64) (res domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.userRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (a *userUsecase) Update(c context.Context, ar *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	return a.userRepo.Update(ctx, ar)
}

func (a *userUsecase) Store(c context.Context, m *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err = a.userRepo.Store(ctx, m)

	return
}

func (a *userUsecase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	existedUser, err := a.userRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	if existedUser == (domain.User{}) {
		return domain.ErrNotFound
	}

	return a.userRepo.Delete(ctx, id)
}
