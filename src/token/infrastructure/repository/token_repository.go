package repository

import (
	"context"
	"log"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/common/utility"
	"pemuda-peduli/src/token/domain/entity"
	"pemuda-peduli/src/token/domain/interfaces"
	"strconv"
	"strings"
)

// TokenRepository
type TokenRepository struct {
	db *db.ConnectTo
	interfaces.ITokenRepository
}

// NewTokenRepository
func NewTokenRepository(db *db.ConnectTo) TokenRepository {
	return TokenRepository{db: db}
}

// Create data token
func (c *TokenRepository) Insert(ctx context.Context, data *entity.TokenEntity) (err error) {

	tx := c.db.DBExec.MustBegin()

	sqlItem := `INSERT INTO pp_token `
	var strField strings.Builder
	var strValue strings.Builder
	filedItem := utility.GetNamedStruct(*data)
	for _, field := range filedItem {
		if field != "id" {
			strField.WriteString(field + ",")
			strValue.WriteString(":" + field + ",")
		}
	}

	sqlItem += "(" + strings.TrimSuffix(strField.String(), ",") + ")" + " VALUES(" + strings.TrimSuffix(strValue.String(), ",") + ")"
	resp, err := tx.NamedExec(sqlItem, data)
	if err != nil {
		log.Println("Error insert pp_token:", err)
		tx.Rollback()
		return
	}

	tx.Commit()
	data.ID, _ = resp.LastInsertId()
	return
}

func (c TokenRepository) Update(ctx context.Context, data *entity.TokenEntity) (err error) {
	tx := c.db.DBRead.MustBegin()

	// Update Data delivery order
	sql := `Update pp_token SET `
	var str strings.Builder
	fields := utility.GetNamedStruct(*data)
	for _, field := range fields {
		if field == "id" || field == "created_at" {
			continue
		}
		str.WriteString(field + "=:" + field + ", ")
	}
	queryCondition := strings.TrimSuffix(str.String(), ", ")

	sql += queryCondition + " WHERE id = " + strconv.FormatInt(data.ID, 10)
	_, err = tx.NamedExec(sql, data)
	err = tx.Commit()

	return
}

// CheckAvailableData
func (c *TokenRepository) CheckDevice(deviceID, deviceType string) (data entity.TokenEntity, err error) {
	if err = c.db.DBRead.Get(&data, "SELECT * FROM pp_token WHERE device_id = $1 AND device_type = $2", deviceID, deviceType); err != nil {
		log.Println("Error get pp_token:", err)
		return
	}
	return
}

func (c *TokenRepository) CheckTokenDevice(token, deviceID, deviceType string) (data entity.TokenEntity, err error) {
	if err = c.db.DBRead.Get(&data, "SELECT * FROM pp_token WHERE device_id = $1 AND device_type = $2 AND token = $3", deviceID, deviceType, token); err != nil {
		log.Println("Error get zoolyfe_token:", err)
		return
	}
	return
}
