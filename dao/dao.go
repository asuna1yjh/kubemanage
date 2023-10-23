package dao

import (
	"context"

	"gorm.io/gorm"
)

/*
负责处理数据层的逻辑，包括数据库的连接、事务的处理、数据的增删改查等。
*/
type ShareDaoFactory struct {
	ctx  context.Context
	User UserRepoInterface
}

func NewShareDaoFactory(ctx context.Context, db *gorm.DB) *ShareDaoFactory {
	return &ShareDaoFactory{
		User: NewUserRepo(ctx, db),
	}
}
