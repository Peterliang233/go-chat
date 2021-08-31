package user

import (
	"github.com/Peterliang233/go-chat/database"
	"github.com/Peterliang233/go-chat/errmsg"
	"github.com/Peterliang233/go-chat/model"
)

// CheckLogin 检查用户登录信息是否正确
func CheckLogin(user *model.User) (code int, err error) {
	user.Password = ScryptPassword(user.Password)

	var u model.User

	if err := database.Db.
		Where(map[string]interface{}{
			"username": user.Username,
			"password": user.Password,
		}).First(&u).
		Error; err != nil {
		return errmsg.ErrMysqlServer, err
	}

	return errmsg.Success, nil
}
