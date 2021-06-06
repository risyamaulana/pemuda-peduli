package applications

import (
	"encoding/json"
	"pemuda-peduli/src/achievement/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateAchievement struct {
	AchievementName  string `json:"achievement_name" valid:"required"`
	AchievementTotal int64  `json:"achievement_total"`
	Description      string `json:"description"`
}

type UpdateAchievement struct {
	AchievementName  string `json:"achievement_name" valid:"required"`
	AchievementTotal int64  `json:"achievement_total"`
	Description      string `json:"description"`
}

type AchievementQuery struct {
	Limit         string                   `json:"limit" valid:"required"`
	Offset        string                   `json:"offset" valid:"required"`
	Filter        []AchievementFilterQuery `json:"filters"`
	Order         string                   `json:"order"`
	Sort          string                   `json:"sort"`
	CreatedAtFrom string                   `json:"created_at_from"`
	CreatedAtTo   string                   `json:"created_at_to"`
	PublishAtFrom string                   `json:"publish_at_from"`
	PublishAtTo   string                   `json:"publish_at_to"`
}

type AchievementFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadAchievement struct {
	IDPPCPAchievement string     `json:"id"`
	AchievementName   string     `json:"achievement_name"`
	AchievementTotal  int64      `json:"achievement_total"`
	Description       string     `json:"description"`
	Status            string     `json:"status"`
	CreatedAt         time.Time  `json:"created_at"`
	CreatedBy         *string    `json:"created_by"`
	UpdatedAt         *time.Time `json:"updated_at"`
	UpdatedBy         *string    `json:"updated_by"`
	PublishedAt       *time.Time `json:"published_at"`
	PublishedBy       *string    `json:"published_by"`
	IsDeleted         bool       `json:"is_deleted"`
}

func GetCreatePayload(body []byte) (payload CreateAchievement, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateAchievement, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload AchievementQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateAchievement) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdateAchievement) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r AchievementQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateAchievement) ToEntity() (data entity.AchievementEntity) {
	data = entity.AchievementEntity{
		AchievementName:  r.AchievementName,
		AchievementTotal: r.AchievementTotal,
		Description:      r.Description,
		CreatedAt:        time.Now(),
	}
	return
}

func (r UpdateAchievement) ToEntity() (data entity.AchievementEntity) {
	data = entity.AchievementEntity{
		AchievementName:  r.AchievementName,
		AchievementTotal: r.AchievementTotal,
		Description:      r.Description,
	}
	return
}

func (r AchievementQuery) ToEntity() (data entity.AchievementQueryEntity) {
	filters := []entity.AchievementFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.AchievementFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.AchievementQueryEntity{
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

func ToPayload(data entity.AchievementEntity) (response ReadAchievement) {
	response = ReadAchievement{
		IDPPCPAchievement: data.IDPPCPAchievement,
		AchievementName:   data.AchievementName,
		AchievementTotal:  data.AchievementTotal,
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
