package entity

import "time"

type TestimoniEntity struct {
	ID                int64      `db:"id"`
	IDPPCPTestimoni   string     `db:"id_pp_cp_testimoni"`
	Name              string     `db:"name"`
	Role              string     `db:"role"`
	ThumbnailPhotoURL string     `db:"thumbnail_photo_url"`
	Message           string     `db:"message"`
	Status            string     `db:"status"`
	CreatedAt         time.Time  `db:"created_at"`
	CreatedBy         *string    `db:"created_by"`
	UpdatedAt         *time.Time `db:"updated_at"`
	UpdatedBy         *string    `db:"updated_by"`
	PublishedAt       *time.Time `db:"published_at"`
	PublishedBy       *string    `db:"published_by"`
	IsDeleted         bool       `db:"is_deleted"`
}

type TestimoniQueryEntity struct {
	Limit         string `db:"limit"`
	Offset        string `db:"offset"`
	Filter        []TestimoniFilterQueryEntity
	Order         string `db:"order"`
	Sort          string `db:"sort"`
	CreatedAtFrom string `db:"created_at_from"`
	CreatedAtTo   string `db:"created_at_to"`
	PublishAtFrom string `db:"publish_at_from"`
	PublishAtTo   string `db:"publish_at_to"`
}

type TestimoniFilterQueryEntity struct {
	Field   string
	Keyword string
}
