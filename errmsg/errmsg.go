package errmsg

const (
	ErrUsernameExist = iota
	ErrMysqlServer
	ErrPassword
	Success = 200
	Error   = 300
)

var CodeMsg = map[int]string{
	ErrUsernameExist: "用户名存在",
	ErrMysqlServer:   "数据库处理错误",
	ErrPassword:      "密码错误",
	Success:          "成功",
	Error:            "失败",
}
