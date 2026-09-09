package sample

func ValidateSample(s Sample) error {
	if s.ID().IsZero() {
		return ErrInvalidID
	}
	if s.CreatedAt().IsZero() {
		return ErrInvalidCreatedAt
	}
	return nil
}
