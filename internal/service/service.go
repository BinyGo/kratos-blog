package service

import (
	pb "kratos-blog/api/blog/v1"
	"kratos-blog/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewBlogService, NewAuthService)

type BlogService struct {
	pb.UnimplementedBlogServer
	log *log.Helper

	article *biz.ArticleUsecase
	user    *biz.UserUseCase
	auth    *biz.AuthUseCase
}

func NewBlogService(article *biz.ArticleUsecase, user *biz.UserUseCase, auth *biz.AuthUseCase, logger log.Logger) *BlogService {
	return &BlogService{
		article: article,
		user:    user,
		auth:    auth,
		log:     log.NewHelper(logger),
	}
}
