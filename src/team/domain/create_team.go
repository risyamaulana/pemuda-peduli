package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/team/domain/entity"
	"pemuda-peduli/src/team/infrastructure/repository"

	flagDom "pemuda-peduli/src/team_flag/domain"
	flagRep "pemuda-peduli/src/team_flag/infrastructure/repository"
)

func CreateTeam(ctx context.Context, db *db.ConnectTo, data *entity.TeamEntity) (err error) {
	repo := repository.NewTeamRepository(db)
	flagRepo := flagRep.NewTeamFlagRepository(db)

	// Check data flag for flag id
	flagData, errFlag := flagDom.GetTeamFlag(ctx, &flagRepo, data.FlagID)
	if errFlag != nil {
		err = errors.New("Flag id is unauthorized / not found")
		return
	}
	data.FlagID = flagData.ID
	data.FlagName = flagData.Name

	err = repo.Insert(ctx, data)
	return
}
