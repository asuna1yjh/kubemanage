package user

import (
	"gin_demo/controllers/common"
	"gin_demo/controllers/types"
	"gin_demo/pkg/jwt"

	"github.com/gin-gonic/gin"
)

/*
controller -> logic -> dao -> logic -> controller
处理逻辑，controller负责参数校验，返回数据
*/

// @Summary 用户登录
// @Description 用户登录接口
// @Tags 用户
// @Accept json
// @Produce json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /login [post]
func (u *UserRouter) LoginHandler(c *gin.Context) {
	// 1. 参数校验
	p := new(types.UserLoginRequest)
	if err := common.Parameter(c, p); err != nil {
		return
	}
	// 2. 逻辑处理
	data, err := u.us.Login(p)
	if err != nil {
		if err == common.ErrorInvalidPassword {
			common.ResponseError(c, common.CodeInvalidPassword)
			return
		} else if err == common.ErrorUserNotExist {
			common.ResponseError(c, common.CodeUserNotExist)
			return
		}
		common.ResponseError(c, common.CodeServerBusy)
		return
	}
	// 3. 返回数据
	token, err := jwt.GenToken(data.UserID, p.Username)
	if err != nil {
		return
	}
	common.ResponseSuccess(c, token)
}

func (u *UserRouter) RegisterHandler(c *gin.Context) {
	// 1. 参数校验
	p := new(types.RegisterRequest)
	if err := common.Parameter(c, p); err != nil {
		return
	}
	// 2. 逻辑处理由logic层处理
	data, err := u.us.Register(p)
	if err != nil {
		common.ResponseError(c, common.CodeServerBusy)
		return
	}
	res := &types.UserRegisterResponse{
		Username: data.UserName,
		Phone:    data.Phone,
		NickName: data.NickName,
		Email:    data.Email,
		Avatar:   data.Avatar,
	}
	// 3. 返回数据
	common.ResponseSuccess(c, res)
}

func (u *UserRouter) InfoHandler(c *gin.Context) {
	// 1. 参数校验
	//p := new(types.UserLoginRequest)
	//if err := common.Parameter(c, p); err != nil {
	//	return
	//}
	id, err := common.GetCurrentUserID(c)
	if err != nil {
		return
	}
	// 2. 逻辑处理
	info, err := u.us.GetUserInfo(id)
	if err != nil {
		common.ResponseError(c, common.CodeServerBusy)
		return
	}
	common.ResponseSuccess(c, info)
}
