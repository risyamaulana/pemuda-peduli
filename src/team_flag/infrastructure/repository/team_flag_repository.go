package repository

import (
	"context"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/team_flag/domain/entity"
	"pemuda-peduli/src/team_flag/domain/interfaces"
)

// TeamFlagRepository
type TeamFlagRepository struct {
	db *db.ConnectTo
	interfaces.ITeamFlagRepository
}

// NewTeamFlagRepository
func NewTeamFlagRepository(db *db.ConnectTo) TeamFlagRepository {
	return TeamFlagRepository{db: db}
}

func (c *TeamFlagRepository) Find(ctx context.Context) (response []entity.TeamFlagEntity, err error) {
	sql := `SELECT * FROM pp_cp_team_flag ORDER BY id ASC`

	// Result query
	if err = c.db.DBRead.Select(&response, sql); err != nil {
		return
	}
	return
}

func (c *TeamFlagRepository) Get(ctx context.Context, id int64) (response entity.TeamFlagEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_cp_team_flag WHERE id = $1", id); err != nil {
		return
	}
	return
}
