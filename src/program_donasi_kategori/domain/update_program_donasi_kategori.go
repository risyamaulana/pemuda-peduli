package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/program_donasi_kategori/domain/entity"
	"pemuda-peduli/src/program_donasi_kategori/domain/interfaces"
)

func UpdateProgramDonasiKategori(ctx context.Context, repo interfaces.IProgramDonasiKategoriRepository, data entity.ProgramDonasiKategoriEntity, id string) (response entity.ProgramDonasiKategoriEntity, err error) {
	// Check available daata
	_, err = GetProgramDonasiKategori(ctx, repo, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}

	response, err = repo.Update(ctx, data, id)
	response.IDPPCPProgramDonasiKategori = id
	return
}
