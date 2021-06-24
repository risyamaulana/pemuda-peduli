package domain

import (
	"context"
	"pemuda-peduli/src/program_donasi_kategori/domain/entity"
	"pemuda-peduli/src/program_donasi_kategori/domain/interfaces"
)

func CreateProgramDonasiKategori(ctx context.Context, repo interfaces.IProgramDonasiKategoriRepository, data *entity.ProgramDonasiKategoriEntity) (err error) {
	err = repo.Insert(ctx, data)
	return
}
