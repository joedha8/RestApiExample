package router

import (
	"net/http"

	"github.com/renosyah/RestApiExample/api"
	"github.com/renosyah/RestApiExample/model"
)

func HandlerPing(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	return (&model.PingData{}).GetPing(), nil
}
