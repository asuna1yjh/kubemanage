package logic

import (
	"errors"
	"gin_demo/controllers/common"
	"gin_demo/controllers/types"
	"gin_demo/dao"
	"gin_demo/models"
	"gin_demo/pkg/snowflake"
	"gin_demo/pkg/utils"

	"go.uber.org/zap"
)

type UserUseCase struct {
	db *dao.ShareDaoFactory
}

type UserInterface interface {
	Login(*types.UserLoginRequest) (*models.User, error)
	Register(*types.RegisterRequest) (*models.User, error)
	GetUserInfo(int64) (data *models.User, err error)
}

func NewUserUseCase(db *dao.ShareDaoFactory) UserInterface {
	return &UserUseCase{
		db: db,
	}
}

func (u *UserUseCase) Login(p *types.UserLoginRequest) (data *models.User, err error) {
	// 1. 根据用户名查询用户
	data, err = u.db.User.FindBy(p.Username)
	if err != nil {
		zap.L().Error("查询用户失败", zap.Error(err))
		return nil, common.ErrorServerBusy
	}
	// 1.1 用户不存在
	if data.UserName == "" {
		zap.L().Error(common.ErrorUserNotExist.Error())
		return nil, common.ErrorUserNotExist
	}
	// 2. 判断密码是否正确
	if data.Password != utils.Md5Sum([]byte(p.Password)) {
		zap.L().Error(common.ErrorInvalidPassword.Error())
		return nil, common.ErrorInvalidPassword
	}
	return
}

func (u *UserUseCase) Register(p *types.RegisterRequest) (data *models.User, err error) {
	userdata, err := u.db.User.FindBy(p.Username)
	if err != nil {
		zap.L().Error("查询用户失败", zap.Error(err))
		return nil, err
	}
	if userdata.UserName != "" {
		zap.L().Info("用户名已存在")
		return nil, errors.New("用户名已存在")
	}
	// 1.1 对密码进行加密
	p.Password = utils.Md5Sum([]byte(p.Password))
	// 1.2 生成用户ID
	UserID := snowflake.GenID()
	// 2. 保存用户
	d := &models.User{
		UserName: p.Username,
		Password: p.Password,
		Phone:    p.Phone,
		UserID:   UserID,
	}
	data, err = u.db.User.Save(d)
	if err != nil {
		return
	}
	return data, nil
}

func (u *UserUseCase) GetUserInfo(UserID int64) (data *models.User, err error) {
	data, err = u.db.User.FindByUserID(UserID)
	if err != nil {
		zap.L().Error("查询用户失败", zap.Error(err))
		return nil, err
	}
	return
}
