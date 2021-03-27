package applications

import (
	"context"
	"encoding/json"

	"github.com/asaskevich/govalidator"
)

type Login struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required,minstringlength(8)"`
}

func GetLoginPayload(body []byte) (payload Login, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r *Login) Validate(ctx context.Context) (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}

	return
}
