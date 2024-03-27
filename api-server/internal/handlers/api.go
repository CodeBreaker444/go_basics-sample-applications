package handlers

import (
	"apiserver/internal/middleware"
	"encoding/json"
	"fmt"
	"net/http"
	"apiserver/api"
	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler_call(r *chi.Mux){
	r.Use(chimiddle.StripSlashes)
	r.Route("/details",func(router chi.Router){
		router.Use(middleware.Authorization)
		router.Get("/decodedToken",decodedToken)
	})
	fmt.Printf("handler called")
}
func decodedToken(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Decoded Token")
	var err error
	var resp map[string](string) = map[string]string{"token":"decodedToken"}
	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err!=nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	
}