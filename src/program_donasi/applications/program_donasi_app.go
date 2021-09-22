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
	"pemuda-peduli/src/program_donasi/domain"
	"pemuda-peduli/src/program_donasi/infrastructure/repository"
	"strconv"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var DB *db.ConnectTo

// ProgramDonasiApp ...
type ProgramDonasiApp struct {
	interfaces.IApplication
}

// NewProgramDonasiApp ...
func NewProgramDonasiApp(db *db.ConnectTo) *ProgramDonasiApp {
	// Place where we init infrastructure, repo etc
	s := ProgramDonasiApp{}
	DB = db
	return &s
}

// Initialize will be called when application run
func (s *ProgramDonasiApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("ProgramDonasi app initialized")
}

// Destroy will be called when app shutdowns
func (s *ProgramDonasiApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("ProgramDonasi app released...")
}

// Route declaration
func (s *ProgramDonasiApp) addRoute(r *router.Router) {
	r.POST("/program-donasi/create", middleware.CheckAuthToken(DB, createProgramDonasi))

	r.PUT("/program-donasi/{id}", middleware.CheckAuthToken(DB, updateProgramDonasi))
	r.PUT("/program-donasi/publish/{id}", middleware.CheckAuthToken(DB, publishProgramDonasi))
	r.PUT("/program-donasi/hide/{id}", middleware.CheckAuthToken(DB, hideProgramDonasi))

	r.POST("/program-donasi/list", middleware.CheckAuthToken(DB, findProgramDonasis))
	r.GET("/program-donasi/{id}", middleware.CheckAuthToken(DB, getProgramDonasi))

	r.DELETE("/program-donasi/{id}", middleware.CheckAuthToken(DB, deleteProgramDonasi))
}

// ============== Handler for each route start here ============

func createProgramDonasi(ctx *fasthttp.RequestCtx) {
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
	responseData, err := domain.CreateProgramDonasi(ctx, DB, &data, &dataDetail)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateProgramDonasi(ctx *fasthttp.RequestCtx) {
	programDonasiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	payload, err := GetUpdatePayload(ctx.Request.Body())
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
	responseData, err := domain.UpdateProgramDonasi(ctx, DB, data, dataDetail, programDonasiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func publishProgramDonasi(ctx *fasthttp.RequestCtx) {
	programDonasiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewProgramDonasiRepository(DB)
	responseData, err := domain.PublishProgramDonasi(ctx, &repo, programDonasiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func hideProgramDonasi(ctx *fasthttp.RequestCtx) {
	programDonasiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewProgramDonasiRepository(DB)
	responseData, err := domain.HideProgramDonasi(ctx, &repo, programDonasiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteProgramDonasi(ctx *fasthttp.RequestCtx) {
	programDonasiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewProgramDonasiRepository(DB)
	responseData, err := domain.DeleteProgramDonasi(ctx, &repo, programDonasiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findProgramDonasis(ctx *fasthttp.RequestCtx) {
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

	responseData, count, err := domain.FindProgramDonasi(ctx, DB, &data)

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
	response := []ReadProgramDonasi{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getProgramDonasi(ctx *fasthttp.RequestCtx) {
	programDonasiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	responseData, err := domain.GetProgramDonasi(ctx, DB, programDonasiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
