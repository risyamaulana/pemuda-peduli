package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/program_donasi_kategori/domain/entity"
	"pemuda-peduli/src/program_donasi_kategori/domain/interfaces"
)

func FindProgramDonasiKategori(ctx context.Context, repo interfaces.IProgramDonasiKategoriRepository, data *entity.ProgramDonasiKategoriQueryEntity) (response []entity.ProgramDonasiKategoriEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	return
}

func GetProgramDonasiKategori(ctx context.Context, repo interfaces.IProgramDonasiKategoriRepository, id string) (response entity.ProgramDonasiKategoriEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	return
}
