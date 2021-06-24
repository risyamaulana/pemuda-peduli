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
	"pemuda-peduli/src/program_donasi_rutin/domain"
	"pemuda-peduli/src/program_donasi_rutin/infrastructure/repository"
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

// ProgramDonasiRutinApp ...
type ProgramDonasiRutinApp struct {
	interfaces.IApplication
}

// NewProgramDonasiRutinApp ...
func NewProgramDonasiRutinApp() *ProgramDonasiRutinApp {
	// Place where we init infrastructure, repo etc
	s := ProgramDonasiRutinApp{}
	return &s
}

// Initialize will be called when application run
func (s *ProgramDonasiRutinApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("ProgramDonasiRutin app initialized")
}

// Destroy will be called when app shutdowns
func (s *ProgramDonasiRutinApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("ProgramDonasiRutin app released...")
}

// Route declaration
func (s *ProgramDonasiRutinApp) addRoute(r *router.Router) {
	r.POST("/program-donasi-rutin/create", middleware.CheckAuthToken(createProgramDonasiRutin))

	r.PUT("/program-donasi-rutin/{id}", middleware.CheckAuthToken(updateProgramDonasiRutin))
	r.PUT("/program-donasi-rutin/publish/{id}", middleware.CheckAuthToken(publishProgramDonasiRutin))
	r.PUT("/program-donasi-rutin/hide/{id}", middleware.CheckAuthToken(hideProgramDonasiRutin))

	r.POST("/program-donasi-rutin/list", middleware.CheckAuthToken(findProgramDonasiRutins))
	r.GET("/program-donasi-rutin/{id}", middleware.CheckAuthToken(getProgramDonasiRutin))

	r.DELETE("/program-donasi-rutin/{id}", middleware.CheckAuthToken(deleteProgramDonasiRutin))
}

// ============== Handler for each route start here ============

func createProgramDonasiRutin(ctx *fasthttp.RequestCtx) {
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
	responseData, err := domain.CreateProgramDonasiRutin(ctx, DB, &data, &dataDetail)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateProgramDonasiRutin(ctx *fasthttp.RequestCtx) {
	programDonasiID := fmt.Sprintf("%s", ctx.UserValue("id"))
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
	responseData, err := domain.EditProgramDonasiRutin(ctx, DB, data, dataDetail, programDonasiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func publishProgramDonasiRutin(ctx *fasthttp.RequestCtx) {
	programDonasiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewProgramDonasiRutinRepository(DB)
	responseData, err := domain.PublishProgramDonasiRutin(ctx, &repo, programDonasiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func hideProgramDonasiRutin(ctx *fasthttp.RequestCtx) {
	programDonasiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewProgramDonasiRutinRepository(DB)
	responseData, err := domain.HideProgramDonasiRutin(ctx, &repo, programDonasiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteProgramDonasiRutin(ctx *fasthttp.RequestCtx) {
	programDonasiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewProgramDonasiRutinRepository(DB)
	responseData, err := domain.DeleteProgramDonasiRutin(ctx, &repo, programDonasiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findProgramDonasiRutins(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewProgramDonasiRutinRepository(DB)

	responseData, count, err := domain.FindProgramDonasiRutin(ctx, &repo, &data)

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
	response := []ReadProgramDonasiRutin{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getProgramDonasiRutin(ctx *fasthttp.RequestCtx) {
	programDonasiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewProgramDonasiRutinRepository(DB)
	responseData, err := domain.GetProgramDonasiRutin(ctx, &repo, programDonasiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
