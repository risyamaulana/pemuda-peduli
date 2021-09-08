package repository

import (
	"context"
	"errors"
	"log"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/common/utility"
	"pemuda-peduli/src/user/domain/entity"
	"pemuda-peduli/src/user/domain/interfaces"
	"strconv"
	"strings"
)

// UserRepository
type UserRepository struct {
	db *db.ConnectTo
	interfaces.IUserRepository
}

// NewUserRepository
func NewUserRepository(db *db.ConnectTo) UserRepository {
	return UserRepository{db: db}
}

// Create User
func (c *UserRepository) Insert(ctx context.Context, data *entity.UserEntity) (err error) {
	tx := c.db.DBExec.MustBegin()

	sql := `INSERT INTO pp_user `
	var strField strings.Builder
	var strValue strings.Builder
	filedItem := utility.GetNamedStruct(*data)
	for _, field := range filedItem {
		if field != "id" {
			strField.WriteString(field + ",")
			strValue.WriteString(":" + field + ",")
		}
	}

	sql += "(" + strings.TrimSuffix(strField.String(), ",") + ")" + " VALUES(" + strings.TrimSuffix(strValue.String(), ",") + ")"
	resp, err := tx.NamedExec(sql, data)
	if err != nil {
		log.Println("Error insert pp_user:", err)
		tx.Rollback()
		return
	}

	tx.Commit()
	data.ID, _ = resp.LastInsertId()
	return
}

// Update
func (c *UserRepository) Update(ctx context.Context, data entity.UserEntity) (response entity.UserEntity, err error) {
	tx := c.db.DBExec.MustBegin()

	// Update Data delivery order
	sql := `Update pp_user SET `
	var str strings.Builder
	fields := utility.GetNamedStruct(data)
	for _, field := range fields {
		if field == "id" || field == "id_user" || field == "created_at" {
			continue
		}
		str.WriteString(field + "=:" + field + ", ")
	}
	queryCondition := strings.TrimSuffix(str.String(), ", ")

	sql += queryCondition + " WHERE id = '" + strconv.FormatInt(data.ID, 10) + "'"
	log.Print("QUERY : ", sql)
	_, err = tx.NamedExec(sql, data)
	if err != nil {
		log.Println("Error insert pp_user:", err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	response = data

	return
}

// Get
func (c *UserRepository) Get(ctx context.Context, id string) (response entity.UserEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_user WHERE id_user = $1 AND is_deleted = false", id); err != nil {
		return
	}
	return
}

func (c *UserRepository) GetByEmail(ctx context.Context, email string) (response entity.UserEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_user WHERE email = $1 AND is_deleted = false", email); err != nil {
		return
	}
	return
}

func (c *UserRepository) GetForLogin(ctx context.Context, username string) (response entity.UserEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_user WHERE (username = $1 OR email = $1 OR phone_number = $1) AND is_deleted = false", username); err != nil {
		return
	}
	if response.IsDeleted {
		err = errors.New("Failed: user not found")
		return
	}
	return
}

func (c *UserRepository) GetDuplicateCheck(ctx context.Context, username, phoneNumber, email string) (response entity.UserEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_user WHERE (username = $1 OR email = $2 OR phone_number = $3) AND is_deleted = false", username, email, phoneNumber); err != nil {
		return
	}
	if response.IsDeleted {
		err = errors.New("Failed: user not found")
		return
	}
	return
}

func (c *UserRepository) GetByToken(ctx context.Context, token string) (response entity.UserEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_user WHERE token_reset = $1", token); err != nil {
		return
	}
	return
}
