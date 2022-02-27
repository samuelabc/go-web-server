package healthRoute

import (
	"encoding/json"
	"fmt"
	"net/http"

	errorModel "web-server/model/error"
	contextHelper "web-server/src/helper/context"
	errorHelper "web-server/src/helper/error"
)

func PostHealth(w http.ResponseWriter, r *http.Request) *errorModel.AppError {
	var err error

	payload := r.Context().Value(contextHelper.ContextKeyJSONPayload)
	fmt.Println("in controller, payload:", payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(payload)
	if err != nil {
		return errorHelper.DECODE_ERROR(err)
	}
	return nil
}
