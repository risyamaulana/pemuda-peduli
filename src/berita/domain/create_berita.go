package domain

import (
	"context"
	"pemuda-peduli/src/berita/domain/entity"
	"pemuda-peduli/src/berita/domain/interfaces"
)

func CreateBerita(ctx context.Context, repo interfaces.IBeritaRepository, data *entity.BeritaEntity) (err error) {
	err = repo.Insert(ctx, data)
	return
}
