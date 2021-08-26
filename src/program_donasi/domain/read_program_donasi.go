package domain

import (
	"context"
	"errors"
	"log"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/common/utility"
	"pemuda-peduli/src/program_donasi/domain/entity"
	"pemuda-peduli/src/program_donasi/infrastructure/repository"

	penggalangDanaDom "pemuda-peduli/src/penggalang_dana/domain"
)

func FindProgramDonasi(ctx context.Context, db *db.ConnectTo, data *entity.ProgramDonasiQueryEntity) (response []entity.ProgramDonasiEntity, count int, err error) {
	// Repo
	repo := repository.NewProgramDonasiRepository(db)

	response, count, err = repo.Find(ctx, data)
	for i, data := range response {
		// Get Detail
		dataDetail, _ := repo.GetDetail(ctx, data.IDPPCPProgramDonasi)

		response[i].Detail = dataDetail

		// Get Penggalang Dana
		dataPenggalangDana, _ := penggalangDanaDom.GetPenggalangDana(ctx, db, data.IDPPCPPenggalangDana)
		response[i].PenggalangDana = dataPenggalangDana
	}
	return
}

func GetProgramDonasi(ctx context.Context, db *db.ConnectTo, id string) (response entity.ProgramDonasiEntity, err error) {
	// Repo
	repo := repository.NewProgramDonasiRepository(db)

	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}

	// Get Detail
	dataDetail, errDetail := repo.GetDetail(ctx, response.IDPPCPProgramDonasi)
	if errDetail != nil {
		log.Println("ERR GET DETAIL : ", err)
	}

	response.Detail = dataDetail

	// Get Penggalang Dana
	dataPenggalangDana, _ := penggalangDanaDom.GetPenggalangDana(ctx, db, response.IDPPCPPenggalangDana)
	response.PenggalangDana = dataPenggalangDana

	log.Println(utility.PrettyPrint(response.Detail))
	return
}
