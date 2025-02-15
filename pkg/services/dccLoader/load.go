package dccLoader

import (
	"io"

	"github.com/gravestench/dcc"
)

func (s *Service) Load(filepath string) (*dcc.DCC, error) {
	if s.cache != nil {
		cachedData, isCached := s.cache.Retrieve(filepath)
		if isCached {
			return cachedData.(*dcc.DCC), nil
		}
	}

	s.logger.Info().Msgf("loading %v", filepath)

	stream, err := s.mpq.Load(filepath)
	if err != nil {
		s.logger.Fatal().Msgf("loading file %q: %v", filepath, err)
	}

	data, err := io.ReadAll(stream)
	if err != nil {
		s.logger.Fatal().Msgf("reading data: %v", err)
	}

	dccImage, err := dcc.FromBytes(data)
	if err != nil {
		s.logger.Fatal().Msgf("parsing dcc: %v", err)
	}

	if s.cache != nil {
		if err = s.cache.Insert(filepath, dccImage, len(data)); err != nil {
			s.logger.Error().Msgf("caching file %q: %v", err)
		}
	}

	return dccImage, nil
}
