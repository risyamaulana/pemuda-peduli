package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/program_donasi/domain/entity"
	"pemuda-peduli/src/program_donasi/domain/interfaces"
)

func FindProgramDonasi(ctx context.Context, repo interfaces.IProgramDonasiRepository, data *entity.ProgramDonasiQueryEntity) (response []entity.ProgramDonasiEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	for i, data := range response {
		// Get Detail
		dataDetail, _ := repo.GetDetail(ctx, data.IDPPCPProgramDonasi)

		response[i].Detail = dataDetail
	}
	return
}

func GetProgramDonasi(ctx context.Context, repo interfaces.IProgramDonasiRepository, id string) (response entity.ProgramDonasiEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}

	// Get Detail
	dataDetail, _ := repo.GetDetail(ctx, response.IDPPCPProgramDonasi)

	response.Detail = dataDetail
	return
}
