package biz

import (
	"context"
	"errors"
	"kratos-blog/internal/conf"
	"math/rand"
	"time"

	pb "kratos-blog/api/blog/v1"

	"github.com/golang-jwt/jwt/v4"
)

type AuthUseCase struct {
	key      string
	userRepo UserRepo
}

func NewAuthUseCase(conf *conf.Auth, userRepo UserRepo) *AuthUseCase {
	return &AuthUseCase{
		key:      conf.ApiKey,
		userRepo: userRepo,
	}
}

func (uc *AuthUseCase) Login(ctx context.Context, user *User) (*pb.LoginReply, error) {
	user, err := uc.userRepo.FindByUsername(ctx, user.Username)
	if err != nil {
		return nil, err
	}
	err = uc.userRepo.VerifyPassword(ctx, user, user.Password)
	if err != nil {
		return nil, err
	}
	// generate token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": user.Id,
		"exp": time.Now().Unix() + 60*60*24,
	})
	signedString, err := claims.SignedString([]byte(uc.key))

	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{Token: signedString}, nil

}

func (uc *AuthUseCase) Register(ctx context.Context, user *User) (*pb.RegisterReply, error) {
	user, err := uc.userRepo.FindByUsername(ctx, user.Username)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("用户名已存在")
	}
	user.Id = rand.Int63()
	err = uc.userRepo.Save(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterReply{Id: user.Id}, nil
}
