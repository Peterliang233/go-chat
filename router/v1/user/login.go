package user

import (
	"github.com/Peterliang233/go-chat/model"
	Service "github.com/Peterliang233/go-chat/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 用户登录
func Login(c *gin.Context) {
	var login model.User

	_ = c.ShouldBind(&login)

	code, err := Service.CheckLogin(&login)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  "登录失败",
			"data": map[string]interface{}{
				"username": login.Username,
			},
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  "登录成功",
		"data": map[string]interface{}{
			"username": login.Username,
		},
	})

}
