package random

import (
	"github.com/google/uuid"
)

func SecUUID() string {
	u, err := uuid.NewUUID()
	if err != nil {
		u = uuid.New()
	}
	return u.String()
}
