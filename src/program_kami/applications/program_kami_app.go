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
	"pemuda-peduli/src/program_kami/domain"
	"pemuda-peduli/src/program_kami/infrastructure/repository"
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

// ProgramKamiApp ...
type ProgramKamiApp struct {
	interfaces.IApplication
}

// NewProgramKamiApp ...
func NewProgramKamiApp() *ProgramKamiApp {
	// Place where we init infrastructure, repo etc
	s := ProgramKamiApp{}
	return &s
}

// Initialize will be called when application run
func (s *ProgramKamiApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("Program Kami app initialized")
}

// Destroy will be called when app shutdowns
func (s *ProgramKamiApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("Program Kami app released...")
}

// Route declaration
func (s *ProgramKamiApp) addRoute(r *router.Router) {
	r.POST("/program-kami/create", middleware.CheckAuthToken(createProgramKami))

	r.PUT("/program-kami/{id}", middleware.CheckAuthToken(updateProgramKami))
	r.PUT("/program-kami/publish/{id}", middleware.CheckAuthToken(publishProgramKami))
	r.PUT("/program-kami/hide/{id}", middleware.CheckAuthToken(hideProgramKami))

	r.POST("/program-kami/list", middleware.CheckAuthToken(findProgramKamis))
	r.GET("/program-kami/{id}", middleware.CheckAuthToken(getProgramKami))

	r.DELETE("/program-kami/{id}", middleware.CheckAuthToken(deleteProgramKami))
}

// ============== Handler for each route start here ============

func createProgramKami(ctx *fasthttp.RequestCtx) {
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

	data, dataDetail := payload.ToEntity()
	repo := repository.NewProgramKamiRepository(DB)
	responseData, err := domain.CreateProgramKami(ctx, &repo, &data, &dataDetail)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(responseData, true), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateProgramKami(ctx *fasthttp.RequestCtx) {
	programKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
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

	data, dataDetail := payload.ToEntity()
	repo := repository.NewProgramKamiRepository(DB)
	responseData, err := domain.UpdateProgramKami(ctx, &repo, data, dataDetail, programKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData, true), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func publishProgramKami(ctx *fasthttp.RequestCtx) {
	programKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewProgramKamiRepository(DB)
	responseData, err := domain.PublishProgramKami(ctx, &repo, programKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData, false), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func hideProgramKami(ctx *fasthttp.RequestCtx) {
	programKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewProgramKamiRepository(DB)
	responseData, err := domain.HideProgramKami(ctx, &repo, programKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData, false), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteProgramKami(ctx *fasthttp.RequestCtx) {
	programKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewProgramKamiRepository(DB)
	responseData, err := domain.DeleteProgramKami(ctx, &repo, programKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData, false), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findProgramKamis(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewProgramKamiRepository(DB)

	responseData, count, err := domain.FindProgramKami(ctx, &repo, &data)

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
	response := []ReadProgramKami{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp, false))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getProgramKami(ctx *fasthttp.RequestCtx) {
	programKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewProgramKamiRepository(DB)
	responseData, err := domain.GetProgramKami(ctx, &repo, programKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData, true), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
