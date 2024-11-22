package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOK  = "OK"
	StatusErr = "Error"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusErr,
		Error:  err.Error(),
	}
}

func ValidateError(err validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range err {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is required"), err.Field())
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is invalid"), err.Field())
		}
	}
	return Response{
		Status: StatusErr,
		Error:  strings.Join(errMsgs, ", "),
	}
}
