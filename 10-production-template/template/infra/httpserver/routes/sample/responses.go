package sample

import (
	"time"

	appsample "template/application/usecase/sample"
	"template/infra/httpserver/utils"
)

type SampleResponse struct {
	ID        string    `json:"id"`
	Name      *string   `json:"name,omitempty"`
	Number    *int      `json:"number,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

func ToApiResponse(r appsample.SampleResponse) SampleResponse {
	return SampleResponse{
		ID:        r.ID,
		Name:      r.Name,
		Number:    r.Number,
		CreatedAt: r.CreatedAt,
	}
}

type GetManySampleResponse = utils.PageResponse[SampleResponse]
