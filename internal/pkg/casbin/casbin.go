package casbin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"kratos-blog/internal/conf"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport"
	_ "github.com/go-sql-driver/mysql"
)

type JwtUser struct {
	UID int64 `json:"uid"`
}

type contextKey string

const (
	defaultRBACModel = `
	[request_definition]
	r = sub, obj, act
	
	[policy_definition]
	p = sub, obj, act
	
	[role_definition]
	g = _, _
	
	[policy_effect]
	e = some(where (p.eft == allow))
	
	[matchers]
	m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act || r.sub == "root"
`
)

// loadRbacModel 加载RBAC模型
func loadRbacModel() (model.Model, error) {
	return model.NewModelFromString(defaultRBACModel)
}

func Server(config *conf.Data) middleware.Middleware {
	var log *log.Helper

	a, err := xormadapter.NewAdapter(config.Casbin.Driver, config.Casbin.Source, true)
	if err != nil {
		log.Fatalf("error: xormadapter-NewAdapter: %s", err)
	}
	m, err := loadRbacModel()
	if err != nil {
		log.Fatalf("error: loadRbacModel: %s", err)
	}
	enforcer, err := casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}

	//enforcer.LoadPolicy()

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {

			if enforcer == nil {
				return nil, errors.New("enforcer is nil")
			}
			user := JwtUser{}
			if claims, ok := jwt.FromContext(ctx); ok {
				arr, _ := json.Marshal(claims)
				json.Unmarshal(arr, &user)
			}
			if user.UID == 0 {
				return nil, errors.New("jwt get user failed")
			}
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, errors.New("transport.FromClientContext err")
			}

			fmt.Println("----tr.Endpoint()--------", tr.Endpoint())   // http://172.31.242.72:8000
			fmt.Println("----tr.Operation()--------", tr.Operation()) ///api.blog.v1.Blog/ListArticle

			//验证权限 最后的的参数act,待改为动态规则
			allowed, err := enforcer.Enforce(fmt.Sprintf("%d", user.UID), tr.Operation(), "get")
			//添加一条规则权限
			//enforcer.AddPolicy(fmt.Sprintf("%d", user.UID), tr.Operation(), "get")

			if err != nil {
				return nil, err
			}
			if !allowed {
				return nil, errors.New("casbin allowed false")
			}

			return handler(ctx, req)
		}
	}
}
