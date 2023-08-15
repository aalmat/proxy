package response

type Entity struct {
	ID      string            `json:"id"`
	Status  string            `json:"status"`
	Headers map[string]string `json:"headers"`
	Length  uint64            `json:"length"`
}
