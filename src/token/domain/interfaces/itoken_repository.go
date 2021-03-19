package interfaces

import (
	"context"
	"pemuda-peduli/src/token/domain/entity"
)

type ITokenRepository interface {
	Insert(ctx context.Context, data *entity.TokenEntity) (err error)

	Update(ctx context.Context, data *entity.TokenEntity) (err error)

	CheckDevice(deviceID, deviceType string) (data entity.TokenEntity, err error)
	CheckTokenDevice(token, deviceID, deviceType string) (data entity.TokenEntity, err error)
}
