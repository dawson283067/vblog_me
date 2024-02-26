package user

import "context"

type Service interface {
	// 用户创建
	CreateUser(context.Context, *CreateUserRequest) (*User, error)

	// 查询用户列表
	QueryUser(context.Context, *QueryUserReqeust) (*UserSet, error)

	// 查看用户详情
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)

	// 用户修改

	// 用户删除

}

// 用户创建的参数
type CreateUserRequest struct {
	Username string
	Password string
	Role string
	Label map[string]string
}

// 查询用户列表
type QueryUserReqeust struct {
	// 分页大小，一页多少个
	PageSize int
	// 当前页，查询那一页的数据
	PageNumber int
	// 根据用户name查找用户
	Username string
}

type UserSet struct {
	// 总共有多少个
	Total int64
	// 当前查询的数据清单
	Items []*User
}

type DescribeUserRequest struct {
	UserId int64
}

