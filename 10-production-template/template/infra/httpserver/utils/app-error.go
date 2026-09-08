package utils

import (
	"log/slog"
	"net/http"

	apperrors "template/application/errors"
)

func WriteAppError(w http.ResponseWriter, err error) {
	switch apperrors.KindOf(err) {
	case apperrors.InvalidInput:
		WriteError(w, http.StatusBadRequest, err.Error())

	case apperrors.Unauthorized:
		WriteError(w, http.StatusUnauthorized, err.Error())

	case apperrors.Domain:
		WriteError(w, http.StatusBadRequest, err.Error())

	case apperrors.NotFound:
		WriteError(w, http.StatusNotFound, err.Error())

	case apperrors.Conflict:
		WriteError(w, http.StatusConflict, err.Error())

	case apperrors.Internal:
		slog.Error("request failed", "err", err)
		WriteError(w, http.StatusInternalServerError, "internal server error")

	default:
		slog.Error("request failed", "err", err)
		WriteError(w, http.StatusInternalServerError, "internal server error")
	}
}
