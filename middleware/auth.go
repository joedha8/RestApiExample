package middleware

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/joedha8/RestApiExample/api"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		time := r.Header.Get("time")

		if time == "" {
			api.RespondError(w, api.MessageUnauthorized, http.StatusUnauthorized)

			return
		}

		encriptBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			api.RespondError(w, api.MessageUnauthorized, http.StatusUnauthorized)

			return
		}

		jsonBody, errVlidate := validateBody(strings.Trim(string(encriptBody), `"`), time)
		if errVlidate != nil {
			api.RespondError(w, api.MessageUnauthorized, http.StatusUnauthorized)

			return
		}

		r.Body = ioutil.NopCloser(strings.NewReader(jsonBody))

		next.ServeHTTP(w, r)
	})
}
