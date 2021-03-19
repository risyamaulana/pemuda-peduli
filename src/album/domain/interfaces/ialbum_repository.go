package interfaces

import (
	"context"
	"pemuda-peduli/src/album/domain/entity"
)

type IAlbumRepository interface {
	Insert(ctx context.Context, data *entity.AlbumEntity) (err error)

	Update(ctx context.Context, data entity.AlbumEntity, id string) (response entity.AlbumEntity, err error)

	Find(ctx context.Context, data *entity.AlbumQueryEntity) (company []entity.AlbumEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.AlbumEntity, err error)
}
