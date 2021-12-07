package applications

import (
	"encoding/json"
	"pemuda-peduli/src/program_donasi_fundraiser/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateProgramDonasiFundraiser struct {
	Title       string   `json:"title" valid:"required"`
	SubTitle    string   `json:"sub_title"`
	SEOURL      string   `json:"seo_url"`
	Target      *float64 `json:"target" valid:"required"`
	Description string   `json:"description"`
}

type ProgramDonasiFundraiserQuery struct {
	Limit         string                               `json:"limit" valid:"required"`
	Offset        string                               `json:"offset" valid:"required"`
	Filter        []ProgramDonasiFundraiserFilterQuery `json:"filters"`
	Order         string                               `json:"order"`
	Sort          string                               `json:"sort"`
	CreatedAtFrom string                               `json:"created_at_from"`
	CreatedAtTo   string                               `json:"created_at_to"`
	PublishAtFrom string                               `json:"publish_at_from"`
	PublishAtTo   string                               `json:"publish_at_to"`
}

type ProgramDonasiFundraiserFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadProgramDonasiFundraiser struct {
	IDPPCPProgramDonasiFundraiser string `json:"id"`
	IDPPCPPenggalangDana          string `json:"id_pp_cp_penggalang_dana"`
	IDPPCPProgramDonasi           string `json:"id_pp_cp_program_donasi"`

	IDUser        string `json:"id_user"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	NamaLengkap   string `json:"nama_lengkap"`
	NamaPanggilan string `json:"nama_panggilan"`
	Alamat        string `json:"alamat"`

	Title      string `json:"title"`
	SubTitle   string `json:"sub_title"`
	DonasiType string `json:"donasi_type"`

	Tag               string `json:"tag"`
	Content           string `json:"content"`
	ThumbnailImageURL string `json:"thumbnail_image_url"`
	Description       string `json:"description"`

	Status    string     `json:"status"`
	ValidFrom *time.Time `json:"valid_from"`
	ValidTo   *time.Time `json:"valid_to"`

	Nominal         string   `json:"nominal"`
	Target          *float64 `json:"target"`
	DonationCollect float64  `json:"donation_collect"`

	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   *string    `json:"created_by"`
	UpdatedAt   *time.Time `json:"updated_at"`
	UpdatedBy   *string    `json:"updated_by"`
	PublishedAt *time.Time `json:"published_at"`
	PublishedBy *string    `json:"published_by"`

	IsDeleted bool `json:"is_deleted"`
	IsShow    bool `json:"is_show"`

	KitaBisaLink     *string `json:"kitabisa_link"`
	AyoBantuLink     *string `json:"ayobantu_link"`
	IDPPCPMasterQris *string `json:"id_pp_cp_master_qris"`
	QrisImageURL     *string `json:"qris_image_url"`

	SEOURL string `json:"seo_url"`
}

func GetCreatePayload(body []byte) (payload CreateProgramDonasiFundraiser, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload ProgramDonasiFundraiserQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateProgramDonasiFundraiser) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}

	return
}

func (r ProgramDonasiFundraiserQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateProgramDonasiFundraiser) ToEntity(idDonasi string) (data entity.ProgramDonasiFundraiserEntity) {
	data = entity.ProgramDonasiFundraiserEntity{
		IDPPCPProgramDonasi: idDonasi,
		Title:               r.Title,
		SubTitle:            r.SubTitle,
		Description:         r.Description,
		Target:              r.Target,
		IsDeleted:           false,
		IsShow:              true,
		SEOURL:              r.SEOURL,
	}
	return
}

func (r ProgramDonasiFundraiserQuery) ToEntity() (data entity.ProgramDonasiFundraiserQueryEntity) {
	filters := []entity.ProgramDonasiFundraiserFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.ProgramDonasiFundraiserFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.ProgramDonasiFundraiserQueryEntity{
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

func ToPayload(data entity.ProgramDonasiFundraiserEntity) (response ReadProgramDonasiFundraiser) {
	response = ReadProgramDonasiFundraiser{
		IDPPCPProgramDonasiFundraiser: data.IDPPCPProgramDonasiFundraiser,
		IDPPCPPenggalangDana:          data.IDPPCPPenggalangDana,
		IDPPCPProgramDonasi:           data.IDPPCPProgramDonasi,
		IDUser:                        data.IDUser,
		Username:                      data.Username,
		Email:                         data.Email,
		PhoneNumber:                   data.PhoneNumber,
		NamaLengkap:                   data.NamaLengkap,
		NamaPanggilan:                 data.NamaPanggilan,
		Alamat:                        data.Alamat,
		Title:                         data.Title,
		SubTitle:                      data.SubTitle,
		DonasiType:                    data.DonasiType,
		Tag:                           data.Tag,
		Content:                       data.Content,
		ThumbnailImageURL:             data.ThumbnailImageURL,
		Description:                   data.Description,
		Status:                        data.Status,
		ValidFrom:                     data.ValidFrom,
		ValidTo:                       data.ValidTo,
		Nominal:                       data.Nominal,
		Target:                        data.Target,
		DonationCollect:               data.DonationCollect,
		CreatedAt:                     data.CreatedAt,
		CreatedBy:                     data.CreatedBy,
		UpdatedAt:                     data.UpdatedAt,
		UpdatedBy:                     data.UpdatedBy,
		PublishedAt:                   data.PublishedAt,
		PublishedBy:                   data.PublishedBy,
		IsDeleted:                     data.IsDeleted,
		IsShow:                        data.IsShow,
		KitaBisaLink:                  data.KitaBisaLink,
		AyoBantuLink:                  data.AyoBantuLink,
		IDPPCPMasterQris:              data.IDPPCPMasterQris,
		QrisImageURL:                  data.QrisImageURL,
		SEOURL:                        data.SEOURL,
	}
	return
}
