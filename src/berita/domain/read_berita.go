package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/berita/domain/entity"
	"pemuda-peduli/src/berita/domain/interfaces"
)

func FindBerita(ctx context.Context, repo interfaces.IBeritaRepository, data *entity.BeritaQueryEntity) (response []entity.BeritaEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	for i, data := range response {
		// Get Detail
		dataDetail, _ := repo.GetDetail(ctx, data.IDPPCPBerita)

		response[i].Detail = dataDetail
	}
	return
}

func GetBerita(ctx context.Context, repo interfaces.IBeritaRepository, id string) (response entity.BeritaEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}

	// Get Detail
	dataDetail, _ := repo.GetDetail(ctx, response.IDPPCPBerita)

	response.Detail = dataDetail
	return
}
