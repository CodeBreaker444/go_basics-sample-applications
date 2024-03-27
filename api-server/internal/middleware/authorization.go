package middleware

import (
	"apiserver/api"
	"errors"
	"fmt"
	"net/http"
	log "github.com/sirupsen/logrus"
)

var unAuthorisedError = errors.New("Invalid token")
func Authorization(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token=r.Header.Get("Authorization")
		fmt.Printf("\nToken:%v",token)
		//var err error
		if token==""{
			log.Error(unAuthorisedError)
			api.RequestErrorHandler(w, unAuthorisedError)
			return
		}

		next.ServeHTTP(w,r)
	})
}