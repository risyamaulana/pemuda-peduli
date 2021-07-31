package applications

import (
	"encoding/json"
	"pemuda-peduli/src/album/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateAlbum struct {
	Title             string `json:"title" valid:"required"`
	SubTitle          string `json:"sub_title" valid:"required"`
	Tag               string `json:"tag"`
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"url"`
}

type UpdateAlbum struct {
	Title             string `json:"title" valid:"required"`
	SubTitle          string `json:"sub_title" valid:"required"`
	Tag               string `json:"tag"`
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"url"`
}

type AlbumQuery struct {
	Limit         string             `json:"limit" valid:"required"`
	Offset        string             `json:"offset" valid:"required"`
	Filter        []AlbumFilterQuery `json:"filters"`
	Order         string             `json:"order"`
	Sort          string             `json:"sort"`
	CreatedAtFrom string             `json:"created_at_from"`
	CreatedAtTo   string             `json:"created_at_to"`
	PublishAtFrom string             `json:"publish_at_from"`
	PublishAtTo   string             `json:"publish_at_to"`
}

type AlbumFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadAlbum struct {
	IDPPCPAlbum       string     `json:"id"`
	Title             string     `json:"title"`
	SubTitle          string     `json:"sub_title"`
	Tag               string     `json:"tag"`
	ThumbnailImageURL string     `json:"thumbnail_image_url"`
	Status            string     `json:"status"`
	CreatedAt         time.Time  `json:"created_at"`
	CreatedBy         *string    `json:"created_by"`
	UpdatedAt         *time.Time `json:"updated_at"`
	UpdatedBy         *string    `json:"updated_by"`
	PublishedAt       *time.Time `json:"published_at"`
	PublishedBy       *string    `json:"published_by"`
	IsDeleted         bool       `json:"is_deleted"`
}

func GetCreatePayload(body []byte) (payload CreateAlbum, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateAlbum, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload AlbumQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateAlbum) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdateAlbum) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r AlbumQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateAlbum) ToEntity() (data entity.AlbumEntity) {
	data = entity.AlbumEntity{
		Title:             r.Title,
		SubTitle:          r.SubTitle,
		Tag:               r.Tag,
		ThumbnailImageURL: r.ThumbnailImageURL,
		CreatedAt:         time.Now().UTC(),
	}
	return
}

func (r UpdateAlbum) ToEntity() (data entity.AlbumEntity) {
	data = entity.AlbumEntity{
		Title:             r.Title,
		SubTitle:          r.SubTitle,
		Tag:               r.Tag,
		ThumbnailImageURL: r.ThumbnailImageURL,
	}
	return
}

func (r AlbumQuery) ToEntity() (data entity.AlbumQueryEntity) {
	filters := []entity.AlbumFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.AlbumFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.AlbumQueryEntity{
		Limit:         r.Limit,
		Offset:        r.Offset,
		Filter:        filters,
		Order:         r.Order,
		Sort:          r.Sort,
		CreatedAtFrom: r.CreatedAtFrom,
		CreatedAtTo:   r.CreatedAtTo,
		PublishAtFrom: r.PublishAtFrom,
		PublishAtTo:   r.PublishAtTo,
	}
	return
}

func ToPayload(data entity.AlbumEntity) (response ReadAlbum) {
	response = ReadAlbum{
		IDPPCPAlbum:       data.IDPPCPAlbum,
		Title:             data.Title,
		SubTitle:          data.SubTitle,
		Tag:               data.Tag,
		ThumbnailImageURL: data.ThumbnailImageURL,
		Status:            data.Status,
		CreatedAt:         data.CreatedAt,
		CreatedBy:         data.CreatedBy,
		UpdatedAt:         data.UpdatedAt,
		UpdatedBy:         data.UpdatedBy,
		PublishedAt:       data.PublishedAt,
		PublishedBy:       data.PublishedBy,
		IsDeleted:         data.IsDeleted,
	}
	return
}
