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
	"pemuda-peduli/src/program_donasi_fundraiser/domain"
	"strconv"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var DB *db.ConnectTo

// ProgramDonasiFundraiserApp ...
type ProgramDonasiFundraiserApp struct {
	interfaces.IApplication
}

// NewProgramDonasiFundraiserApp ...
func NewProgramDonasiFundraiserApp(db *db.ConnectTo) *ProgramDonasiFundraiserApp {
	// Place where we init infrastructure, repo etc
	s := ProgramDonasiFundraiserApp{}
	DB = db
	return &s
}

// Initialize will be called when application run
func (s *ProgramDonasiFundraiserApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("ProgramDonasiRutin app initialized")
}

// Destroy will be called when app shutdowns
func (s *ProgramDonasiFundraiserApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("ProgramDonasiFundraiser app released...")
}

func (s *ProgramDonasiFundraiserApp) addRoute(r *router.Router) {
	r.POST("/fundraiser/create/{id}", middleware.CheckUserToken(DB, createProgramDonasiFundraiser))
	r.POST("/fundraiser/list", middleware.CheckAuthToken(DB, findProgramDonasiFundraiser))
	r.GET("/fundraiser/{id}", middleware.CheckAuthToken(DB, getProgramDonasiFundraiser))
	r.GET("/fundraiser/seo/{seo_url}", middleware.CheckAuthToken(DB, getProgramDonasiFundraiserSeo))

}

func createProgramDonasiFundraiser(ctx *fasthttp.RequestCtx) {
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
	data := payload.ToEntity(ctx.UserValue("id").(string))

	err = domain.CreateProgramDonasiFundraiser(ctx, DB, &data)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findProgramDonasiFundraiser(ctx *fasthttp.RequestCtx) {
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

	responses, count, err := domain.FindProgramDonasiFundraiser(ctx, DB, payload.ToEntity())
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
	response := []ReadProgramDonasiFundraiser{}
	for _, resp := range responses {
		response = append(response, ToPayload(resp))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getProgramDonasiFundraiser(ctx *fasthttp.RequestCtx) {
	data, err := domain.GetProgramDonasiFundraiser(ctx, DB, ctx.UserValue("id").(string))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func getProgramDonasiFundraiserSeo(ctx *fasthttp.RequestCtx) {
	data, err := domain.GetProgramDonasiFundraiserSeo(ctx, DB, ctx.UserValue("seo_url").(string))

	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
