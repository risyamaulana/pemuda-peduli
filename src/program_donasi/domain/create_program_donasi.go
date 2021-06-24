package domain

import (
	"context"
	"log"
	"pemuda-peduli/src/program_donasi/domain/entity"
	"pemuda-peduli/src/program_donasi/domain/interfaces"
)

func CreateProgramDonasi(ctx context.Context, repo interfaces.IProgramDonasiRepository, data *entity.ProgramDonasiEntity, dataDetail *entity.ProgramDonasiDetailEntity) (response entity.ProgramDonasiEntity, err error) {
	err = repo.Insert(ctx, data)

	// Insert Detail
	dataDetail.IDPPCPProgramDonasi = data.IDPPCPProgramDonasi
	dataDetail.Tag = data.Tag
	if errDetail := repo.InsertDetail(ctx, dataDetail); errDetail != nil {
		log.Println("ERR Insert Detail: ", errDetail)
	}

	response, _ = GetProgramDonasi(ctx, repo, data.IDPPCPProgramDonasi)

	return
}
