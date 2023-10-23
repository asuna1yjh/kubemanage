package common

import (
	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"

// getCurrentUserID 获取当前登录的用户ID
func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
