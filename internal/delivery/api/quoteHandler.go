package api

import (
	"AiCheto/internal/entity"
	"AiCheto/internal/usecases"

	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// 		--- DTO ---

// CreateQuoteRequest описывает тело POST /quotes.
type CreateQuoteRequest struct {
	Author string `json:"author" example:"Confucius"`
	Quote  string `json:"quote"  example:"Life is really simple, but we insist on making it complicated."`
}

// 		--- handler ---

type QuoteHandler struct{ uc *usecases.QuoteService }

func NewQuoteHandler(uc *usecases.QuoteService) *QuoteHandler { return &QuoteHandler{uc: uc} }

func (h *QuoteHandler) Register(r *mux.Router) {
	r.HandleFunc("/quotes", h.create).Methods(http.MethodPost)
	r.HandleFunc("/quotes", h.list).Methods(http.MethodGet)
	r.HandleFunc("/quotes/random", h.random).Methods(http.MethodGet)
	r.HandleFunc("/quotes/{id}", h.delete).Methods(http.MethodDelete)
}

// create godoc
// @Summary      Create a quote
// @Description  Adds a new quote (ID генерируется на сервере)
// @Tags         quotes
// @Accept       json
// @Produce      json
// @Param        payload  body      api.CreateQuoteRequest  true  "Quote to create"
// @Success      201      {object}  entity.Quote
// @Failure      400      {string}  string  "invalid request"
// @Router       /quotes [post]
func (h *QuoteHandler) create(w http.ResponseWriter, r *http.Request) {
	var req CreateQuoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Author == "" || req.Quote == "" {
		http.Error(w, "author and quote are required", http.StatusBadRequest)
		return
	}

	q := &entity.Quote{Author: req.Author, Text: req.Quote}
	if err := h.uc.Create(r.Context(), q); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// сделал логи для более удобного тестирования через curl
	log.Printf("created quote id=%s author=%q", q.ID, q.Author)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", "/quotes/"+q.ID.String())
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(q)
}

// list godoc
// @Summary      List quotes
// @Description  Returns all quotes or filtered by author (case-insensitive)
// @Tags         quotes
// @Produce      json
// @Param        author  query     string  false  "Author to filter"
// @Success      200     {array}   entity.Quote
// @Router       /quotes [get]
func (h *QuoteHandler) list(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	quotes, err := h.uc.List(r.Context(), author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(quotes)
}

// random godoc
// @Summary      Random quote
// @Tags         quotes
// @Produce      json
// @Success      200  {object}  entity.Quote
// @Failure      500  {string}  string  "no quotes available"
// @Router       /quotes/random [get]
func (h *QuoteHandler) random(w http.ResponseWriter, r *http.Request) {
	q, err := h.uc.Random(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(q)
}

// delete godoc
// @Summary      Delete quote
// @Tags         quotes
// @Param        id   path      string  true  "Quote ID (UUID)"
// @Success      204  "deleted"
// @Failure      400  {string}  string  "invalid id"
// @Failure      404  {string}  string  "quote not found"
// @Router       /quotes/{id} [delete]
func (h *QuoteHandler) delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	if err := h.uc.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
