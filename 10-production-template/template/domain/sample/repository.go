package sample

import (
	"context"
	valueobjects "template/domain/value-objects"
)

type SampleRepository interface {
	CreateSample(ctx context.Context, sample Sample) error
	FindByID(ctx context.Context, id valueobjects.SampleID) (*Sample, error)
}
