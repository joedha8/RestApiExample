package middleware

import (
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

			return
		}

		encriptBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			api.RespondError(w, api.MessageUnauthorized, http.StatusUnauthorized)

			return
		}

		jsonBody, errVlidate := validateBody(string(encriptBody), time, key)
		if errVlidate != nil {
			api.RespondError(w, api.MessageUnauthorized, http.StatusUnauthorized)

			return
		}

		r.Body = ioutil.NopCloser(strings.NewReader(jsonBody))

		next.ServeHTTP(w, r)
	})
}
