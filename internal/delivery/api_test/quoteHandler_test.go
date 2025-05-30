package api_test_test

import (
	"AiCheto/internal/delivery/api"
	"AiCheto/internal/repositories"
	"AiCheto/internal/usecases"
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func router() *mux.Router {
	repo := repositories.New()
	uc := usecases.NewQuoteUsecase(repo)
	h := api.NewQuoteHandler(uc)

	r := mux.NewRouter()
	h.Register(r)
	return r
}

func TestCreateAndList(t *testing.T) {
	srv := httptest.NewServer(router())
	defer srv.Close()

	// POST /quotes
	body := []byte(`{"author":"Test","quote":"Hello"}`)
	resp, err := http.Post(srv.URL+"/quotes", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("post failed: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("want 201, got %d", resp.StatusCode)
	}

	// GET /quotes?author=Test
	res, err := http.Get(srv.URL + "/quotes?author=Test")
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("want 200, got %d", res.StatusCode)
	}
	var quotes []map[string]any
	if err := json.NewDecoder(res.Body).Decode(&quotes); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(quotes) != 1 {
		t.Fatalf("expect 1 quote, got %d", len(quotes))
	}
}

func TestRandomAndDelete(t *testing.T) {
	srv := httptest.NewServer(router())
	defer srv.Close()

	// создаём две цитаты
	for _, txt := range []string{"One", "Two"} {
		_, _ = http.Post(srv.URL+"/quotes", "application/json",
			bytes.NewReader([]byte(`{"author":"Tester","quote":"`+txt+`"}`)))
	}

	// GET /quotes/random
	res, _ := http.Get(srv.URL + "/quotes/random")
	var q map[string]string
	_ = json.NewDecoder(res.Body).Decode(&q)

	// DELETE /quotes/{id}
	req, _ := http.NewRequest(http.MethodDelete, srv.URL+"/quotes/"+q["id"], nil)
	delResp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("delete failed: %v", err)
	}
	if delResp.StatusCode != http.StatusNoContent {
		t.Fatalf("want 204, got %d", delResp.StatusCode)
	}
}
