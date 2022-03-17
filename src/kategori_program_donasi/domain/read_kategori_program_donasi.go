package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/kategori_program_donasi/domain/entity"
	"pemuda-peduli/src/kategori_program_donasi/domain/interfaces"
)

func FindKategoriProgramDonasi(ctx context.Context, repo interfaces.IKategoriProgramDonasiRepository, data *entity.KategoriProgramDonasiQueryEntity) (response []entity.KategoriProgramDonasiEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	return
}

func GetKategoriProgramDonasi(ctx context.Context, repo interfaces.IKategoriProgramDonasiRepository, id string) (response entity.KategoriProgramDonasiEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	return
}
