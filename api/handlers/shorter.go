package handlers

import (
	"encoding/json"
	"net/http"
	"url-shorter/api/models"
	"url-shorter/generator"
	"url-shorter/storage"
)

type ShortenHandler struct {
	Store storage.Storage
}

func (s *ShortenHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "неверный метод", http.StatusMethodNotAllowed)
		return
	}

	var request models.Request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "ошибка записи json", http.StatusBadRequest)
		return
	}

	short, err := s.Store.GetShortLink(request.URL)

	if err == nil {
		response := models.Response{
			Short: short,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	for {
		short, _ = generator.Rand_generate()

		_, err := s.Store.GetOriginLink(short)

		if err != nil {
			break
		}

	}

	err = s.Store.Save(request.URL, short)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.Response{
		Short: short,
	}
	json.NewEncoder(w).Encode(response)

}
