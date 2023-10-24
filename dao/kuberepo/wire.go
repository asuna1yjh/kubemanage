package kuberepo

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewKubeFactory, NewDeployment)
