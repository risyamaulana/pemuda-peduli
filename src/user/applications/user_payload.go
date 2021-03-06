package applications

import (
	"context"
	"encoding/json"
	"errors"
	"pemuda-peduli/src/common/utility"
	"pemuda-peduli/src/user/domain/entity"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
)

type Readuser struct {
	ID            string     `json:"id"` // id user (UUID)
	Username      string     `json:"username"`
	Email         string     `json:"email"`
	NamaLengkap   string     `json:"nama_lengkap"`
	NamaPanggilan string     `json:"nama_panggilan"`
	Alamat        string     `json:"alamat"`
	PhoneNumber   string     `json:"phone_number"`
	IsDeleted     bool       `json:"is_deleted"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}

func ToPayload(data entity.UserEntity) (response Readuser) {
	response = Readuser{
		ID:            data.IDUser,
		Username:      data.Username,
		Email:         data.Email,
		NamaLengkap:   data.NamaLengkap,
		NamaPanggilan: data.NamaPanggilan,
		Alamat:        data.Alamat,
		PhoneNumber:   data.PhoneNumber,
		IsDeleted:     data.IsDeleted,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}
	return
}

// ============================ Register ============================

type RegisterUser struct {
	Email         string `json:"email" valid:"required,email"`
	PhoneNumber   string `json:"phone_number" valid:"required"`
	Password      string `json:"password" valid:"required,minstringlength(8)"`
	ConfPassword  string `json:"conf_password" valid:"required,minstringlength(8)"`
	NamaLengkap   string `json:"nama_lengkap" valid:"required"`
	NamaPanggilan string `json:"nama_panggilan" valid:"required"`
	Alamat        string `json:"alamat"`
}

// ============================ Forgot Password ============================
type ForgotPassword struct {
	Email string `json:"email" valid:"required,email"`
}

// ============================ Reset Password ============================
type ResetPassword struct {
	NewPassword     string `json:"new_password" valid:"required,minstringlength(8)"`
	ConfNewPassword string `json:"conf_new_password" valid:"required,minstringlength(8)"`
	Token           string `json:"token" valid:"required"`
}

func GetRegisterUserPayload(body []byte) (payload RegisterUser, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetForgotPasswordPayload(body []byte) (payload ForgotPassword, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetResetPasswordPayload(body []byte) (payload ResetPassword, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r RegisterUser) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}

	if r.Password != r.ConfPassword {
		err = errors.New("Failed: Password not match")
		return
	}

	return
}

func (r ForgotPassword) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r ResetPassword) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}

	if r.NewPassword != r.ConfNewPassword {
		err = errors.New("Failed: Password not match")
		return
	}
	return
}

func (r RegisterUser) ToEntity() (data entity.UserEntity) {
	username := strings.Split(r.Email, "@")

	salt := utility.GenerateSalt(4)
	password := utility.GeneratePass(salt, r.Password)

	phoneNumber := utility.FormatPhoneNumber(r.PhoneNumber)

	data = entity.UserEntity{
		IDUser:        utility.GetUUID(),
		Username:      username[0] + "-" + utility.GenerateSalt(5),
		Salt:          salt,
		Password:      password,
		Email:         r.Email,
		NamaLengkap:   r.NamaLengkap,
		NamaPanggilan: r.NamaPanggilan,
		Alamat:        r.Alamat,
		PhoneNumber:   phoneNumber,
		IsDeleted:     false,
		CreatedAt:     time.Now().UTC(),
	}

	return
}

// ============================ Update User ============================
type UpdateUser struct {
	NamaLengkap   string `json:"nama_lengkap" valid:"required"`
	NamaPanggilan string `json:"nama_panggilan" valid:"required"`
	Alamat        string `json:"alamat"`
}

func GetUpdateUserPayload(body []byte) (payload UpdateUser, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r UpdateUser) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}

	return
}

func (r UpdateUser) ToEntity(ctx context.Context) (data entity.UserEntity) {
	currentDate := time.Now().UTC()

	userID := ctx.Value("user_id").(string)

	data = entity.UserEntity{
		IDUser:        userID,
		NamaLengkap:   r.NamaLengkap,
		NamaPanggilan: r.NamaPanggilan,
		Alamat:        r.Alamat,
		IsDeleted:     false,
		UpdatedAt:     &currentDate,
	}

	return
}

// ============================ Change Password ============================
type ChangePassword struct {
	OldPassword  string `json:"old_password" valid:"required,minstringlength(8)"`
	Password     string `json:"new_password" valid:"required,minstringlength(8)"`
	ConfPassword string `json:"confirm_new_password" valid:"required,minstringlength(8)"`
}

func GetChangePasswordPayload(body []byte) (payload ChangePassword, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r ChangePassword) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}

	if r.Password != r.ConfPassword {
		err = errors.New("Failed: Password not match")
		return
	}

	return
}

func (r ChangePassword) ToEntity(ctx context.Context) (data entity.UserEntity) {
	currentDate := time.Now().UTC()

	userID := ctx.Value("user_id").(string)

	salt := utility.GenerateSalt(4)
	password := utility.GeneratePass(salt, r.Password)

	data = entity.UserEntity{
		IDUser:    userID,
		Salt:      salt,
		Password:  password,
		IsDeleted: false,
		UpdatedAt: &currentDate,
	}

	return
}
