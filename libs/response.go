package libs

import (
	"encoding/json"
	"net/http"
)

/*
 * handle API Response & Error
 * agar response API lebih rapi dan mudah dikembangkan
 *
 */
func HandleResponse(httpCode int, w http.ResponseWriter, data interface{}, message ...string) {
    msg := ""
    if len(message) > 0 {
        msg = message[0]
    }

    var output = map[string]interface{}{
        "status": httpCode,
        "message": msg,
        "data": data,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(httpCode)
    json.NewEncoder(w).Encode(output)
}

func HandleError(httpCode int, w http.ResponseWriter, message string) {
    var output = map[string]interface{}{
        "status": httpCode,
        "message": message,
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(httpCode)
    json.NewEncoder(w).Encode(output)
}