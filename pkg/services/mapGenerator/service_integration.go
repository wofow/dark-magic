package mapGenerator

import (
	"github.com/gravestench/runtime"

	"github.com/gravestench/dark-magic/pkg/services/webRouter"
)

var (
	_ runtime.Service              = &Service{}
	_ runtime.HasLogger            = &Service{}
	_ runtime.HasDependencies      = &Service{}
	_ webRouter.IsRouteInitializer = &Service{}
	_ GeneratesDiablo2Maps         = &Service{}
)

type Dependency = GeneratesDiablo2Maps

type GeneratesDiablo2Maps interface {
	Seed() uint64
	SetSeed(uint64)
	GenerateMap(act, difficulty uint) (*WorldMap, error)
}
