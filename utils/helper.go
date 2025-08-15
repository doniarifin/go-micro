package helper

import (
	"github.com/google/uuid"
)

func NewUUID() string {
	newUid := uuid.New()
	return newUid.String()
}
