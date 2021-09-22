package entity

import "time"

type TeamEntity struct {
	ID                int64      `db:"id"`
	IDPPCPTeam        string     `db:"id_pp_cp_team"`
	FlagID            int64      `db:"flag_id"`
	FlagName          string     `db:"flag_name"`
	Name              string     `db:"name"`
	Role              string     `db:"role"`
	ThumbnailPhotoURL string     `db:"thumbnail_photo_url"`
	FacebookLink      string     `db:"facebook_link"`
	GoogleLink        string     `db:"google_link"`
	InstagramLink     string     `db:"instagram_link"`
	LinkedinLink      string     `db:"linkedin_link"`
	Status            string     `db:"status"`
	CreatedAt         time.Time  `db:"created_at"`
	CreatedBy         *string    `db:"created_by"`
	UpdatedAt         *time.Time `db:"updated_at"`
	UpdatedBy         *string    `db:"updated_by"`
	PublishedAt       *time.Time `db:"published_at"`
	PublishedBy       *string    `db:"published_by"`
	IsDeleted         bool       `db:"is_deleted"`
}

type TeamQueryEntity struct {
	Limit         string `db:"limit"`
	Offset        string `db:"offset"`
	Filter        []TeamFilterQueryEntity
	Order         string `db:"order"`
	Sort          string `db:"sort"`
	CreatedAtFrom string `db:"created_at_from"`
	CreatedAtTo   string `db:"created_at_to"`
	PublishAtFrom string `db:"publish_at_from"`
	PublishAtTo   string `db:"publish_at_to"`
}

type TeamFilterQueryEntity struct {
	Field   string
	Keyword string
}
