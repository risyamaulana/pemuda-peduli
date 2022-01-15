package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/program_donasi_fundraiser/domain/entity"
	"pemuda-peduli/src/program_donasi_fundraiser/infrastructure/repository"
	"time"
)

func UpdateDonationCollectFundraiser(ctx context.Context, db *db.ConnectTo, id string, amount float64) (response entity.ProgramDonasiFundraiserEntity, err error) {
	// Repo
	repo := repository.NewProgramDonasiFundraiserRepository(db)

	currentDate := time.Now().UTC()
	// Check available data
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("failed update data collect for fundraiser")
		return
	}

	response.DonationCollect = response.DonationCollect + amount
	response.UpdatedAt = &currentDate

	err = repo.Update(ctx, &response)
	return
}
