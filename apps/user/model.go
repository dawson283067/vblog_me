package user

type User struct {
	// 用户Id
	Id int

	// 创建时间，时间戳 10位，秒
  	CreatedAt int64

	// 更新时间，时间戳 10位，秒
  	UpdatedAt int64

	// 用户参数
	*CreateUserRequest
}