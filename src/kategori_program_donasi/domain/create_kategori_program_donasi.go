package domain

import (
	"context"
	"pemuda-peduli/src/kategori_program_donasi/domain/entity"
	"pemuda-peduli/src/kategori_program_donasi/domain/interfaces"
)

func CreateKategoriProgramDonasi(ctx context.Context, repo interfaces.IKategoriProgramDonasiRepository, data *entity.KategoriProgramDonasiEntity) (err error) {
	err = repo.Insert(ctx, data)
	return
}
