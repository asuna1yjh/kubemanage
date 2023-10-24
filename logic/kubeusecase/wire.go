package kubeusecase

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewDeploymentUseCase)
