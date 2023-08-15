package service

import "github.com/aalmat/proxy/internal/domain/response"

func (s *Service) SaveResponse(resp response.Entity) (err error) {
	err = s.responseRepository.StoreResponse(resp)

	return
}
