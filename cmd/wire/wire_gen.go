// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitRouter(ctx context.Context, db *gorm.DB) routes.Option {
	userRepo := dao.NewUserRepo(ctx, db)
	shareDaoFactory := dao.NewShareDaoFactory(userRepo)
	userUseCase := logic.NewUserUseCase(shareDaoFactory)
	option := user.NewUserRouter(userUseCase)
	return option
}

func InitKubeRouter() routes.Option {
	kubeFactory := kuberepo.NewKubeFactory()
	deployment := kuberepo.NewDeployment(kubeFactory)
	deploymentUseCase := kubeusecase.NewDeploymentUseCase(deployment)
	kubecontrollerDeployment := kubecontroller.NewDeployment(deploymentUseCase)
	option := kubecontroller.InitKubeRouter(kubecontrollerDeployment)
	return option
}