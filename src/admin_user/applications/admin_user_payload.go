package applications

import (
	"encoding/json"
	"errors"
	"pemuda-peduli/src/admin_user/domain/entity"
	"pemuda-peduli/src/common/utility"

	roleApp "pemuda-peduli/src/role/applications"

	"time"

	"github.com/asaskevich/govalidator"
)

type CreateAdminUser struct {
	Username        string `json:"username" valid:"required,minstringlength(5)"`
	Password        string `json:"password" valid:"required,minstringlength(8),alphanum"`
	ConfirmPassword string `json:"confirm_password" valid:"required,minstringlength(8),alphanum"`
	Email           string `json:"email" valid:"required,email"`
	NamaLengkap     string `json:"nama_lengkap" valid:"required"`
	Alamat          string `json:"alamat"`
	Role            string `json:"role" valid:"required"`
}

type UpdateAdminUser struct {
	Username    string `json:"username" valid:"required,minstringlength(5)"`
	Email       string `json:"email" valid:"required"`
	NamaLengkap string `json:"nama_lengkap" valid:"required"`
	Alamat      string `json:"alamat"`
}

type UsernameAdmin struct {
	Username string `json:"username" valid:"required,minstringlength(5)"`
}

// TODO: Change password
type ChangePassword struct {
	OldPassword        string `json:"old_password" valid:"required,minstringlength(8),alphanum"`
	NewPassword        string `json:"new_password" valid:"required,minstringlength(8),alphanum"`
	ConfirmNewPassword string `json:"confirm_new_password" valid:"required,minstringlength(8),alphanum"`
}

// TODO: Reset password
type ResetPassword struct {
	ID string `json:"id" valid:"required"`
}

// TODO: Change role
type ChangeRole struct {
	ID   string `json:"id" valid:"required"`
	Role string `json:"role" valid:"required"`
}

type AdminUserQuery struct {
	Limit         string                 `json:"limit" valid:"required"`
	Offset        string                 `json:"offset" valid:"required"`
	Filter        []AdminUserFilterQuery `json:"filters"`
	Order         string                 `json:"order"`
	Sort          string                 `json:"sort"`
	CreatedAtFrom string                 `json:"created_at_from"`
	CreatedAtTo   string                 `json:"created_at_to"`
	PublishAtFrom string                 `json:"publish_at_from"`
	PublishAtTo   string                 `json:"publish_at_to"`
}

type AdminUserFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadAdminUser struct {
	IDPPCPAdminUser string           `json:"id"`
	Username        string           `json:"username"`
	Email           string           `json:"email"`
	NamaLengkap     string           `json:"nama_lengkap"`
	Alamat          string           `json:"alamat"`
	Role            string           `json:"-"`
	RoleData        roleApp.ReadRole `json:"role"`
	CreatedAt       time.Time        `json:"created_at"`
	CreatedBy       *string          `json:"created_by"`
	UpdatedAt       *time.Time       `json:"updated_at"`
	UpdatedBy       *string          `json:"updated_by"`
	IsDeleted       bool             `json:"is_deleted"`
}

type ResetPasswordResponse struct {
	NewPassword string `json:"new_password"`
}

func GetCreatePayload(body []byte) (payload CreateAdminUser, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateAdminUser, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUsernameAdminPayload(body []byte) (payload UsernameAdmin, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetChangePasswordPayload(body []byte) (payload ChangePassword, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetResetPasswordPayload(body []byte) (payload ResetPassword, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetChangeRolePayload(body []byte) (payload ChangeRole, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload AdminUserQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateAdminUser) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}

	if r.Password != r.ConfirmPassword {
		err = errors.New("Password not match")
		return
	}

	return
}

func (r UpdateAdminUser) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UsernameAdmin) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r ChangePassword) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}

	// get user reference
	if r.NewPassword != r.ConfirmNewPassword {
		err = errors.New("Password not match")
	}
	return
}

func (r ResetPassword) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}

	return
}

func (r ChangeRole) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}

	return
}

func (r AdminUserQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateAdminUser) ToEntity() (data entity.AdminUserEntity) {
	// Create Salt
	salt := utility.GenerateSalt(4)
	password := utility.GeneratePass(salt, r.Password)
	data = entity.AdminUserEntity{
		Username:    r.Username,
		Salt:        salt,
		Password:    password,
		Email:       r.Email,
		NamaLengkap: r.NamaLengkap,
		Alamat:      r.Alamat,
		Role:        r.Role,
		CreatedAt:   time.Now(),
	}
	return
}

func (r UpdateAdminUser) ToEntity() (data entity.AdminUserEntity) {
	data = entity.AdminUserEntity{
		Username:    r.Username,
		Email:       r.Email,
		NamaLengkap: r.NamaLengkap,
		Alamat:      r.Alamat,
	}
	return
}

func (r ChangePassword) ToEntity() (data entity.AdminUserEntity) {
	salt := utility.GenerateSalt(4)
	password := utility.GeneratePass(salt, r.NewPassword)
	data = entity.AdminUserEntity{
		Salt:     salt,
		Password: password,
	}
	return
}

func (r AdminUserQuery) ToEntity() (data entity.AdminUserQueryEntity) {
	filters := []entity.AdminUserFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.AdminUserFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.AdminUserQueryEntity{
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

func ToPayload(data entity.AdminUserEntity) (response ReadAdminUser) {
	response = ReadAdminUser{
		IDPPCPAdminUser: data.IDPPCPAdminUser,
		Username:        data.Username,
		Email:           data.Email,
		NamaLengkap:     data.NamaLengkap,
		Alamat:          data.Alamat,
		Role:            data.Role,
		RoleData:        roleApp.ToPayload(data.RoleData),
		CreatedAt:       data.CreatedAt,
		CreatedBy:       data.CreatedBy,
		UpdatedAt:       data.UpdatedAt,
		UpdatedBy:       data.UpdatedBy,
		IsDeleted:       data.IsDeleted,
	}
	return
}
