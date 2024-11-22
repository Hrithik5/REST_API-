package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/hrithik5/student-api/internal/types"
	"github.com/hrithik5/student-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Creating a Studnet")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// Validate Request
		if err := validator.New().Struct(student); err != nil {

			validateErrs := err.(validator.ValidationErrors)
			response.WriteJSON(w, http.StatusBadRequest, response.ValidateError(validateErrs))
			return
		}

		response.WriteJSON(w, http.StatusCreated, map[string]string{"sucess": "OK"})
	}
}
