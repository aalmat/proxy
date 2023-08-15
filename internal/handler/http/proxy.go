package http

import (
	"github.com/aalmat/proxy/internal/domain/request"
	"github.com/aalmat/proxy/internal/domain/response"
	"github.com/aalmat/proxy/internal/service"
	httpResponse "github.com/aalmat/proxy/pkg/server/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"io"
	"net/http"
)

type ProxyHandler struct {
	service *service.Service
}

func NewProxyHandler(service *service.Service) *ProxyHandler {
	return &ProxyHandler{
		service: service,
	}
}

func (h *ProxyHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.do)

	return r
}

func (h *ProxyHandler) do(w http.ResponseWriter, r *http.Request) {
	req := request.Entity{}
	if err := render.Bind(r, &req); err != nil {
		httpResponse.BadRequest(w, r, err, nil)
		return
	}

	res, err := DoProxyRequest(&req)
	if err != nil {
		httpResponse.BadRequest(w, r, err, nil)
		return
	}

	err = h.service.SaveRequest(res.ID, req)
	if err != nil {
		httpResponse.InternalServerError(w, r, err)
		return
	}

	err = h.service.SaveResponse(*res)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	httpResponse.OK(w, r, map[string]interface{}{
		"data": res,
	})

}

func DoProxyRequest(req *request.Entity) (*response.Entity, error) {
	client := &http.Client{}
	request, err := http.NewRequest(req.Method, req.Url, nil)

	if err != nil {
		return nil, err
	}

	for k, v := range req.Headers {
		request.Header.Set(k, v)
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	for k, v := range resp.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}

	id := uuid.New().String()

	proxyResponse := response.Entity{
		ID:      id,
		Status:  resp.Status,
		Headers: headers,
		Length:  uint64(len(respBody)),
	}

	return &proxyResponse, nil

}
