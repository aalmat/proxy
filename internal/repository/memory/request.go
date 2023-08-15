package memory

import (
	"github.com/aalmat/proxy/internal/domain/request"
	"sync"
)

type RequestRepository struct {
	db map[string]request.Entity
	sync.RWMutex
}

func (r *RequestRepository) StoreRequest(id string, request request.Entity) error {
	r.Lock()
	defer r.Unlock()

	r.db[id] = request

	return nil
}

func NewRequestRepository() *RequestRepository {
	return &RequestRepository{
		db: make(map[string]request.Entity),
	}
}
