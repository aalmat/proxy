package response

type Repository interface {
	StoreResponse(response Entity) error
}
