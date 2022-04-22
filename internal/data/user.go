package data

import (
	"context"
	"kratos-blog/internal/biz"
	"kratos-blog/internal/data/ent/user"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (rp *userRepo) Find(ctx context.Context, id int64) (*biz.User, error) {
	p, err := rp.data.db.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:       p.ID,
		Username: p.Username,
		Password: p.Password,
	}, nil
}

func (rp *userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	p, err := rp.data.db.User.Query().Where(user.Username(username)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:       p.ID,
		Username: p.Username,
		Password: p.Password,
	}, nil
}

func (rp *userRepo) Save(ctx context.Context, user *biz.User) error {
	_, err := rp.data.db.User.
		Create().
		SetID(user.Id).
		SetUsername(user.Username).
		SetPassword(user.Password).
		Save(ctx)
	return err
}

func (rp *userRepo) VerifyPassword(ctx context.Context, user *biz.User, password string) error {
	return nil
}
