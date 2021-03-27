package applications

import (
	"encoding/json"
	"pemuda-peduli/src/role/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateRole struct {
	RoleType  string `json:"role_type" valid:"required"`
	RoleLevel int    `json:"role_level" valid:"required"`
}

type UpdateRole struct {
	RoleType  string `json:"role_type" valid:"required"`
	RoleLevel int    `json:"role_level" valid:"required"`
}

type RoleQuery struct {
	Limit         string            `json:"limit" valid:"required"`
	Offset        string            `json:"offset" valid:"required"`
	Filter        []RoleFilterQuery `json:"filters"`
	Order         string            `json:"order"`
	Sort          string            `json:"sort"`
	CreatedAtFrom string            `json:"created_at_from"`
	CreatedAtTo   string            `json:"created_at_to"`
}

type RoleFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadRole struct {
	IDPPCPMasterRole string     `json:"id"`
	RoleType         string     `json:"role_type"`
	RoleLevel        int        `json:"role_level"`
	CreatedAt        time.Time  `json:"created_at"`
	CreatedBy        *string    `json:"created_by"`
	UpdatedAt        *time.Time `json:"updated_at"`
	UpdatedBy        *string    `json:"updated_by"`
}

func GetCreatePayload(body []byte) (payload CreateRole, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateRole, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload RoleQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateRole) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdateRole) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r RoleQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateRole) ToEntity() (data entity.RoleEntity) {
	data = entity.RoleEntity{
		RoleType:  r.RoleType,
		RoleLevel: r.RoleLevel,
	}
	return
}

func (r UpdateRole) ToEntity() (data entity.RoleEntity) {
	data = entity.RoleEntity{
		RoleType:  r.RoleType,
		RoleLevel: r.RoleLevel,
	}
	return
}

func (r RoleQuery) ToEntity() (data entity.RoleQueryEntity) {
	filters := []entity.RoleFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.RoleFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.RoleQueryEntity{
		Limit:         r.Limit,
		Offset:        r.Offset,
		Filter:        filters,
		Order:         r.Order,
		Sort:          r.Sort,
		CreatedAtFrom: r.CreatedAtFrom,
		CreatedAtTo:   r.CreatedAtTo,
	}
	return
}

func ToPayload(data entity.RoleEntity) (response ReadRole) {
	response = ReadRole{
		IDPPCPMasterRole: data.IDPPCPMasterRole,
		RoleType:         data.RoleType,
		RoleLevel:        data.RoleLevel,
		CreatedAt:        data.CreatedAt,
		CreatedBy:        data.CreatedBy,
		UpdatedAt:        data.UpdatedAt,
		UpdatedBy:        data.UpdatedBy,
	}
	return
}
