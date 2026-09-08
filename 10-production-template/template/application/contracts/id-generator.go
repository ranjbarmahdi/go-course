package contracts

type UUID string

type IDGenerator interface {
	NewUUID() UUID
	ParseUUID(value string) (UUID, error)
}
