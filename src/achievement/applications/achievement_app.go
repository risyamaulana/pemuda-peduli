package applications

import (
	"errors"
	"fmt"
	"log"
	"math"
	"pemuda-peduli/src/achievement/domain"
	"pemuda-peduli/src/achievement/infrastructure/repository"
	"pemuda-peduli/src/common/handler"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/common/interfaces"
	"pemuda-peduli/src/common/middleware"
	"pemuda-peduli/src/common/utility"
	"strconv"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var (
	DB *db.ConnectTo
)

// db init hardcoded temporary for testing
func init() {
	DB = db.NewDBConnectionFactory(0)
}

// AchievementApp ...
type AchievementApp struct {
	interfaces.IApplication
}

// NewAchievementApp ...
func NewAchievementApp() *AchievementApp {
	// Place where we init infrastructure, repo etc
	s := AchievementApp{}
	return &s
}

// Initialize will be called when application run
func (s *AchievementApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("Achievement app initialized")
}

// Destroy will be called when app shutdowns
func (s *AchievementApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("Achievement app released...")
}

// Route declaration
func (s *AchievementApp) addRoute(r *router.Router) {
	r.POST("/achievement/create", middleware.CheckAuthToken(createAchievement))

	r.PUT("/achievement/{id}", middleware.CheckAuthToken(updateAchievement))
	r.PUT("/achievement/publish/{id}", middleware.CheckAuthToken(publishAchievement))
	r.PUT("/achievement/hide/{id}", middleware.CheckAuthToken(hideAchievement))

	r.POST("/achievement/list", middleware.CheckAuthToken(findAchievements))
	r.GET("/achievement/{id}", middleware.CheckAuthToken(getAchievement))

	r.DELETE("/achievement/{id}", middleware.CheckAuthToken(deleteAchievement))
}

// ============== Handler for each route start here ============

func createAchievement(ctx *fasthttp.RequestCtx) {
	payload, err := GetCreatePayload(ctx.Request.Body())
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Bad JSON Payload"))))
		log.Println("Error Bad Request JSON Payload:", err)
		return
	}

	if err := payload.Validate(); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	data := payload.ToEntity()
	repo := repository.NewAchievementRepository(DB)
	if err := domain.CreateAchievement(ctx, &repo, &data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateAchievement(ctx *fasthttp.RequestCtx) {
	achievementID := fmt.Sprintf("%s", ctx.UserValue("id"))
	payload, err := GetCreatePayload(ctx.Request.Body())
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Bad JSON Payload"))))
		log.Println("Error Bad Request JSON Payload:", err)
		return
	}

	if err := payload.Validate(); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	data := payload.ToEntity()
	repo := repository.NewAchievementRepository(DB)
	responseData, err := domain.UpdateAchievement(ctx, &repo, data, achievementID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func publishAchievement(ctx *fasthttp.RequestCtx) {
	achievementID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewAchievementRepository(DB)
	responseData, err := domain.PublishAchievement(ctx, &repo, achievementID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func hideAchievement(ctx *fasthttp.RequestCtx) {
	achievementID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewAchievementRepository(DB)
	responseData, err := domain.HideAchievement(ctx, &repo, achievementID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteAchievement(ctx *fasthttp.RequestCtx) {
	achievementID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewAchievementRepository(DB)
	responseData, err := domain.DeleteAchievement(ctx, &repo, achievementID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findAchievements(ctx *fasthttp.RequestCtx) {
	payload, err := GetQueryPayload(ctx.Request.Body())
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Bad JSON Payload"))))
		log.Println("Error Bad Request JSON Payload:", err)
		return
	}

	if err := payload.Validate(); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	data := payload.ToEntity()
	repo := repository.NewAchievementRepository(DB)

	responseData, count, err := domain.FindAchievement(ctx, &repo, &data)

	// TOTAL PAGE
	limit, _ := strconv.Atoi(payload.Limit)
	page, _ := strconv.Atoi(payload.Offset)
	pageTotal := math.Ceil(float64(count) / float64(limit))

	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(nil, err, page, limit, int(pageTotal), count)))
		log.Println(err)
		return
	}

	// Return data as json
	response := []ReadAchievement{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getAchievement(ctx *fasthttp.RequestCtx) {
	achievementID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewAchievementRepository(DB)
	responseData, err := domain.GetAchievement(ctx, &repo, achievementID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
