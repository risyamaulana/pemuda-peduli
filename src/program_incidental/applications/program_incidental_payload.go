package applications

import (
	"encoding/json"
	"pemuda-peduli/src/program_incidental/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateProgramIncidental struct {
	Title             string `json:"title" valid:"required"`
	SubTitle          string `json:"sub_title" valid:"required"`
	Content           string `json:"content" valid:"required"`
	Tag               string `json:"tag"`
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"url"`
	Description       string `json:"description"`
}

type UpdateProgramIncidental struct {
	Title             string `json:"title" valid:"required"`
	SubTitle          string `json:"sub_title" valid:"required"`
	Content           string `json:"content" valid:"required"`
	Tag               string `json:"tag"`
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"url"`
	Description       string `json:"description"`
}

type ProgramIncidentalQuery struct {
	Limit         string                         `json:"limit" valid:"required"`
	Offset        string                         `json:"offset" valid:"required"`
	Filter        []ProgramIncidentalFilterQuery `json:"filters"`
	Order         string                         `json:"order"`
	Sort          string                         `json:"sort"`
	CreatedAtFrom string                         `json:"created_at_from"`
	CreatedAtTo   string                         `json:"created_at_to"`
	PublishAtFrom string                         `json:"publish_at_from"`
	PublishAtTo   string                         `json:"publish_at_to"`
}

type ProgramIncidentalFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadProgramIncidental struct {
	IDPPCPProgramIncidental string     `json:"id"`
	Title                   string     `json:"title"`
	SubTitle                string     `json:"sub_title"`
	Tag                     string     `json:"tag"`
	ThumbnailImageURL       string     `json:"thumbnail_image_url"`
	Content                 string     `json:"content,omitempty"`
	Description             string     `json:"description"`
	Status                  string     `json:"status"`
	CreatedAt               time.Time  `json:"created_at"`
	CreatedBy               *string    `json:"created_by"`
	UpdatedAt               *time.Time `json:"updated_at"`
	UpdatedBy               *string    `json:"updated_by"`
	PublishedAt             *time.Time `json:"published_at"`
	PublishedBy             *string    `json:"published_by"`
	IsDeleted               bool       `json:"is_deleted"`
}

func GetCreatePayload(body []byte) (payload CreateProgramIncidental, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateProgramIncidental, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload ProgramIncidentalQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateProgramIncidental) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}

	return
}

func (r UpdateProgramIncidental) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}

	return
}

func (r ProgramIncidentalQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateProgramIncidental) ToEntity() (data entity.ProgramIncidentalEntity) {
	data = entity.ProgramIncidentalEntity{
		Title:             r.Title,
		SubTitle:          r.SubTitle,
		Tag:               r.Tag,
		ThumbnailImageURL: r.ThumbnailImageURL,
		Content:           r.Content,
		Description:       r.Description,
		CreatedAt:         time.Now().UTC(),
	}

	return
}

func (r UpdateProgramIncidental) ToEntity() (data entity.ProgramIncidentalEntity) {

	data = entity.ProgramIncidentalEntity{

		Title:             r.Title,
		SubTitle:          r.SubTitle,
		Tag:               r.Tag,
		Content:           r.Content,
		ThumbnailImageURL: r.ThumbnailImageURL,
		Description:       r.Description,
	}

	return
}

func (r ProgramIncidentalQuery) ToEntity() (data entity.ProgramIncidentalQueryEntity) {
	filters := []entity.ProgramIncidentalFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.ProgramIncidentalFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.ProgramIncidentalQueryEntity{
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

func ToPayload(data entity.ProgramIncidentalEntity, isDetail bool) (response ReadProgramIncidental) {

	response = ReadProgramIncidental{
		IDPPCPProgramIncidental: data.IDPPCPProgramIncidental,
		Title:                   data.Title,
		SubTitle:                data.SubTitle,
		Tag:                     data.Tag,
		ThumbnailImageURL:       data.ThumbnailImageURL,
		Content:                 data.Content,
		Description:             data.Description,
		Status:                  data.Status,
		CreatedAt:               data.CreatedAt,
		CreatedBy:               data.CreatedBy,
		UpdatedAt:               data.UpdatedAt,
		UpdatedBy:               data.UpdatedBy,
		PublishedAt:             data.PublishedAt,
		PublishedBy:             data.PublishedBy,
		IsDeleted:               data.IsDeleted,
	}
	return
}
