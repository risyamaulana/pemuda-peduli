package domain

import (
	"context"
	"pemuda-peduli/src/album/domain/entity"
	"pemuda-peduli/src/album/domain/interfaces"
)

func CreateAlbum(ctx context.Context, repo interfaces.IAlbumRepository, data *entity.AlbumEntity) (err error) {
	err = repo.Insert(ctx, data)
	return
}
