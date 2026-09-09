package valueobjects

import "template/domain/domainerror"

type SampleID struct {
	value string
}

func NewSampleID(value string) (SampleID, error) {
	if err := parseUUID(value); err != nil {
		return SampleID{}, domainerror.New(domainerror.InvalidInput, "invalid sample id")
	}
	return SampleID{value: value}, nil
}

func SampleIDFromTrusted(value string) SampleID {
	return SampleID{value: value}
}

func (id SampleID) String() string { return id.value }
func (id SampleID) IsZero() bool   { return id.value == "" }
