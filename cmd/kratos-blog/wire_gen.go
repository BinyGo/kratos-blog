// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
	"kratos-blog/internal/biz"
	"kratos-blog/internal/conf"
	"kratos-blog/internal/data"
	"kratos-blog/internal/server"
	"kratos-blog/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, auth *conf.Auth, registry *conf.Registry, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	articleRepo := data.NewArticleRepo(dataData, logger)
	articleUsecase := biz.NewArticleUsecase(articleRepo, logger)
	blogService := service.NewBlogService(articleUsecase, logger)
	userRepo := data.NewUserRepo(dataData, logger)
	authUseCase := biz.NewAuthUseCase(auth, userRepo)
	userUseCase := biz.NewUserUseCase(userRepo, logger, authUseCase)
	authService := service.NewAuthService(userUseCase, authUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, auth, confData, blogService, authService, logger, tracerProvider)
	grpcServer := server.NewGRPCServer(confServer, blogService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, httpServer, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
