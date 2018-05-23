package got

import (
	"github.com/google/uuid"
)

func uniqueID() string {
	return uuid.New().String()
}
