package home

import "net/http"

type handler struct {
	fs http.FileSystem
}

func NewHandler(fs http.FileSystem) *handler {
	return &handler{fs: fs}
}

func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
