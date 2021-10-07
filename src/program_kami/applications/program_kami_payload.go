package applications

import (
	"encoding/json"
	"errors"
	"pemuda-peduli/src/program_kami/domain/entity"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateProgramKami struct {
	Title                string        `json:"title" valid:"required"`
	SubTitle             string        `json:"sub_title" valid:"required"`
	Content              string        `json:"content" valid:"required"`
	Achievements         []Achievement `json:"achievements"`
	Tag                  string        `json:"tag"`
	ThumbnailImageURL    string        `json:"thumbnail_image_url" valid:"url"`
	BeneficariesImageURL []string      `json:"beneficaries_image_url" valid:"url"`
	Description          string        `json:"description"`
}

type UpdateProgramKami struct {
	Title                string        `json:"title" valid:"required"`
	SubTitle             string        `json:"sub_title" valid:"required"`
	Content              string        `json:"content" valid:"required"`
	Achievements         []Achievement `json:"achievements"`
	Tag                  string        `json:"tag"`
	ThumbnailImageURL    string        `json:"thumbnail_image_url" valid:"url"`
	BeneficariesImageURL []string      `json:"beneficaries_image_url" valid:"url"`
	Description          string        `json:"description"`
}

type ProgramKamiQuery struct {
	Limit         string                   `json:"limit" valid:"required"`
	Offset        string                   `json:"offset" valid:"required"`
	Filter        []ProgramKamiFilterQuery `json:"filters"`
	Order         string                   `json:"order"`
	Sort          string                   `json:"sort"`
	CreatedAtFrom string                   `json:"created_at_from"`
	CreatedAtTo   string                   `json:"created_at_to"`
	PublishAtFrom string                   `json:"publish_at_from"`
	PublishAtTo   string                   `json:"publish_at_to"`
}

type ProgramKamiFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadProgramKami struct {
	IDPPCPProgramKami    string        `json:"id"`
	Title                string        `json:"title"`
	SubTitle             string        `json:"sub_title"`
	Tag                  string        `json:"tag"`
	ThumbnailImageURL    string        `json:"thumbnail_image_url"`
	BeneficariesImageURL []string      `json:"beneficaries_image_url"`
	Achievements         []Achievement `json:"achievements"`
	// Detail            *ReadProgramKamiDetail `json:"detail,omitempty"`
	Content     string     `json:"content,omitempty"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   *string    `json:"created_by"`
	UpdatedAt   *time.Time `json:"updated_at"`
	UpdatedBy   *string    `json:"updated_by"`
	PublishedAt *time.Time `json:"published_at"`
	PublishedBy *string    `json:"published_by"`
	IsDeleted   bool       `json:"is_deleted"`
}

type ReadProgramKamiDetail struct {
	Content string `json:"content"`
}

type Achievement struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func GetCreatePayload(body []byte) (payload CreateProgramKami, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateProgramKami, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload ProgramKamiQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateProgramKami) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}

	if len(r.BeneficariesImageURL) > 3 {
		err = errors.New("Failed, limit beneficaries image url is 3")
		return
	}
	return
}

func (r UpdateProgramKami) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}

	if len(r.BeneficariesImageURL) > 3 {
		err = errors.New("Failed, limit beneficaries image url is 3")
		return
	}

	return
}

func (r ProgramKamiQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateProgramKami) ToEntity() (data entity.ProgramKamiEntity, dataDetail entity.ProgramKamiDetailEntity) {
	var beneficariesURL strings.Builder
	for _, beneficariesImageURL := range r.BeneficariesImageURL {
		beneficariesURL.WriteString(beneficariesImageURL + "|")
	}

	beneficariesURLValue := strings.TrimSuffix(beneficariesURL.String(), "|")

	var achievements strings.Builder
	for _, achievement := range r.Achievements {
		achievements.WriteString((achievement.Label + "&" + achievement.Value + "|"))
	}
	achievementsValue := strings.TrimSuffix(achievements.String(), "|")

	data = entity.ProgramKamiEntity{
		Title:                r.Title,
		SubTitle:             r.SubTitle,
		Tag:                  r.Tag,
		ThumbnailImageURL:    r.ThumbnailImageURL,
		BeneficariesImageURL: beneficariesURLValue,
		Achievements:         achievementsValue,
		Description:          r.Description,
		CreatedAt:            time.Now().UTC(),
	}
	dataDetail = entity.ProgramKamiDetailEntity{
		Content: r.Content,
		Tag:     r.Tag,
	}

	return
}

func (r UpdateProgramKami) ToEntity() (data entity.ProgramKamiEntity, dataDetail entity.ProgramKamiDetailEntity) {
	var beneficariesURL strings.Builder
	for _, beneficariesImageURL := range r.BeneficariesImageURL {
		beneficariesURL.WriteString(beneficariesImageURL + "|")
	}

	beneficariesURLValue := strings.TrimSuffix(beneficariesURL.String(), "|")

	var achievements strings.Builder
	for _, achievement := range r.Achievements {
		achievements.WriteString((achievement.Label + "&" + achievement.Value + "|"))
	}
	achievementsValue := strings.TrimSuffix(achievements.String(), "|")

	data = entity.ProgramKamiEntity{
		Title:                r.Title,
		SubTitle:             r.SubTitle,
		Tag:                  r.Tag,
		BeneficariesImageURL: beneficariesURLValue,
		Achievements:         achievementsValue,
		ThumbnailImageURL:    r.ThumbnailImageURL,
		Description:          r.Description,
	}

	dataDetail = entity.ProgramKamiDetailEntity{
		Content: r.Content,
		Tag:     r.Tag,
	}
	return
}

func (r ProgramKamiQuery) ToEntity() (data entity.ProgramKamiQueryEntity) {
	filters := []entity.ProgramKamiFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.ProgramKamiFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.ProgramKamiQueryEntity{
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

func ToPayload(data entity.ProgramKamiEntity, isDetail bool) (response ReadProgramKami) {
	// var detail *ReadProgramKamiDetail
	// if isDetail {
	// 	detail = &ReadProgramKamiDetail{
	// 		Content: data.Detail.Content,
	// 	}
	// }

	beneficariesUrl := strings.Split(data.BeneficariesImageURL, "|")

	achievements := []Achievement{}
	if data.Achievements != "" {
		achievementStr := strings.Split(data.Achievements, "|")

		for _, achievement := range achievementStr {
			achievementData := strings.Split(achievement, "&")
			achievements = append(achievements, Achievement{
				Label: achievementData[0],
				Value: achievementData[1],
			})
		}
	}

	response = ReadProgramKami{
		IDPPCPProgramKami:    data.IDPPCPProgramKami,
		Title:                data.Title,
		SubTitle:             data.SubTitle,
		Tag:                  data.Tag,
		ThumbnailImageURL:    data.ThumbnailImageURL,
		BeneficariesImageURL: beneficariesUrl,
		Content:              data.Detail.Content,
		Achievements:         achievements,
		// Detail:            detail,
		Description: data.Description,
		Status:      data.Status,
		CreatedAt:   data.CreatedAt,
		CreatedBy:   data.CreatedBy,
		UpdatedAt:   data.UpdatedAt,
		UpdatedBy:   data.UpdatedBy,
		PublishedAt: data.PublishedAt,
		PublishedBy: data.PublishedBy,
		IsDeleted:   data.IsDeleted,
	}
	return
}
