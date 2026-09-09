package sample

import "template/domain/sample"

func ToAppResponse(s sample.Sample) SampleResponse {
	return SampleResponse{
		ID:        s.ID().String(),
		Name:      s.Name(),
		Number:    s.Number(),
		CreatedAt: s.CreatedAt(),
	}
}
