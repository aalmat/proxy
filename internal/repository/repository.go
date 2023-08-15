package repository

import (
	"github.com/aalmat/proxy/internal/domain/request"
	"github.com/aalmat/proxy/internal/domain/response"
	"github.com/aalmat/proxy/internal/repository/memory"
)

type Configuration func(r *Repository) error

type Repository struct {
	Request  request.Repository
	Response response.Repository
}

func New(configs ...Configuration) (s *Repository, err error) {
	s = &Repository{}

	for _, cfg := range configs {
		if err = cfg(s); err != nil {
			return
		}
	}

	return
}

func WithMemoryStore() Configuration {
	return func(s *Repository) (err error) {
		s.Response = memory.NewResponseRepository()
		s.Request = memory.NewRequestRepository()

		return
	}
}
