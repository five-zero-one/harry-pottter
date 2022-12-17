package service

import (
	"encoding/json"
	"harry-potter/store"
	"harry-potter/store/option"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

const applicationJson = "application/json"

type Service struct {
	mux chi.Router

	r store.Repo
}

func New(r store.Repo) http.Handler {
	s := &Service{
		mux: chi.NewMux(),
		r:   r,
	}
	s.routes()
	return s
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Service) routes() {
	s.mux.Get("/api/characters", s.handleCharacterList)
	s.mux.Get("/api/characters/{characterId}", s.handleUniqueCharacter)
}

func (s *Service) handleCharacterList(w http.ResponseWriter, r *http.Request) {
	opts := s.parseFilterOptions(r)
	cs, err := s.r.Filter(opts)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s.respond(w, cs, http.StatusOK)
}

func (s *Service) handleUniqueCharacter(w http.ResponseWriter, r *http.Request) {
	id, err := s.parseCharacterID(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c, err := s.r.Search(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	s.respond(w, c, http.StatusOK)
}

func (s *Service) parseFilterOptions(r *http.Request) option.FilterOption {
	return option.NewFilter(r.URL.Query())
}

func (s Service) parseCharacterID(r *http.Request) (int, error) {
	p := chi.URLParam(r, "characterId")
	return strconv.Atoi(p)
}

func (s Service) respond(w http.ResponseWriter, v any, status int) {
	w.Header().Set("Content-Type", applicationJson)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
