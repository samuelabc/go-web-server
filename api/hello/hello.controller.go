package helloRoute

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	errorHelper "web-server/helper/error"
	errorModel "web-server/model/error"
	helloModel "web-server/model/hello"
)

func PostHello(w http.ResponseWriter, r *http.Request) *errorModel.AppError {
	var err error
	var payload PostHelloSchema
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return errorHelper.ErrDecode(err)
	}
	// payload := r.Context().Value(contextHelper.ContextKeyJSONPayload).(*PostHelloSchema)

	var response helloModel.HelloResponse
	response.Name = payload.Name
	response.Age = payload.Age
	response.Description = fmt.Sprintf("Hi, my name is %s and I'm %d years old.", payload.Name, payload.Age)
	response.Timestamp = time.Now().String()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return errorHelper.ErrEncode(err)
	}
	return nil
}
