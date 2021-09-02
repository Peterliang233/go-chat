package user

import (
	"github.com/Peterliang233/go-chat/errmsg"
	"github.com/Peterliang233/go-chat/model"
	Service "github.com/Peterliang233/go-chat/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Registry 用户注册
func Registry(c *gin.Context) {
	var u model.User
	_ = c.ShouldBind(&u)

	//fmt.Printf("%v\n",u)
	code := Service.CheckUsername(u.Username)

	if code != errmsg.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  errmsg.CodeMsg[code],
			"data": map[string]interface{}{
				"username": u.Username,
			},
		})

		return
	}

	var err error

	code, err = Service.AddUser(&u)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": code,
			"msg":  errmsg.CodeMsg[code],
			"data": map[string]interface{}{
				"username": u.Username,
			},
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  "注册成功",
		"data": map[string]interface{}{
			"username": u.Username,
		},
	})
}
