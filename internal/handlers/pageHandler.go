package handlers

import (
	"net/http"

	"github.com/gsrthegreat/CookIt/templates"
)

type PageHandler struct{}

func NewPageHandler() *PageHandler {
	return &PageHandler{}
}

func (h *PageHandler) HomePageHandler(w http.ResponseWriter, r *http.Request) {
	templates.Home().Render(r.Context(), w)
}

func (h *PageHandler) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	templates.Login().Render(r.Context(), w)
}

func (h *PageHandler) RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	templates.Register().Render(r.Context(), w)
}
