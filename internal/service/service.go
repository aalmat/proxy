package service

import (
	"github.com/aalmat/proxy/internal/domain/request"
	"github.com/aalmat/proxy/internal/domain/response"
)

type Configuration func(s *Service) error

type Service struct {
	requestRepository  request.Repository
	responseRepository response.Repository
}

func New(configs ...Configuration) (s *Service, err error) {
	s = &Service{}

	for _, cfg := range configs {
		if err = cfg(s); err != nil {
			return
		}
	}
	return
}

func WithRequestRepository(requestRepository request.Repository) Configuration {
	return func(s *Service) error {
		s.requestRepository = requestRepository
		return nil
	}
}

func WithResponseRepository(responseRepository response.Repository) Configuration {
	return func(s *Service) error {
		s.responseRepository = responseRepository
		return nil
	}
}
