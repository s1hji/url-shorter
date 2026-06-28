package handlers

import (
	"net/http"
	"strings"
	"url-shorter/storage"
)

type RedirectHandler struct {
	Storage storage.Storage
}

func (re *RedirectHandler) RedirectServeHttp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "неверный метод", http.StatusMethodNotAllowed)
		return
	}

	short := strings.TrimPrefix(r.URL.Path, "/")
	if short == "" {
		http.Error(w, "такой короткой ссылки не сущетсвует", http.StatusBadRequest)
		return
	}

	origin, err := re.Storage.GetOriginLink(short)
	if err != nil {
		http.Error(w, "ссылка не найдена", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, origin, http.StatusFound)
}
