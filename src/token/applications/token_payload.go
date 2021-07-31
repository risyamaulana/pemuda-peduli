package applications

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"pemuda-peduli/src/token/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type TokenPayload struct {
	Name       string `json:"name" valid:"required"`
	SecretKey  string `json:"secret_key" valid:"required"`
	DeviceID   string `json:"device_id" valid:"required"`
	DeviceType string `json:"device_type" valid:"required"`
}

type ReadToken struct {
	ID                  int64      `json:"id"`
	Name                string     `json:"name"`
	DeviceID            string     `json:"device_id"`
	DeviceType          string     `json:"device_type"`
	Token               string     `json:"token"`
	TokenExpired        time.Time  `json:"token_expired"`
	RefreshToken        string     `json:"refresh_token"`
	RefreshTokenExpired time.Time  `json:"refresh_token_expired"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
}

func GetPayload(body []byte) (payload TokenPayload, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r TokenPayload) Validate(ctx context.Context) (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}

	// Validate Name & Secret Key
	name := os.Getenv("TOKEN_NAME")
	secretKey := os.Getenv("TOKEN_SECRET_KEY")
	if name != r.Name || secretKey != r.SecretKey {
		err = errors.New("Token authorization failed, Name Or Secret key not valid")
	}

	return
}

func (r TokenPayload) ToEntity() (data entity.TokenEntity) {
	data = entity.TokenEntity{
		Name:       r.Name,
		DeviceID:   r.DeviceID,
		DeviceType: r.DeviceType,
		CreatedAt:  time.Now().UTC(),
	}
	return
}
func ToPayload(data entity.TokenEntity) (response ReadToken) {
	response = ReadToken{
		ID:                  data.ID,
		Name:                data.Name,
		DeviceID:            data.DeviceID,
		DeviceType:          data.DeviceType,
		Token:               data.Token,
		TokenExpired:        data.TokenExpired,
		RefreshToken:        data.RefreshToken,
		RefreshTokenExpired: data.RefreshTokenExpired,
		CreatedAt:           data.CreatedAt,
		UpdatedAt:           data.UpdatedAt,
	}
	return
}
