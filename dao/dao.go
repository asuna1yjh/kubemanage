package dao

import (
	"github.com/google/wire"
)

/*
负责处理数据层的逻辑，包括数据库的连接、事务的处理、数据的增删改查等。
*/
type ShareDaoFactory struct {
	User *UserRepo
}

func NewShareDaoFactory(repo *UserRepo) *ShareDaoFactory {
	return &ShareDaoFactory{
		User: repo,
	}
}

var ProviderSet = wire.NewSet(NewUserRepo, NewShareDaoFactory)
