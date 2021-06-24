package applications

import (
	"encoding/json"
	"pemuda-peduli/src/program_donasi_kategori/domain/entity"

	"github.com/asaskevich/govalidator"
)

type CreateProgramDonasiKategori struct {
	KategoriName string `json:"kategori_name" valid:"required"`
}

type UpdateProgramDonasiKategori struct {
	KategoriName string `json:"kategori_name" valid:"required"`
}

type ProgramDonasiKategoriQuery struct {
	Limit         string                             `json:"limit" valid:"required"`
	Offset        string                             `json:"offset" valid:"required"`
	Filter        []ProgramDonasiKategoriFilterQuery `json:"filters"`
	Order         string                             `json:"order"`
	Sort          string                             `json:"sort"`
	CreatedAtFrom string                             `json:"created_at_from"`
	CreatedAtTo   string                             `json:"created_at_to"`
	PublishAtFrom string                             `json:"publish_at_from"`
	PublishAtTo   string                             `json:"publish_at_to"`
}

type ProgramDonasiKategoriFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadProgramDonasiKategori struct {
	IDPPCPProgramDonasiKategori string `json:"id"`
	KategoriName                string `json:"kategori_name"`
}

func GetCreatePayload(body []byte) (payload CreateProgramDonasiKategori, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateProgramDonasiKategori, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload ProgramDonasiKategoriQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateProgramDonasiKategori) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdateProgramDonasiKategori) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r ProgramDonasiKategoriQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateProgramDonasiKategori) ToEntity() (data entity.ProgramDonasiKategoriEntity) {
	data = entity.ProgramDonasiKategoriEntity{
		KategoriName: r.KategoriName,
	}
	return
}

func (r UpdateProgramDonasiKategori) ToEntity() (data entity.ProgramDonasiKategoriEntity) {
	data = entity.ProgramDonasiKategoriEntity{
		KategoriName: r.KategoriName,
	}
	return
}

func (r ProgramDonasiKategoriQuery) ToEntity() (data entity.ProgramDonasiKategoriQueryEntity) {
	filters := []entity.ProgramDonasiKategoriFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.ProgramDonasiKategoriFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.ProgramDonasiKategoriQueryEntity{
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

func ToPayload(data entity.ProgramDonasiKategoriEntity) (response ReadProgramDonasiKategori) {
	response = ReadProgramDonasiKategori{
		IDPPCPProgramDonasiKategori: data.IDPPCPProgramDonasiKategori,
		KategoriName:                data.KategoriName,
	}
	return
}
