package service

import (
	"context"
	pb "kratos-blog/api/blog/v1"
	"kratos-blog/internal/biz"

	"go.opentelemetry.io/otel"
)

func (s *BlogService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	s.log.Infof("Register input data %v", req)
	return s.auth.Register(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})

}

func (s *BlogService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	s.log.Infof("Login input data %v", req)
	return s.auth.Login(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
}

// func (s *BlogService) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutReply, error) {
// 	return nil, nil
// }

func (s *BlogService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetUser")
	defer span.End()
	p, err := s.user.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{Id: p.Id, Username: p.Username}, nil
}

// func (s *BlogService) VerifyPassword(ctx context.Context, req *pb.VerifyPasswordReq) (*pb.VerifyPasswordReply, error) {
// 	return nil, nil
// }
