package applications

import (
	"encoding/json"
	"pemuda-peduli/src/program_kami/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateProgramKami struct {
	Title             string `json:"title" valid:"required,alphanum"`
	SubTitle          string `json:"sub_title" valid:"required,alphanum"`
	Content           string `json:"content" valid:"required"`
	Tag               string `json:"tag"`
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"url"`
	Description       string `json:"description"`
}

type UpdateProgramKami struct {
	Title             string `json:"title" valid:"required,alphanum"`
	SubTitle          string `json:"sub_title" valid:"required,alphanum"`
	Content           string `json:"content" valid:"required"`
	Tag               string `json:"tag"`
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"url"`
	Description       string `json:"description"`
}

type ProgramKamiQuery struct {
	Limit         string                   `json:"limit" valid:"required"`
	Offset        string                   `json:"offset" valid:"required"`
	Filter        []ProgramKamiFilterQuery `json:"filters"`
	Order         string                   `json:"order"`
	Sort          string                   `json:"sort"`
	CreatedAtFrom string                   `json:"created_at_from"`
	CreatedAtTo   string                   `json:"created_at_to"`
	PublishAtFrom string                   `json:"publish_at_from"`
	PublishAtTo   string                   `json:"publish_at_to"`
}

type ProgramKamiFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadProgramKami struct {
	IDPPCPProgramKami string                 `json:"id"`
	Title             string                 `json:"title"`
	SubTitle          string                 `json:"sub_title"`
	Tag               string                 `json:"tag"`
	ThumbnailImageURL string                 `json:"thumbnail_image_url"`
	Detail            *ReadProgramKamiDetail `json:"detail,omitempty"`
	Description       string                 `json:"description"`
	Status            string                 `json:"status"`
	CreatedAt         time.Time              `json:"created_at"`
	CreatedBy         *string                `json:"created_by"`
	UpdatedAt         *time.Time             `json:"updated_at"`
	UpdatedBy         *string                `json:"updated_by"`
	PublishedAt       *time.Time             `json:"published_at"`
	PublishedBy       *string                `json:"published_by"`
	IsDeleted         bool                   `json:"is_deleted"`
}

type ReadProgramKamiDetail struct {
	Content string `json:"content"`
}

func GetCreatePayload(body []byte) (payload CreateProgramKami, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateProgramKami, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload ProgramKamiQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateProgramKami) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdateProgramKami) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r ProgramKamiQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateProgramKami) ToEntity() (data entity.ProgramKamiEntity, dataDetail entity.ProgramKamiDetailEntity) {
	data = entity.ProgramKamiEntity{
		Title:             r.Title,
		SubTitle:          r.SubTitle,
		Tag:               r.Tag,
		ThumbnailImageURL: r.ThumbnailImageURL,
		Description:       r.Description,
		CreatedAt:         time.Now(),
	}
	dataDetail = entity.ProgramKamiDetailEntity{
		Content: r.Content,
		Tag:     r.Tag,
	}

	return
}

func (r UpdateProgramKami) ToEntity() (data entity.ProgramKamiEntity, dataDetail entity.ProgramKamiDetailEntity) {
	data = entity.ProgramKamiEntity{
		Title:             r.Title,
		SubTitle:          r.SubTitle,
		Tag:               r.Tag,
		ThumbnailImageURL: r.ThumbnailImageURL,
		Description:       r.Description,
	}

	dataDetail = entity.ProgramKamiDetailEntity{
		Content: r.Content,
		Tag:     r.Tag,
	}
	return
}

func (r ProgramKamiQuery) ToEntity() (data entity.ProgramKamiQueryEntity) {
	filters := []entity.ProgramKamiFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.ProgramKamiFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.ProgramKamiQueryEntity{
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

func ToPayload(data entity.ProgramKamiEntity, isDetail bool) (response ReadProgramKami) {
	var detail *ReadProgramKamiDetail
	if isDetail {
		detail = &ReadProgramKamiDetail{
			Content: data.Detail.Content,
		}
	}

	response = ReadProgramKami{
		IDPPCPProgramKami: data.IDPPCPProgramKami,
		Title:             data.Title,
		SubTitle:          data.SubTitle,
		Tag:               data.Tag,
		ThumbnailImageURL: data.ThumbnailImageURL,
		Detail:            detail,
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
