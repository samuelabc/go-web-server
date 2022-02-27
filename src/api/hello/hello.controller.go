package helloRoute

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	errorModel "web-server/model/error"
	helloModel "web-server/model/hello"
	contextHelper "web-server/src/helper/context"
	errorHelper "web-server/src/helper/error"
)

func PostHello(w http.ResponseWriter, r *http.Request) *errorModel.AppError {
	var err error

	// var payload PostHelloSchema
	// err := json.NewDecoder(r.Body).Decode(&payload)
	// if err != nil {
	// 	return errorHelper.DECODE_ERROR(err)
	// }
	payload := r.Context().Value(contextHelper.ContextKeyJSONPayload).(*PostHelloSchema)

	var response helloModel.HelloResponse
	response.Name = payload.Name
	response.Age = payload.Age
	response.Description = fmt.Sprintf("Hi, my name is %s and I'm %d years old.", payload.Name, payload.Age)
	response.Timestamp = time.Now().String()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return errorHelper.ENCODE_ERROR(err)
	}
	return nil
}
