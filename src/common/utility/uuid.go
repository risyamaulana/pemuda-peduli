package utility

import "github.com/google/uuid"

func GetUUID() (uid string) {
	uid = uuid.New().String()
	return
}
