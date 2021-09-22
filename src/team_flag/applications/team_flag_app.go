package applications

import (
	"fmt"
	"log"
	"pemuda-peduli/src/common/handler"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/common/interfaces"
	"pemuda-peduli/src/common/utility"
	"pemuda-peduli/src/team_flag/domain"
	"pemuda-peduli/src/team_flag/infrastructure/repository"
	"strconv"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var DB *db.ConnectTo

// TeamFlagApp ...
type TeamFlagApp struct {
	interfaces.IApplication
}

// NewTeamFlagApp ...
func NewTeamFlagApp(db *db.ConnectTo) *TeamFlagApp {
	// Place where we init infrastructure, repo etc
	s := TeamFlagApp{}
	DB = db
	return &s
}

// Initialize will be called when application run
func (s *TeamFlagApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("Team app initialized")
}

// Destroy will be called when app shutdowns
func (s *TeamFlagApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("Team app released...")
}

// Route declaration
func (s *TeamFlagApp) addRoute(r *router.Router) {
	r.POST("/team-flag/list", listTeamFlag)
	r.GET("/team-flag/{id}", getTeamFlag)
}

// ============== Handler for each route start here ============

func listTeamFlag(ctx *fasthttp.RequestCtx) {
	repo := repository.NewTeamFlagRepository(DB)

	responseData, err := domain.ListTeamFlag(ctx, &repo)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}

	// Return data as json
	response := []ReadTeamFlag{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp))
	}
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func getTeamFlag(ctx *fasthttp.RequestCtx) {
	teamID, err := strconv.ParseInt(ctx.UserValue("id").(string), 10, 64)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println("Error Bad Request JSON Payload:", err)
		return
	}
	repo := repository.NewTeamFlagRepository(DB)
	responseData, err := domain.GetTeamFlag(ctx, &repo, teamID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
