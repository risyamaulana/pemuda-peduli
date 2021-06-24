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

var (
	DB *db.ConnectTo
)

// db init hardcoded temporary for testing
func init() {
	DB = db.NewDBConnectionFactory(0)
}

// ProgramDonasiApp ...
type ProgramDonasiApp struct {
	interfaces.IApplication
}

// NewProgramDonasiApp ...
func NewProgramDonasiApp() *ProgramDonasiApp {
	// Place where we init infrastructure, repo etc
	s := ProgramDonasiApp{}
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
	r.POST("/program-donasi/create", middleware.CheckAuthToken(createProgramDonasi))

	r.PUT("/program-donasi/{id}", middleware.CheckAuthToken(updateProgramDonasi))
	r.PUT("/program-donasi/publish/{id}", middleware.CheckAuthToken(publishProgramDonasi))
	r.PUT("/program-donasi/hide/{id}", middleware.CheckAuthToken(hideProgramDonasi))

	r.POST("/program-donasi/list", middleware.CheckAuthToken(findProgramDonasis))
	r.GET("/program-donasi/{id}", middleware.CheckAuthToken(getProgramDonasi))

	r.DELETE("/program-donasi/{id}", middleware.CheckAuthToken(deleteProgramDonasi))
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
	repo := repository.NewProgramDonasiRepository(DB)
	responseData, err := domain.CreateProgramDonasi(ctx, &repo, &data, &dataDetail)
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
	repo := repository.NewProgramDonasiRepository(DB)
	responseData, err := domain.UpdateProgramDonasi(ctx, &repo, data, dataDetail, programDonasiID)
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
	repo := repository.NewProgramDonasiRepository(DB)

	responseData, count, err := domain.FindProgramDonasi(ctx, &repo, &data)

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
	repo := repository.NewProgramDonasiRepository(DB)
	responseData, err := domain.GetProgramDonasi(ctx, &repo, programDonasiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
