package domain

import (
	"context"
	"log"
	"pemuda-peduli/src/program_kami/domain/entity"
	"pemuda-peduli/src/program_kami/domain/interfaces"
)

func CreateProgramKami(ctx context.Context, repo interfaces.IProgramKamiRepository, data *entity.ProgramKamiEntity, dataDetail *entity.ProgramKamiDetailEntity) (response entity.ProgramKamiEntity, err error) {
	err = repo.Insert(ctx, data)

	// Insert Detail
	dataDetail.IDPPCPProgramKami = data.IDPPCPProgramKami
	dataDetail.Tag = data.Tag
	if errDetail := repo.InsertDetail(ctx, dataDetail); errDetail != nil {
		log.Println("ERR Insert Detail: ", errDetail)
	}

	response, _ = GetProgramKami(ctx, repo, data.IDPPCPProgramKami)

	return
}
