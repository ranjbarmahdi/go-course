package sample

import "time"

type SampleResponse struct {
	ID        string
	Name      *string
	Number    *int
	CreatedAt time.Time
}
