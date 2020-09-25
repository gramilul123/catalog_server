package app

import (
	"catalog_server/models"
	"catalog_server/utils"
	"context"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/signin", "/login"}
		requestPath := r.URL.Path

		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)

				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {

			sendError(w, "Missing auth token")

			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {

			sendError(w, "Invalid/Malformed auth token")

			return
		}

		tokenPart := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {

			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {

			sendError(w, "Malformed authentication token")

			return
		}

		if !token.Valid {

			sendError(w, "Token is not valid")

			return
		}

		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func sendError(w http.ResponseWriter, error string) {
	response := make(map[string]interface{})
	response = utils.Message(false, error)
	w.WriteHeader(http.StatusForbidden)
	w.Header().Add("Content-Type", "application/json")
	utils.Respond(w, response)
}
