package dc6Loader

import (
	"github.com/rs/zerolog"

	"github.com/gravestench/runtime"

	"github.com/gravestench/dark-magic/pkg/cache"
	"github.com/gravestench/dark-magic/pkg/services/configFile"
	"github.com/gravestench/dark-magic/pkg/services/mpqLoader"
)

type Service struct {
	logger *zerolog.Logger
	mpq    mpqLoader.Dependency
	config configFile.Dependency
	cache  *cache.Cache
}

func (s *Service) Init(rt runtime.R) {

}

func (s *Service) Name() string {
	return "DC6 Loader"
}

func (s *Service) BindLogger(logger *zerolog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *zerolog.Logger {
	return s.logger
}
