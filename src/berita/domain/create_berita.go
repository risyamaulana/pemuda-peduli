package domain

import (
	"context"
	"log"
	"pemuda-peduli/src/berita/domain/entity"
	"pemuda-peduli/src/berita/domain/interfaces"
)

func CreateBerita(ctx context.Context, repo interfaces.IBeritaRepository, data *entity.BeritaEntity, dataDetail *entity.BeritaDetailEntity) (response entity.BeritaEntity, err error) {
	err = repo.Insert(ctx, data)
	if err != nil {
		return
	}

	// Insert Detail
	dataDetail.IDPPCPBerita = data.IDPPCPBerita
	dataDetail.Tag = data.Tag
	if errDetail := repo.InsertDetail(ctx, dataDetail); errDetail != nil {
		log.Println("ERR Insert Detail: ", errDetail)
	}

	response, _ = GetBerita(ctx, repo, data.IDPPCPBerita)

	return
}
