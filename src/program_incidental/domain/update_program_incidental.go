package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/program_incidental/common/constants"
	"pemuda-peduli/src/program_incidental/domain/entity"
	"pemuda-peduli/src/program_incidental/domain/interfaces"
	"time"
)

func UpdateProgramIncidental(ctx context.Context, repo interfaces.IProgramIncidentalRepository, data entity.ProgramIncidentalEntity, id string) (response entity.ProgramIncidentalEntity, err error) {
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

	data.ID = checkData.ID
	data.IDPPCPProgramIncidental = checkData.IDPPCPProgramIncidental
	data.Status = checkData.Status
	data.CreatedAt = checkData.CreatedAt
	data.CreatedBy = checkData.CreatedBy

	data.UpdatedAt = &currentDate

	data.PublishedAt = checkData.PublishedAt
	data.PublishedBy = checkData.PublishedBy

	data.IsDeleted = checkData.IsDeleted

	response, err = repo.Update(ctx, checkData, id)
	return
}

func PublishProgramIncidental(ctx context.Context, repo interfaces.IProgramIncidentalRepository, id string) (response entity.ProgramIncidentalEntity, err error) {
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
	checkData.Status = constants.StatusPublished
	checkData.UpdatedAt = &currentDate
	checkData.PublishedAt = &currentDate
	checkData.IsDeleted = false
	response, err = repo.Update(ctx, checkData, id)
	return
}

func HideProgramIncidental(ctx context.Context, repo interfaces.IProgramIncidentalRepository, id string) (response entity.ProgramIncidentalEntity, err error) {
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

	checkData.Status = constants.StatusHidden
	checkData.UpdatedAt = &currentDate
	checkData.PublishedAt = nil
	checkData.IsDeleted = false
	response, err = repo.Update(ctx, checkData, id)
	return
}

func DeleteProgramIncidental(ctx context.Context, repo interfaces.IProgramIncidentalRepository, id string) (response entity.ProgramIncidentalEntity, err error) {
	currentDate := time.Now().UTC()
	// Check available daata
	checkData, err := repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}

	checkData.Status = constants.StatusDeleted
	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = true
	response, err = repo.Update(ctx, checkData, id)
	return
}
