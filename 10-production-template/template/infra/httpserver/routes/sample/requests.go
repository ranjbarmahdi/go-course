package sample

type CreateSampleRequest struct {
	Name   string `json:"name" validate:"required"`
	Number int    `json:"number" validate:"required"`
}
