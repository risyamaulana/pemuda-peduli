package entity

import "time"

type HubungiKamiEntity struct {
	ID                int64      `db:"id"`
	IDPPCPHubungiKami string     `db:"id_pp_cp_hubungi_kami"`
	Icon              string     `db:"icon"`
	Link              string     `db:"link"`
	Title             string     `db:"title"`
	Status            string     `db:"status"`
	CreatedAt         time.Time  `db:"created_at"`
	CreatedBy         *string    `db:"created_by"`
	UpdatedAt         *time.Time `db:"updated_at"`
	UpdatedBy         *string    `db:"updated_by"`
	PublishedAt       *time.Time `db:"published_at"`
	PublishedBy       *string    `db:"published_by"`
	IsDeleted         bool       `db:"is_deleted"`
}

type HubungiKamiQueryEntity struct {
	Limit         string `db:"limit"`
	Offset        string `db:"offset"`
	Filter        []HubungiKamiFilterQueryEntity
	Order         string `db:"order"`
	Sort          string `db:"sort"`
	CreatedAtFrom string `db:"created_at_from"`
	CreatedAtTo   string `db:"created_at_to"`
	PublishAtFrom string `db:"publish_at_from"`
	PublishAtTo   string `db:"publish_at_to"`
}

type HubungiKamiFilterQueryEntity struct {
	Field   string
	Keyword string
}
