package sample

import "template/domain/domainerror"

var (
	ErrInvalidID        = domainerror.New(domainerror.InvalidInput, "Invalid sample id")
	ErrNotFound         = domainerror.New(domainerror.NotFound, "sample not found")
	ErrInvalidCreatedAt = domainerror.New(domainerror.InvalidInput, "Invalid created at")
)
