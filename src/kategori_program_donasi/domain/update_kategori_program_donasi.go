package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/kategori_program_donasi/domain/entity"
	"pemuda-peduli/src/kategori_program_donasi/domain/interfaces"
	"time"
)

func UpdateKategoriProgramDonasi(ctx context.Context, repo interfaces.IKategoriProgramDonasiRepository, data entity.KategoriProgramDonasiEntity, id string) (response entity.KategoriProgramDonasiEntity, err error) {
	currentDate := time.Now().UTC()
	// Check available daata
	checkData, err := repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	if checkData.IsDeleted {
		err = errors.New("Can't update this data")
		return
	}

	checkData.Name = data.Name
	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = false

	response, err = repo.Update(ctx, checkData, id)
	return
}

func DeleteKategoriProgramDonasi(ctx context.Context, repo interfaces.IKategoriProgramDonasiRepository, id string) (response entity.KategoriProgramDonasiEntity, err error) {
	currentDate := time.Now().UTC()
	// Check available data
	checkData, err := repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}

	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = true
	response, err = repo.Update(ctx, checkData, id)
	return
}
