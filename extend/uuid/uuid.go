package sUuid

import (
	"github.com/google/uuid"
)

func NextUUID() uuid.UUID {
	return uuid.New()
}

func Next() string {
	return uuid.New().String()
}
