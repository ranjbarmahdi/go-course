package utils

import (
	"log/slog"

	apperrors "template/application/errors"
	"template/domain/domainerror"
)

func TranslateDomainError(err error) error {
	if err == nil {
		return nil
	}

	switch domainerror.KindOf(err) {
	case domainerror.InvalidInput:
		return apperrors.New(apperrors.InvalidInput, err.Error())

	case domainerror.NotFound:
		return apperrors.New(apperrors.NotFound, err.Error())

	case domainerror.Conflict:
		return apperrors.New(apperrors.Conflict, err.Error())

	case domainerror.Domain:
		return apperrors.New(apperrors.Domain, err.Error())

	case domainerror.Internal:
		return apperrors.New(apperrors.Internal, err.Error())

	default:
		slog.Error("unmapped domain error", "err", err)
		return apperrors.New(apperrors.Internal, "unexpected error")
	}
}
