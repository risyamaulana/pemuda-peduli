package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/team/common/constants"
	"pemuda-peduli/src/team/domain/entity"
	"pemuda-peduli/src/team/domain/interfaces"
	"pemuda-peduli/src/team/infrastructure/repository"
	"time"

	flagDom "pemuda-peduli/src/team_flag/domain"
	flagRep "pemuda-peduli/src/team_flag/infrastructure/repository"
)

func UpdateTeam(ctx context.Context, db *db.ConnectTo, data entity.TeamEntity, id string) (response entity.TeamEntity, err error) {
	// Repo
	repo := repository.NewTeamRepository(db)
	flagRepo := flagRep.NewTeamFlagRepository(db)

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

	// Check data flag for flag id
	flagData, errFlag := flagDom.GetTeamFlag(ctx, &flagRepo, data.FlagID)
	if errFlag != nil {
		err = errors.New("Flag id is unauthorized / not found")
		return
	}

	checkData.Name = data.Name
	checkData.Role = data.Role
	checkData.ThumbnailPhotoURL = data.ThumbnailPhotoURL
	checkData.FacebookLink = data.FacebookLink
	checkData.GoogleLink = data.GoogleLink
	checkData.InstagramLink = data.InstagramLink
	checkData.LinkedinLink = data.LinkedinLink
	checkData.FlagID = flagData.ID
	checkData.FlagName = flagData.Name

	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = false

	response, err = repo.Update(ctx, checkData, id)
	return
}

func PublishTeam(ctx context.Context, repo interfaces.ITeamRepository, id string) (response entity.TeamEntity, err error) {
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

func HideTeam(ctx context.Context, repo interfaces.ITeamRepository, id string) (response entity.TeamEntity, err error) {
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

func DeleteTeam(ctx context.Context, repo interfaces.ITeamRepository, id string) (response entity.TeamEntity, err error) {
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
