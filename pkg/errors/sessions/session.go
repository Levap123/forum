package sessions

import "github.com/gofrs/uuid"

func GenerateUuid() (string, error) {
	uuid, _ := uuid.NewV7()

	return uuid.String(), nil
}
