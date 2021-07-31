package applications

import (
	"encoding/json"
	"pemuda-peduli/src/testimoni/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateTestimoni struct {
	Name              string `json:"name" valid:"required"`
	Role              string `json:"role" valid:"required"`
	ThumbnailPhotoURL string `json:"thumbnail_photo_url" valid:"required,url"`
	Message           string `json:"message"`
}

type UpdateTestimoni struct {
	Name              string `json:"name" valid:"required"`
	Role              string `json:"role" valid:"required"`
	ThumbnailPhotoURL string `json:"thumbnail_photo_url" valid:"required,url"`
	Message           string `json:"message"`
}

type TestimoniQuery struct {
	Limit         string                 `json:"limit" valid:"required"`
	Offset        string                 `json:"offset" valid:"required"`
	Filter        []TestimoniFilterQuery `json:"filters"`
	Order         string                 `json:"order"`
	Sort          string                 `json:"sort"`
	CreatedAtFrom string                 `json:"created_at_from"`
	CreatedAtTo   string                 `json:"created_at_to"`
	PublishAtFrom string                 `json:"publish_at_from"`
	PublishAtTo   string                 `json:"publish_at_to"`
}

type TestimoniFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadTestimoni struct {
	IDPPCPTestimoni   string     `json:"id"`
	Name              string     `json:"name"`
	Role              string     `json:"role"`
	ThumbnailPhotoURL string     `json:"thumbnail_photo_url"`
	Message           string     `json:"message"`
	Status            string     `json:"status"`
	CreatedAt         time.Time  `json:"created_at"`
	CreatedBy         *string    `json:"created_by"`
	UpdatedAt         *time.Time `json:"updated_at"`
	UpdatedBy         *string    `json:"updated_by"`
	PublishedAt       *time.Time `json:"published_at"`
	PublishedBy       *string    `json:"published_by"`
	IsDeleted         bool       `json:"is_deleted"`
}

func GetCreatePayload(body []byte) (payload CreateTestimoni, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateTestimoni, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload TestimoniQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateTestimoni) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdateTestimoni) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r TestimoniQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateTestimoni) ToEntity() (data entity.TestimoniEntity) {
	data = entity.TestimoniEntity{
		Name:              r.Name,
		Role:              r.Role,
		ThumbnailPhotoURL: r.ThumbnailPhotoURL,
		Message:           r.Message,
		CreatedAt:         time.Now().UTC(),
	}
	return
}

func (r UpdateTestimoni) ToEntity() (data entity.TestimoniEntity) {
	data = entity.TestimoniEntity{
		Name:              r.Name,
		Role:              r.Role,
		ThumbnailPhotoURL: r.ThumbnailPhotoURL,
		Message:           r.Message,
	}
	return
}

func (r TestimoniQuery) ToEntity() (data entity.TestimoniQueryEntity) {
	filters := []entity.TestimoniFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.TestimoniFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.TestimoniQueryEntity{
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

func ToPayload(data entity.TestimoniEntity) (response ReadTestimoni) {
	response = ReadTestimoni{
		IDPPCPTestimoni:   data.IDPPCPTestimoni,
		Name:              data.Name,
		Role:              data.Role,
		ThumbnailPhotoURL: data.ThumbnailPhotoURL,
		Message:           data.Message,
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
