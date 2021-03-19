package entity

import "time"

type PartnerKamiEntity struct {
	ID                int64      `db:"id"`
	IDPPCPPartnerKami string     `db:"id_pp_cp_partner_kami"`
	Name              string     `db:"name"`
	ThumbnailImageURL string     `db:"thumbnail_image_url"`
	Status            string     `db:"status"`
	CreatedAt         time.Time  `db:"created_at"`
	CreatedBy         *string    `db:"created_by"`
	UpdatedAt         *time.Time `db:"updated_at"`
	UpdatedBy         *string    `db:"updated_by"`
	PublishedAt       *time.Time `db:"published_at"`
	PublishedBy       *string    `db:"published_by"`
	IsDeleted         bool       `db:"is_deleted"`
}

type PartnerKamiQueryEntity struct {
	Limit         string `db:"limit"`
	Offset        string `db:"offset"`
	Filter        []PartnerKamiFilterQueryEntity
	Order         string `db:"order"`
	Sort          string `db:"sort"`
	CreatedAtFrom string `db:"created_at_from"`
	CreatedAtTo   string `db:"created_at_to"`
	PublishAtFrom string `db:"publish_at_from"`
	PublishAtTo   string `db:"publish_at_to"`
}

type PartnerKamiFilterQueryEntity struct {
	Field   string
	Keyword string
}
