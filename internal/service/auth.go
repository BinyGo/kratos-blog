package service

import (
	"context"

	pb "kratos-blog/api/blog/v1"
)

type AuthService struct {
	pb.UnimplementedAuthServer
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{}, nil
}
func (s *AuthService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}
func (s *AuthService) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutReply, error) {
	return &pb.LogoutReply{}, nil
}
