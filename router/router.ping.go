package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/joedha8/RestApiExample/api"
	"github.com/joedha8/RestApiExample/model"
)

func HandlerPing(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {

	response := &model.PingData{}

	body, errReadBody := ioutil.ReadAll(r.Body)
	if errReadBody != nil {
		return response, &api.Error{Err: errReadBody, StatusCode: 0, Message: ""}
	}

	data := &model.PingData{}
	errUnmarshal := json.Unmarshal(body, data)
	if errUnmarshal != nil {
		return response, &api.Error{Err: errUnmarshal, StatusCode: 0, Message: ""}
	}

	if data.PingData == "ping" {
		response = data.GetPing()
	}

	return response, nil
}
