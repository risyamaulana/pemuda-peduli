package entity

import "time"

type KontakKamiEntity struct {
	ID               int64      `db:"id"`
	IDPPCPKontakKami string     `db:"id_pp_cp_kontak_kami"`
	SKLegalitas      string     `db:"sk_legalitas"`
	Address          string     `db:"address"`
	ContactName      string     `db:"contact_name"`
	Icon             string     `db:"icon"`
	ContactLink      string     `db:"contact_link"`
	Menu             string     `db:"menu"`
	Status           string     `db:"status"`
	CreatedAt        time.Time  `db:"created_at"`
	CreatedBy        *string    `db:"created_by"`
	UpdatedAt        *time.Time `db:"updated_at"`
	UpdatedBy        *string    `db:"updated_by"`
	PublishedAt      *time.Time `db:"published_at"`
	PublishedBy      *string    `db:"published_by"`
	IsDeleted        bool       `db:"is_deleted"`
}

type KontakKamiQueryEntity struct {
	Limit         string `db:"limit"`
	Offset        string `db:"offset"`
	Filter        []KontakKamiFilterQueryEntity
	Order         string `db:"order"`
	Sort          string `db:"sort"`
	CreatedAtFrom string `db:"created_at_from"`
	CreatedAtTo   string `db:"created_at_to"`
	PublishAtFrom string `db:"publish_at_from"`
	PublishAtTo   string `db:"publish_at_to"`
}

type KontakKamiFilterQueryEntity struct {
	Field   string
	Keyword string
}
