package contracts

type UUID string

type IDGenerator interface {
	NewUUID() UUID
	ParseUUID(value string) (UUID, error)
}

func (u UUID) String() string {
	return string(u)
}
