package request

type Repository interface {
	StoreRequest(id string, req Entity) error
}
