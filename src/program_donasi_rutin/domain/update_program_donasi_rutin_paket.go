package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/program_donasi_rutin/domain/entity"
	"pemuda-peduli/src/program_donasi_rutin/domain/interfaces"
	"pemuda-peduli/src/program_donasi_rutin/infrastructure/repository"
	"time"
)

func EditProgramDonasiRutinPaket(ctx context.Context, db *db.ConnectTo, data entity.ProgramDonasiRutinPaketEntity, id string) (resp entity.ProgramDonasiRutinPaketEntity, err error) {
	// Repo
	repo := repository.NewProgramDonasiRutinRepository(db)

	// Check available daata
	_, errCheckData := repo.GetPaket(ctx, id)
	if errCheckData != nil {
		err = errors.New("Data not found")
		return
	}

	// Check data Donasi Rutin
	_, errCheckData = GetProgramDonasiRutin(ctx, &repo, data.IDPPCPProgramDonasiRutin)
	if errCheckData != nil {
		err = errors.New("Failed, data donasi not found")
		return
	}

	resp, err = updateProgramDonasiRutinPaket(ctx, &repo, data, id)
	return
}

func DeleteProgramDonasiRutinPaket(ctx context.Context, db *db.ConnectTo, id string) (response entity.ProgramDonasiRutinPaketEntity, err error) {
	// Repo
	repo := repository.NewProgramDonasiRutinRepository(db)

	currentDate := time.Now().UTC()
	// Check available daata
	checkData, err := repo.GetPaket(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}

	checkData.IsDeleted = true
	checkData.UpdatedAt = &currentDate
	response, err = updateProgramDonasiRutinPaket(ctx, &repo, checkData, id)
	return
}

func updateProgramDonasiRutinPaket(ctx context.Context, repo interfaces.IProgramDonasiRutinRepository, data entity.ProgramDonasiRutinPaketEntity, id string) (resp entity.ProgramDonasiRutinPaketEntity, err error) {
	resp, err = repo.UpdatePaket(ctx, data, id)
	return
}
