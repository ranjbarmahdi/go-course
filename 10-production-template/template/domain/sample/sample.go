package sample

import (
	"template/domain/shared"
	valueobjects "template/domain/value-objects"
	"time"
)

type Sample struct {
	id        valueobjects.SampleID
	name      *string
	number    *int
	createdAt time.Time
}

func (s Sample) ID() valueobjects.SampleID {
	return s.id
}

func (s Sample) Name() *string {
	return shared.CloneOptionalString(s.name)
}

func (s Sample) Number() *int {
	return shared.CloneOptionalInt(s.number)
}

func (s Sample) CreatedAt() time.Time {
	return s.createdAt
}

func NewSample(
	id valueobjects.SampleID,
	name *string,
	number *int,
	createdAt time.Time,
) (Sample, error) {
	s := Sample{
		id:        id,
		name:      shared.CloneOptionalString(name),
		number:    shared.CloneOptionalInt(number),
		createdAt: createdAt,
	}

	err := ValidateSample(s)
	if err != nil {
		return Sample{}, err
	}

	return s, nil
}
