package applications

import (
	"errors"
	"fmt"
	"log"
	"math"
	"pemuda-peduli/src/common/handler"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/common/interfaces"
	"pemuda-peduli/src/common/middleware"
	"pemuda-peduli/src/common/utility"
	"pemuda-peduli/src/team/domain"
	"pemuda-peduli/src/team/infrastructure/repository"
	"strconv"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var DB *db.ConnectTo

// TeamApp ...
type TeamApp struct {
	interfaces.IApplication
}

// NewTeamApp ...
func NewTeamApp(db *db.ConnectTo) *TeamApp {
	// Place where we init infrastructure, repo etc
	s := TeamApp{}
	DB = db
	return &s
}

// Initialize will be called when application run
func (s *TeamApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("Team app initialized")
}

// Destroy will be called when app shutdowns
func (s *TeamApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("Team app released...")
}

// Route declaration
func (s *TeamApp) addRoute(r *router.Router) {
	r.POST("/team/create", middleware.CheckAuthToken(DB, createTeam))

	r.PUT("/team/{id}", middleware.CheckAuthToken(DB, updateTeam))
	r.PUT("/team/publish/{id}", middleware.CheckAuthToken(DB, publishTeam))
	r.PUT("/team/hide/{id}", middleware.CheckAuthToken(DB, hideTeam))

	r.POST("/team/list", middleware.CheckAuthToken(DB, findTeams))
	r.GET("/team/{id}", middleware.CheckAuthToken(DB, getTeam))

	r.DELETE("/team/{id}", middleware.CheckAuthToken(DB, deleteTeam))
}

// ============== Handler for each route start here ============

func createTeam(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewTeamRepository(DB)
	if err := domain.CreateTeam(ctx, &repo, &data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateTeam(ctx *fasthttp.RequestCtx) {
	teamID := fmt.Sprintf("%s", ctx.UserValue("id"))
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
	repo := repository.NewTeamRepository(DB)
	responseData, err := domain.UpdateTeam(ctx, &repo, data, teamID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func publishTeam(ctx *fasthttp.RequestCtx) {
	teamID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewTeamRepository(DB)
	responseData, err := domain.PublishTeam(ctx, &repo, teamID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func hideTeam(ctx *fasthttp.RequestCtx) {
	teamID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewTeamRepository(DB)
	responseData, err := domain.HideTeam(ctx, &repo, teamID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteTeam(ctx *fasthttp.RequestCtx) {
	teamID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewTeamRepository(DB)
	responseData, err := domain.DeleteTeam(ctx, &repo, teamID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findTeams(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewTeamRepository(DB)

	responseData, count, err := domain.FindTeam(ctx, &repo, &data)

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
	response := []ReadTeam{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getTeam(ctx *fasthttp.RequestCtx) {
	teamID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewTeamRepository(DB)
	responseData, err := domain.GetTeam(ctx, &repo, teamID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
