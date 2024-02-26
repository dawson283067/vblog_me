package impl

import (
	"context"

	"github.com/vblog_me/apps/user"
)

func (i *UserServiceImpl) CreateUser(ctx context.Context, in *user.CreateUserRequest) (*user.User, error) {
	return nil, nil
}

func (i *UserServiceImpl) QueryUser(ctx context.Context, in *user.QueryUserReqeust) (*user.UserSet, error) {
	return nil, nil
}

func (i *UserServiceImpl) DescribeUser(ctx context.Context, in *user.DescribeUserRequest) (*user.User, error) {
	return nil, nil
}