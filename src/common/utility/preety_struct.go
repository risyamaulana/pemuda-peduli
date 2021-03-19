package utility

import (
	"encoding/json"
	"log"
)

// PrettyPrint will transform struct data as json string for nicer log
func PrettyPrint(data interface{}) string {
	JSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	return string(JSON)
}
