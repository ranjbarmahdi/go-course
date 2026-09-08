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
		return apperrors.NewError(apperrors.InvalidInput, err.Error())

	case domainerror.NotFound:
		return apperrors.NewError(apperrors.NotFound, err.Error())

	case domainerror.Conflict:
		return apperrors.NewError(apperrors.Conflict, err.Error())

	case domainerror.Domain:
		return apperrors.NewError(apperrors.Domain, err.Error())

	case domainerror.Internal:
		return apperrors.NewError(apperrors.Internal, err.Error())

	default:
		slog.Error("unmapped domain error", "err", err)
		return apperrors.NewError(apperrors.Internal, "unexpected error")
	}
}
