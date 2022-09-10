package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sky0621/cv-admin/generated/swagger"
)

type ServerImpl struct {
}

func (s *ServerImpl) GetUsersUserIdAttributes(w http.ResponseWriter, r *http.Request, userId swagger.UserId) {
	fmt.Println(userId)
	w.WriteHeader(200)
	w.Write([]byte("OK2"))
}

func main() {
	si := &ServerImpl{}
	r := chi.NewRouter()
	r.Mount("/", swagger.Handler(si))
	http.ListenAndServe(":3000", r)
}
