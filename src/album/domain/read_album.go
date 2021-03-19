package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/album/domain/entity"
	"pemuda-peduli/src/album/domain/interfaces"
)

func FindAlbum(ctx context.Context, repo interfaces.IAlbumRepository, data *entity.AlbumQueryEntity) (response []entity.AlbumEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	return
}

func GetAlbum(ctx context.Context, repo interfaces.IAlbumRepository, id string) (response entity.AlbumEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	return
}
