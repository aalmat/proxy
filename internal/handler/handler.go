package handler

import (
	"github.com/aalmat/proxy/internal/config"
	"github.com/aalmat/proxy/internal/handler/http"
	"github.com/aalmat/proxy/internal/service"
	"github.com/go-chi/chi/v5"
)

type Dependencies struct {
	Configs config.Configs
	Service *service.Service
}

type Configuration func(h *Handler) error

type Handler struct {
	dependencies Dependencies
	HTTP         *chi.Mux
}

func New(d Dependencies, configs ...Configuration) (h *Handler, err error) {
	h = &Handler{
		dependencies: d,
	}

	for _, cfg := range configs {
		if err = cfg(h); err != nil {
			return
		}
	}

	return
}

func WithHTTPHandler() Configuration {
	return func(h *Handler) (err error) {
		h.HTTP = chi.NewRouter()

		proxyHandler := http.NewProxyHandler(h.dependencies.Service)
		h.HTTP.Route("/api/v1", func(r chi.Router) {
			r.Mount("/proxy", proxyHandler.Routes())
		})

		return
	}
}
