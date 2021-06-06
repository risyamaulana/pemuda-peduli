package applications

import (
	"encoding/json"
	"pemuda-peduli/src/team/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateTeam struct {
	Name              string `json:"name" valid:"required"`
	Role              string `json:"role" valid:"required"`
	ThumbnailPhotoURL string `json:"thumbnail_photo_url" valid:"required,url"`
	FacebookLink      string `json:"facebook_link" valid:"url"`
	GoogleLink        string `json:"google_link" valid:"url"`
	InstagramLink     string `json:"instagram_link" valid:"url"`
	LinkedinLink      string `json:"linkedin_link" valid:"url"`
}

type UpdateTeam struct {
	Name              string `json:"name" valid:"required"`
	Role              string `json:"role" valid:"required"`
	ThumbnailPhotoURL string `json:"thumbnail_photo_url" valid:"required,url"`
	FacebookLink      string `json:"facebook_link" valid:"url"`
	GoogleLink        string `json:"google_link" valid:"url"`
	InstagramLink     string `json:"instagram_link" valid:"url"`
	LinkedinLink      string `json:"linkedin_link" valid:"url"`
}

type TeamQuery struct {
	Limit         string            `json:"limit" valid:"required"`
	Offset        string            `json:"offset" valid:"required"`
	Filter        []TeamFilterQuery `json:"filters"`
	Order         string            `json:"order"`
	Sort          string            `json:"sort"`
	CreatedAtFrom string            `json:"created_at_from"`
	CreatedAtTo   string            `json:"created_at_to"`
	PublishAtFrom string            `json:"publish_at_from"`
	PublishAtTo   string            `json:"publish_at_to"`
}

type TeamFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadTeam struct {
	IDPPCPTeam        string     `json:"id"`
	Name              string     `json:"name"`
	Role              string     `json:"role"`
	ThumbnailPhotoURL string     `json:"thumbnail_photo_url"`
	FacebookLink      string     `json:"facebook_link"`
	GoogleLink        string     `json:"google_link"`
	InstagramLink     string     `json:"instagram_link"`
	LinkedinLink      string     `json:"linkedin_link"`
	Status            string     `json:"status"`
	CreatedAt         time.Time  `json:"created_at"`
	CreatedBy         *string    `json:"created_by"`
	UpdatedAt         *time.Time `json:"updated_at"`
	UpdatedBy         *string    `json:"updated_by"`
	PublishedAt       *time.Time `json:"published_at"`
	PublishedBy       *string    `json:"published_by"`
	IsDeleted         bool       `json:"is_deleted"`
}

func GetCreatePayload(body []byte) (payload CreateTeam, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateTeam, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload TeamQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateTeam) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdateTeam) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r TeamQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateTeam) ToEntity() (data entity.TeamEntity) {
	data = entity.TeamEntity{
		Name:              r.Name,
		Role:              r.Role,
		ThumbnailPhotoURL: r.ThumbnailPhotoURL,
		FacebookLink:      r.FacebookLink,
		GoogleLink:        r.GoogleLink,
		InstagramLink:     r.InstagramLink,
		LinkedinLink:      r.LinkedinLink,
		CreatedAt:         time.Now(),
	}
	return
}

func (r UpdateTeam) ToEntity() (data entity.TeamEntity) {
	data = entity.TeamEntity{
		Name:              r.Name,
		Role:              r.Role,
		ThumbnailPhotoURL: r.ThumbnailPhotoURL,
		FacebookLink:      r.FacebookLink,
		GoogleLink:        r.GoogleLink,
		InstagramLink:     r.InstagramLink,
		LinkedinLink:      r.LinkedinLink,
	}
	return
}

func (r TeamQuery) ToEntity() (data entity.TeamQueryEntity) {
	filters := []entity.TeamFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.TeamFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.TeamQueryEntity{
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

func ToPayload(data entity.TeamEntity) (response ReadTeam) {
	response = ReadTeam{
		IDPPCPTeam:        data.IDPPCPTeam,
		Name:              data.Name,
		Role:              data.Role,
		ThumbnailPhotoURL: data.ThumbnailPhotoURL,
		FacebookLink:      data.FacebookLink,
		GoogleLink:        data.GoogleLink,
		InstagramLink:     data.InstagramLink,
		LinkedinLink:      data.LinkedinLink,
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
