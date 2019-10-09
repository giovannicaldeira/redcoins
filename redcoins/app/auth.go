package app

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"redcoins/models"
	u "redcoins/utils"

	jwt "github.com/dgrijalva/jwt-go"
)

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		noAuth := []string{"/redcoins/api/user/new", "/redcoins/api/login", "/redcoins/api/user/reset_password, /"}
		requestPath := r.URL.Path

		for _, value := range noAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Auth")

		if tokenHeader == "" {
			response = u.Message(false, "Token de autenticação invalido")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		// tokenSplitted := strings.Split(tokenHeader, " ")
		// if len(tokenSplitted) != 2 {
		// 	response = u.Message(false, "Token de autenticação invalido")
		// 	w.WriteHeader(http.StatusForbidden)
		// 	w.Header().Add("Content-Type", "application/json")
		// 	u.Respond(w, response)
		// 	return
		// }

		// tokenPart := tokenSplitted[1]

		tokenPart := tokenHeader
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN_PASSWORD")), nil
		})

		if err != nil {
			response = u.Message(false, "Token de autenticacao invalido")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		if !token.Valid {
			response = u.Message(false, "Token de autenticao invalido.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		fmt.Sprintf("User %", tk.UserId)
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
