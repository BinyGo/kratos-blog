package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id       int64
	Username string
	Password string
}

type UserRepo interface {
	Find(ctx context.Context, id int64) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	Save(ctx context.Context, u *User) error
	VerifyPassword(ctx context.Context, u *User, password string) error
}

type UserUseCase struct {
	repo   UserRepo
	authUc *AuthUseCase
}

func NewUserUseCase(repo UserRepo, logger log.Logger, authUc *AuthUseCase) *UserUseCase {
	return &UserUseCase{
		repo:   repo,
		authUc: authUc,
	}
}

func (uc *UserUseCase) GetUser(ctx context.Context, id int64) (p *User, err error) {
	p, err = uc.repo.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (uc *UserUseCase) Logout(ctx context.Context, u *User) error {
	return nil
}
