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
	"pemuda-peduli/src/program_incidental/domain"
	"pemuda-peduli/src/program_incidental/infrastructure/repository"
	"strconv"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var DB *db.ConnectTo

// ProgramIncidentalApp ...
type ProgramIncidentalApp struct {
	interfaces.IApplication
}

// NewProgramIncidentalApp ...
func NewProgramIncidentalApp(db *db.ConnectTo) *ProgramIncidentalApp {
	// Place where we init infrastructure, repo etc
	s := ProgramIncidentalApp{}
	DB = db
	return &s
}

// Initialize will be called when application run
func (s *ProgramIncidentalApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("Program Kami app initialized")
}

// Destroy will be called when app shutdowns
func (s *ProgramIncidentalApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("Program Kami app released...")
}

// Route declaration
func (s *ProgramIncidentalApp) addRoute(r *router.Router) {
	r.POST("/program-incidental/create", middleware.CheckAuthToken(DB, createProgramIncidental))

	r.PUT("/program-incidental/{id}", middleware.CheckAuthToken(DB, updateProgramIncidental))
	r.PUT("/program-incidental/publish/{id}", middleware.CheckAuthToken(DB, publishProgramIncidental))
	r.PUT("/program-incidental/hide/{id}", middleware.CheckAuthToken(DB, hideProgramIncidental))

	r.POST("/program-incidental/list", middleware.CheckAuthToken(DB, findProgramIncidentals))
	r.GET("/program-incidental/{id}", middleware.CheckAuthToken(DB, getProgramIncidental))

	r.DELETE("/program-incidental/{id}", middleware.CheckAuthToken(DB, deleteProgramIncidental))
}

// ============== Handler for each route start here ============

func createProgramIncidental(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewProgramIncidentalRepository(DB)
	responseData, err := domain.CreateProgramIncidental(ctx, &repo, &data)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(responseData, true), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateProgramIncidental(ctx *fasthttp.RequestCtx) {
	programIncidentalID := fmt.Sprintf("%s", ctx.UserValue("id"))
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
	repo := repository.NewProgramIncidentalRepository(DB)
	responseData, err := domain.UpdateProgramIncidental(ctx, &repo, data, programIncidentalID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData, true), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func publishProgramIncidental(ctx *fasthttp.RequestCtx) {
	programIncidentalID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewProgramIncidentalRepository(DB)
	responseData, err := domain.PublishProgramIncidental(ctx, &repo, programIncidentalID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData, false), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func hideProgramIncidental(ctx *fasthttp.RequestCtx) {
	programIncidentalID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewProgramIncidentalRepository(DB)
	responseData, err := domain.HideProgramIncidental(ctx, &repo, programIncidentalID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData, false), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteProgramIncidental(ctx *fasthttp.RequestCtx) {
	programIncidentalID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewProgramIncidentalRepository(DB)
	responseData, err := domain.DeleteProgramIncidental(ctx, &repo, programIncidentalID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData, false), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findProgramIncidentals(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewProgramIncidentalRepository(DB)

	responseData, count, err := domain.FindProgramIncidental(ctx, &repo, &data)

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
	response := []ReadProgramIncidental{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp, false))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getProgramIncidental(ctx *fasthttp.RequestCtx) {
	programIncidentalID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewProgramIncidentalRepository(DB)
	responseData, err := domain.GetProgramIncidental(ctx, &repo, programIncidentalID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData, true), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
