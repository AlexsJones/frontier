package v1

import "net/http"

func ProcessorTestPOST(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}
