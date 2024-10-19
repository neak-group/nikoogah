package uuid

import "github.com/google/uuid"

type UUID struct{ uuid.UUID }

func New() UUID {
	return UUID{
		UUID: uuid.New(),
	}
}

var Nil = UUID{uuid.Nil}

func Parse(str string) (u UUID, err error) {
	u = UUID{}
	u.UUID, err = uuid.Parse(str)
	return u, err
}

type UUIDs []UUID

func NewString() string {
	return uuid.NewString()
}
