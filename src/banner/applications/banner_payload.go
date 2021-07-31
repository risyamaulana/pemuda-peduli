package applications

import (
	"encoding/json"
	"pemuda-peduli/src/banner/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateBanner struct {
	Title             string  `json:"title" valid:"required"`
	SubTitle          string  `json:"sub_title" valid:"required"`
	TitleContent      string  `json:"title_content"`
	ThumbnailImageURL string  `json:"thumbnail_image_url" valid:"url"`
	TitleButtonRight  *string `json:"title_button_right"`
	DeeplinkRight     *string `json:"deeplink_right" valid:"url"`
	TitleButtonLeft   *string `json:"title_button_left"`
	DeeplinkLeft      *string `json:"deeplink_left" valid:"url"`
	Description       *string `json:"description"`
}

type UpdateBanner struct {
	Title             string  `json:"title" valid:"required"`
	SubTitle          string  `json:"sub_title" valid:"required"`
	TitleContent      string  `json:"title_content"`
	ThumbnailImageURL string  `json:"thumbnail_image_url" valid:"url"`
	TitleButtonRight  *string `json:"title_button_right"`
	DeeplinkRight     *string `json:"deeplink_right" valid:"url"`
	TitleButtonLeft   *string `json:"title_button_left"`
	DeeplinkLeft      *string `json:"deeplink_left" valid:"url"`
	Description       *string `json:"description"`
}

type BannerQuery struct {
	Limit         string              `json:"limit" valid:"required"`
	Offset        string              `json:"offset" valid:"required"`
	Filter        []BannerFilterQuery `json:"filters"`
	Order         string              `json:"order"`
	Sort          string              `json:"sort"`
	CreatedAtFrom string              `json:"created_at_from"`
	CreatedAtTo   string              `json:"created_at_to"`
	PublishAtFrom string              `json:"publish_at_from"`
	PublishAtTo   string              `json:"publish_at_to"`
}

type BannerFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadBanner struct {
	IDPPCPBanner      string     `json:"id"`
	Title             string     `json:"title"`
	SubTitle          string     `json:"sub_title"`
	TitleContent      string     `json:"title_content"`
	ThumbnailImageURL string     `json:"thumbnail_image_url"`
	TitleButtonRight  *string    `json:"title_button_right"`
	DeeplinkRight     *string    `json:"deeplink_right"`
	TitleButtonLeft   *string    `json:"title_button_left"`
	DeeplinkLeft      *string    `json:"deeplink_left"`
	Description       *string    `json:"description"`
	Status            string     `json:"status"`
	CreatedAt         time.Time  `json:"created_at"`
	CreatedBy         *string    `json:"created_by"`
	UpdatedAt         *time.Time `json:"updated_at"`
	UpdatedBy         *string    `json:"updated_by"`
	PublishedAt       *time.Time `json:"published_at"`
	PublishedBy       *string    `json:"published_by"`
	IsDeleted         bool       `json:"is_deleted"`
}

func GetCreatePayload(body []byte) (payload CreateBanner, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateBanner, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload BannerQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateBanner) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdateBanner) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r BannerQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateBanner) ToEntity() (data entity.BannerEntity) {
	data = entity.BannerEntity{
		Title:             r.Title,
		SubTitle:          r.SubTitle,
		TitleContent:      r.TitleContent,
		ThumbnailImageURL: r.ThumbnailImageURL,
		TitleButtonRight:  r.TitleButtonRight,
		DeeplinkRight:     r.DeeplinkRight,
		TitleButtonLeft:   r.TitleButtonLeft,
		DeeplinkLeft:      r.DeeplinkLeft,
		Description:       r.Description,
		CreatedAt:         time.Now().UTC(),
	}
	return
}

func (r UpdateBanner) ToEntity() (data entity.BannerEntity) {
	data = entity.BannerEntity{
		Title:             r.Title,
		SubTitle:          r.SubTitle,
		TitleContent:      r.TitleContent,
		ThumbnailImageURL: r.ThumbnailImageURL,
		TitleButtonRight:  r.TitleButtonRight,
		DeeplinkRight:     r.DeeplinkRight,
		TitleButtonLeft:   r.TitleButtonLeft,
		DeeplinkLeft:      r.DeeplinkLeft,
		Description:       r.Description,
	}
	return
}

func (r BannerQuery) ToEntity() (data entity.BannerQueryEntity) {
	filters := []entity.BannerFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.BannerFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.BannerQueryEntity{
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

func ToPayload(data entity.BannerEntity) (response ReadBanner) {
	response = ReadBanner{
		IDPPCPBanner:      data.IDPPCPBanner,
		Title:             data.Title,
		SubTitle:          data.SubTitle,
		TitleContent:      data.TitleContent,
		ThumbnailImageURL: data.ThumbnailImageURL,
		TitleButtonRight:  data.TitleButtonRight,
		DeeplinkRight:     data.DeeplinkRight,
		TitleButtonLeft:   data.TitleButtonLeft,
		DeeplinkLeft:      data.DeeplinkLeft,
		Description:       data.Description,
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
