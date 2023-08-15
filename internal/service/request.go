package service

import "github.com/aalmat/proxy/internal/domain/request"

func (s *Service) SaveRequest(id string, req request.Entity) (err error) {
	err = s.requestRepository.StoreRequest(id, req)

	return
}
