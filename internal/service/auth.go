package service

import (
	"context"
	"encoding/json"
	pb "kratos-blog/api/blog/v1"
	"kratos-blog/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"go.opentelemetry.io/otel"
)

type JwtUser struct {
	UID int64 `json:"uid"`
}

type AuthService struct {
	pb.UnimplementedAuthServer
	user *biz.UserUseCase
	auth *biz.AuthUseCase
	log  *log.Helper
}

func NewAuthService(user *biz.UserUseCase, auth *biz.AuthUseCase, logger log.Logger) *AuthService {
	return &AuthService{
		user: user,
		auth: auth,
		log:  log.NewHelper(logger),
	}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	s.log.Infof("Register input data %v", req)
	return s.auth.Register(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})

}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	s.log.Infof("Login input data %v", req)
	return s.auth.Login(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
}

// func (s *BlogService) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutReply, error) {
// 	return nil, nil
// }

func (s *AuthService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	user := JwtUser{}
	if claims, ok := jwt.FromContext(ctx); ok {
		arr, _ := json.Marshal(claims)
		json.Unmarshal(arr, &user)
	}
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetUser")
	defer span.End()
	p, err := s.user.GetUser(ctx, user.UID)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{Id: p.Id, Username: p.Username}, nil
}

// func (s *BlogService) VerifyPassword(ctx context.Context, req *pb.VerifyPasswordReq) (*pb.VerifyPasswordReply, error) {
// 	return nil, nil
// }
