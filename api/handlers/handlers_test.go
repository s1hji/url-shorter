package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shorter/api/models"
	"url-shorter/storage/memory"
)

func TestShorterHandler(t *testing.T) {
	storage := ShortenHandler{
		Store: memory.NewMemoryStorage(),
	}

	body := bytes.NewBufferString(`{"url":"telegram.org"}`)

	req := httptest.NewRequest(http.MethodPost, "/shorter", body)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	storage.ServeHttp(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("ожидали 200 OK, получили %d", w.Code)
	}

}

func TestDuplicateUrl(t *testing.T) {
	storage := ShortenHandler{
		Store: memory.NewMemoryStorage(),
	}

	body1 := bytes.NewBufferString(`{"url":"telegram.org"}`)
	req1 := httptest.NewRequest(http.MethodPost, "/shorter", body1)
	req1.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	storage.ServeHttp(w1, req1)

	var resp1 models.Response
	json.NewDecoder(w1.Body).Decode(&resp1)

	body2 := bytes.NewBufferString(`{"url":"telegram.org"}`)
	req2 := httptest.NewRequest(http.MethodPost, "/shorter", body2)
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	storage.ServeHttp(w2, req2)

	var resp2 models.Response
	json.NewDecoder(w2.Body).Decode(&resp2)

	if resp1 != resp2 {
		t.Errorf("должно было вывести одинаковый short, но вывело %s %s", resp1.Short, resp2.Short)
	}
}

func TestRedirectHandler(t *testing.T) {
	storage := memory.NewMemoryStorage()

	storage.Save("telegram.org", "tg")

	s := RedirectHandler{
		Storage: storage,
	}

	req := httptest.NewRequest(http.MethodGet, "/tg", nil)
	w := httptest.NewRecorder()
	s.RedirectServeHttp(w, req)

	if w.Code != http.StatusFound {
		t.Errorf("ожидали 302 found, получили %d", w.Code)
	}
}
