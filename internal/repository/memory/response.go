package memory

import (
	"github.com/aalmat/proxy/internal/domain/response"
	"sync"
)

type ResponseRepository struct {
	db map[string]response.Entity
	sync.RWMutex
}

func (r *ResponseRepository) StoreResponse(response response.Entity) error {
	r.Lock()
	defer r.Unlock()

	r.db[response.ID] = response

	return nil
}

func NewResponseRepository() *ResponseRepository {
	return &ResponseRepository{
		db: make(map[string]response.Entity),
	}
}
