//go:build wireinject
// +build wireinject

package wire

import (
	"context"
	"gin_demo/controllers/kubecontroller"
	"gin_demo/controllers/user"
	"gin_demo/dao"
	"gin_demo/dao/kuberepo"
	"gin_demo/logic"
	"gin_demo/logic/kubeusecase"
	"gin_demo/routes"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitRouter(ctx context.Context, db *gorm.DB) routes.Option {
	panic(wire.Build(dao.ProviderSet, logic.ProviderSet, user.ProviderSet))
	return func(r *gin.RouterGroup) {}
}

func InitKubeRouter() routes.Option {
	panic(wire.Build(kubecontroller.ProviderSet, kuberepo.ProviderSet, kubeusecase.ProviderSet))
	return func(r *gin.RouterGroup) {}
}
