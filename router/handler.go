package router

import (
	"encoding/json"
	"net/http"

	"github.com/renosyah/RestApiExample/api"
)

type (
	HandlerFunc func(http.ResponseWriter, *http.Request) (interface{}, *api.Error)
)

func (fn HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var errs []string

	r.ParseForm()

	data, err := fn(w, r)
	if err != nil {
		errs = append(errs, err.Error())
		w.WriteHeader(err.StatusCode)
	}

	resp := api.Response{
		Data: data,
		BaseResponse: api.BaseResponse{
			Errors: errs,
		},
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		return
	}
}
