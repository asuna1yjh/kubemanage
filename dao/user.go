package dao

import (
	"context"
	"gin_demo/models"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

type UserRepo struct {
	ctx context.Context
	db  *gorm.DB
}

// 数据库操作抽象方法集合
type UserRepoInterface interface {
	FindBy(string) (*models.User, error)
	Save(*models.User) (*models.User, error)
	FindByUserID(int64) (*models.User, error)
}

func NewUserRepo(ctx context.Context, db *gorm.DB) UserRepoInterface {
	return &UserRepo{
		ctx: ctx,
		db:  db,
	}
}

func (u *UserRepo) FindByUserID(userID int64) (*models.User, error) {
	var user = new(models.User)
	err := u.db.WithContext(u.ctx).Where("user_id = ?", userID).Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func (u *UserRepo) FindBy(username string) (*models.User, error) {
	var user = new(models.User)
	err := u.db.WithContext(u.ctx).Where("username = ?", username).Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func (u *UserRepo) Save(user *models.User) (*models.User, error) {
	err := u.db.WithContext(u.ctx).Create(user).Error
	if err != nil {
		zap.L().Error("创建用户失败", zap.Error(err))
		return nil, err
	}
	// 1. 查询为空，返回nil
	// 2. 查询出错，返回err
	// 3. 查询成功，返回user
	return user, err
}
