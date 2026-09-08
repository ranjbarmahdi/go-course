package adapters

import (
	"github.com/google/uuid"

	"template/application/contracts"
)

type UUIDGenerator struct{}

func (UUIDGenerator) NewUUID() contracts.UUID {
	return contracts.UUID(uuid.NewString())
}

func (UUIDGenerator) ParseUUID(value string) (contracts.UUID, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return "", err
	}
	return contracts.UUID(id.String()), nil
}
