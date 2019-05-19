package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/renosyah/RestApiExample/api"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		key := r.Header.Get("key")
		time := r.Header.Get("time")

		if key == "" || time == "" {
			api.RespondError(w, api.MessageUnauthorized, http.StatusUnauthorized)
			fmt.Println("key and time empty")
			return
		}

		hashBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			api.RespondError(w, api.MessageUnauthorized, http.StatusUnauthorized)
			fmt.Println("body empty")
			return
		}

		jsonBody := compareHashRequest(key, string(hashBody), time)
		if jsonBody == "" {
			api.RespondError(w, api.MessageUnauthorized, http.StatusUnauthorized)
			fmt.Println("failed compare")
			return
		}

		r.Body = ioutil.NopCloser(strings.NewReader(jsonBody))

		next.ServeHTTP(w, r)
	})
}
