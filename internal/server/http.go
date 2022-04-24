package server

import (
	"context"
	v1 "kratos-blog/api/blog/v1"
	"kratos-blog/internal/conf"
	"kratos-blog/internal/service"

	myCasbin "kratos-blog/internal/pkg/casbin"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	jwt4 "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"
)

func NewWhiteListMatcher() selector.MatchFunc {

	whiteList := make(map[string]struct{})
	whiteList["/api.blog.v1.Auth/Login"] = struct{}{}
	whiteList["/api.blog.v1.Auth/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, cd *conf.Data, blog *service.BlogService, auth *service.AuthService, logger log.Logger) *http.Server {
	// m, _ := model.NewModelFromFile("../../configs/authz/authz_model.conf")
	// a := fileAdapter.NewAdapter("../../configs/authz/authz_policy.csv")
	// a, err := entAdapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/kratos-blog?parseTime=True")
	// if err != nil {
	// 	fmt.Println("-----------------------------")
	// 	fmt.Println(err)
	// }

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			validate.Validator(),
			selector.Server(
				jwt.Server(
					func(token *jwt4.Token) (interface{}, error) {
						return []byte(ac.ApiKey), nil
					},
					jwt.WithSigningMethod(jwt4.SigningMethodHS256),
					jwt.WithClaims(func() jwt4.Claims {
						return &jwt4.MapClaims{}
					}),
				),
				myCasbin.Server(cd),
			).Match(NewWhiteListMatcher()).Build(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)
	v1.RegisterBlogHTTPServer(srv, blog)
	v1.RegisterAuthHTTPServer(srv, auth)
	return srv
}
