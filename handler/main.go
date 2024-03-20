package handler

import (
	"log"
	"net/http"

	"example.com/music-api/store"
	"github.com/gorilla/mux"
)

type Handler struct {
	store  store.Store
	Router *mux.Router
}

func New(store store.Store) *Handler {
	h := &Handler{
		store: store,
	}

	h.Router = mux.NewRouter()

	h.Router.Use(ContentTypeJson)

	h.handleSong()

	return h
}

func (h *Handler) Run() error {
	http.Handle("/", h.Router)
	log.Println("Listening on :8080")
	return http.ListenAndServe(":8080", nil)
}

func ContentTypeJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
