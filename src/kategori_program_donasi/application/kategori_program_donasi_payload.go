package application

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"pemuda-peduli/src/kategori_program_donasi/domain/entity"

	"time"
)

type CreateKategoriProgramDonasi struct {
	Name string `json:"name" valid:"required"`
}

type UpdateKategoriProgramDonasi struct {
	Name string `json:"name" valid:"required"`
}

type KategoriProgramDonasiQuery struct {
	Limit  string                             `json:"limit" valid:"required"`
	Offset string                             `json:"offset" valid:"required"`
	Filter []KategoriProgramDonasiFilterQuery `json:"filters"`
	Order  string                             `json:"order"`
	Sort   string                             `json:"sort"`
}

type KategoriProgramDonasiFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadKategoriProgramDonasi struct {
	IDPPCPKategoriProgramDonasi string     `json:"id"`
	Name                        string     `json:"name"`
	CreatedAt                   time.Time  `json:"created_at"`
	CreatedBy                   *string    `json:"created_by"`
	UpdatedAt                   *time.Time `json:"updated_at"`
	UpdatedBy                   *string    `json:"updated_by"`
}

func GetCreatePayload(body []byte) (payload CreateKategoriProgramDonasi, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateKategoriProgramDonasi, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload KategoriProgramDonasiQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateKategoriProgramDonasi) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdateKategoriProgramDonasi) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r KategoriProgramDonasiQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateKategoriProgramDonasi) ToEntity() (data entity.KategoriProgramDonasiEntity) {
	data = entity.KategoriProgramDonasiEntity{
		Name:      r.Name,
		CreatedAt: time.Now().UTC(),
		IsDeleted: false,
	}
	return
}

func (r UpdateKategoriProgramDonasi) ToEntity() (data entity.KategoriProgramDonasiEntity) {
	data = entity.KategoriProgramDonasiEntity{
		Name:      r.Name,
		CreatedAt: time.Now().UTC(),
		IsDeleted: false,
	}
	return
}

func (r KategoriProgramDonasiQuery) ToEntity() (data entity.KategoriProgramDonasiQueryEntity) {
	filters := []entity.KategoriProgramDonasiFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.KategoriProgramDonasiFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.KategoriProgramDonasiQueryEntity{
		Limit:  r.Limit,
		Offset: r.Offset,
		Filter: filters,
		Order:  r.Order,
		Sort:   r.Sort,
	}
	return
}

func ToPayload(data entity.KategoriProgramDonasiEntity) (response ReadKategoriProgramDonasi) {
	response = ReadKategoriProgramDonasi{
		IDPPCPKategoriProgramDonasi: data.IDPPCPKategoriProgramDonasi,
		Name:                        data.Name,
		CreatedAt:                   data.CreatedAt,
		CreatedBy:                   data.CreatedBy,
		UpdatedAt:                   data.UpdatedAt,
		UpdatedBy:                   data.UpdatedBy,
	}
	return
}
