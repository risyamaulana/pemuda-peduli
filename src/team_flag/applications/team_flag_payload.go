package applications

import (
	"pemuda-peduli/src/team_flag/domain/entity"
	"time"
)

type ReadTeamFlag struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	IsDeleted bool       `json:"is_deleted"`
}

func ToPayload(data entity.TeamFlagEntity) (response ReadTeamFlag) {
	response = ReadTeamFlag{
		ID:        data.ID,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		IsDeleted: data.IsDeleted,
	}
	return
}
