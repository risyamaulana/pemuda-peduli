package entity

import "time"

type BannerEntity struct {
	ID                int64      `db:"id"`
	IDPPCPBanner      string     `db:"id_pp_cp_banner"`
	Tag               string     `db:"tag"`
	Title             string     `db:"title"`
	SubTitle          string     `db:"sub_title"`
	TitleContent      string     `db:"title_content"`
	ThumbnailImageURL string     `db:"thumbnail_image_url"`
	TitleButtonRight  *string    `db:"title_button_right"`
	DeeplinkRight     *string    `db:"deeplink_right"`
	TitleButtonLeft   *string    `db:"title_button_left"`
	DeeplinkLeft      *string    `db:"deeplink_left"`
	Description       *string    `db:"description"`
	Status            string     `db:"status"`
	CreatedAt         time.Time  `db:"created_at"`
	CreatedBy         *string    `db:"created_by"`
	UpdatedAt         *time.Time `db:"updated_at"`
	UpdatedBy         *string    `db:"updated_by"`
	PublishedAt       *time.Time `db:"published_at"`
	PublishedBy       *string    `db:"published_by"`
	IsDeleted         bool       `db:"is_deleted"`
}

type BannerQueryEntity struct {
	Limit         string `db:"limit"`
	Offset        string `db:"offset"`
	Filter        []BannerFilterQueryEntity
	Order         string `db:"order"`
	Sort          string `db:"sort"`
	CreatedAtFrom string `db:"created_at_from"`
	CreatedAtTo   string `db:"created_at_to"`
	PublishAtFrom string `db:"publish_at_from"`
	PublishAtTo   string `db:"publish_at_to"`
}

type BannerFilterQueryEntity struct {
	Field   string
	Keyword string
}
