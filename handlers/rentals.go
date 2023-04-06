package handlers

import "net/http"

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type rentalHandler struct {
}

func InitializeHandler() Handler {
	return &rentalHandler{}
}

func (rh *rentalHandler) Get(w http.ResponseWriter, r *http.Request) {

}

func (rh *rentalHandler) List(w http.ResponseWriter, r *http.Request) {

}
