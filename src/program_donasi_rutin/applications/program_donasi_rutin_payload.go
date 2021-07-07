package applications

import (
	"encoding/json"
	"pemuda-peduli/src/program_donasi_rutin/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateProgramDonasiRutin struct {
	IDPPCPProgramDonasiKategori string `json:"id_kategori" valid:"required"`

	Title             string `json:"title" valid:"required"`
	SubTitle          string `json:"sub_title" valid:"required"`
	Content           string `json:"content" valid:"required"`
	Benefit           string `json:"benefit"`
	Tag               string `json:"tag"`
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"url"`

	Description string `json:"description"`
}

type UpdateProgramDonasiRutin struct {
	IDPPCPProgramDonasiKategori string `json:"id_kategori" valid:"required"`

	Title             string `json:"title" valid:"required"`
	SubTitle          string `json:"sub_title" valid:"required"`
	Content           string `json:"content" valid:"required"`
	Benefit           string `json:"benefit"`
	Tag               string `json:"tag"`
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"url"`
	IsShow            *bool  `json:"is_show" valid:"required"`

	Description string `json:"description"`
}

type ProgramDonasiRutinQuery struct {
	Limit         string                          `json:"limit" valid:"required"`
	Offset        string                          `json:"offset" valid:"required"`
	Filter        []ProgramDonasiRutinFilterQuery `json:"filters"`
	Order         string                          `json:"order"`
	Sort          string                          `json:"sort"`
	CreatedAtFrom string                          `json:"created_at_from"`
	CreatedAtTo   string                          `json:"created_at_to"`
	PublishAtFrom string                          `json:"publish_at_from"`
	PublishAtTo   string                          `json:"publish_at_to"`
}

type ProgramDonasiRutinFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadProgramDonasiRutin struct {
	IDPPCPProgramDonasiRutin    string     `json:"id"`
	IDPPCPProgramDonasiKategori string     `json:"id_kategori"`
	KategoriName                string     `json:"kategori_name"`
	Title                       string     `json:"title"`
	SubTitle                    string     `json:"sub_title"`
	Content                     string     `json:"content"`
	Benefit                     string     `json:"benefit"`
	Tag                         string     `json:"tag"`
	ThumbnailImageURL           string     `json:"thumbnail_image_url"`
	ValidFrom                   *time.Time `json:"valid_from"`
	ValidTo                     *time.Time `json:"valid_to"`
	Target                      *float64   `json:"target"`
	Description                 string     `json:"description"`
	Status                      string     `json:"status"`
	CreatedAt                   time.Time  `json:"created_at"`
	CreatedBy                   *string    `json:"created_by"`
	UpdatedAt                   *time.Time `json:"updated_at"`
	UpdatedBy                   *string    `json:"updated_by"`
	PublishedAt                 *time.Time `json:"published_at"`
	PublishedBy                 *string    `json:"published_by"`
	IsDeleted                   bool       `json:"is_deleted"`
	IsShow                      bool       `json:"is_show"`
}

func GetCreatePayload(body []byte) (payload CreateProgramDonasiRutin, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateProgramDonasiRutin, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload ProgramDonasiRutinQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateProgramDonasiRutin) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdateProgramDonasiRutin) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r ProgramDonasiRutinQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateProgramDonasiRutin) ToEntity() (data entity.ProgramDonasiRutinEntity, dataDetail entity.ProgramDonasiRutinDetailEntity) {
	data = entity.ProgramDonasiRutinEntity{
		IDPPCPProgramDonasiKategori: r.IDPPCPProgramDonasiKategori,
		Title:                       r.Title,
		SubTitle:                    r.SubTitle,
		Tag:                         r.Tag,
		ThumbnailImageURL:           r.ThumbnailImageURL,
		Description:                 r.Description,
		CreatedAt:                   time.Now(),
	}

	dataDetail = entity.ProgramDonasiRutinDetailEntity{
		Content: r.Content,
		Benefit: r.Benefit,
		Tag:     r.Tag,
	}
	return
}

func (r UpdateProgramDonasiRutin) ToEntity() (data entity.ProgramDonasiRutinEntity, dataDetail entity.ProgramDonasiRutinDetailEntity) {
	data = entity.ProgramDonasiRutinEntity{
		IDPPCPProgramDonasiKategori: r.IDPPCPProgramDonasiKategori,
		Title:                       r.Title,
		SubTitle:                    r.SubTitle,
		Tag:                         r.Tag,
		ThumbnailImageURL:           r.ThumbnailImageURL,
		Description:                 r.Description,
		IsShow:                      *r.IsShow,
	}

	dataDetail = entity.ProgramDonasiRutinDetailEntity{
		Content: r.Content,
		Benefit: r.Benefit,
		Tag:     r.Tag,
	}
	return
}

func (r ProgramDonasiRutinQuery) ToEntity() (data entity.ProgramDonasiRutinQueryEntity) {
	filters := []entity.ProgramDonasiRutinFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.ProgramDonasiRutinFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.ProgramDonasiRutinQueryEntity{
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

func ToPayload(data entity.ProgramDonasiRutinEntity) (response ReadProgramDonasiRutin) {
	response = ReadProgramDonasiRutin{
		IDPPCPProgramDonasiRutin:    data.IDPPCPProgramDonasiRutin,
		IDPPCPProgramDonasiKategori: data.IDPPCPProgramDonasiKategori,
		KategoriName:                data.KategoriName,
		Title:                       data.Title,
		SubTitle:                    data.SubTitle,
		Content:                     data.Detail.Content,
		Benefit:                     data.Detail.Benefit,
		Tag:                         data.Tag,
		ThumbnailImageURL:           data.ThumbnailImageURL,
		Description:                 data.Description,
		Status:                      data.Status,
		CreatedAt:                   data.CreatedAt,
		CreatedBy:                   data.CreatedBy,
		UpdatedAt:                   data.UpdatedAt,
		UpdatedBy:                   data.UpdatedBy,
		PublishedAt:                 data.PublishedAt,
		PublishedBy:                 data.PublishedBy,
		IsDeleted:                   data.IsDeleted,
		IsShow:                      data.IsShow,
	}
	return
}
