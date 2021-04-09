package applications

import (
	"errors"
	"fmt"
	"log"
	"math"
	"pemuda-peduli/src/beneficaries/domain"
	"pemuda-peduli/src/beneficaries/infrastructure/repository"
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

// BeneficariesApp ...
type BeneficariesApp struct {
	interfaces.IApplication
}

// NewBeneficariesApp ...
func NewBeneficariesApp() *BeneficariesApp {
	// Place where we init infrastructure, repo etc
	s := BeneficariesApp{}
	return &s
}

// Initialize will be called when application run
func (s *BeneficariesApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("Beneficaries app initialized")
}

// Destroy will be called when app shutdowns
func (s *BeneficariesApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("Beneficaries app released...")
}

// Route declaration
func (s *BeneficariesApp) addRoute(r *router.Router) {
	r.POST("/beneficaries/create", middleware.CheckAuthToken(createBeneficaries))

	r.PUT("/beneficaries/{id}", middleware.CheckAuthToken(updateBeneficaries))
	r.PUT("/beneficaries/publish/{id}", middleware.CheckAuthToken(publishBeneficaries))
	r.PUT("/beneficaries/hide/{id}", middleware.CheckAuthToken(hideBeneficaries))

	r.POST("/beneficaries/list", middleware.CheckAuthToken(findBeneficariess))
	r.GET("/beneficaries/{id}", middleware.CheckAuthToken(getBeneficaries))

	r.DELETE("/beneficaries/{id}", middleware.CheckAuthToken(deleteBeneficaries))
}

// ============== Handler for each route start here ============

func createBeneficaries(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewBeneficariesRepository(DB)
	if err := domain.CreateBeneficaries(ctx, &repo, &data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateBeneficaries(ctx *fasthttp.RequestCtx) {
	beneficariesID := fmt.Sprintf("%s", ctx.UserValue("id"))
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
	repo := repository.NewBeneficariesRepository(DB)
	responseData, err := domain.UpdateBeneficaries(ctx, &repo, data, beneficariesID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func publishBeneficaries(ctx *fasthttp.RequestCtx) {
	beneficariesID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewBeneficariesRepository(DB)
	responseData, err := domain.PublishBeneficaries(ctx, &repo, beneficariesID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func hideBeneficaries(ctx *fasthttp.RequestCtx) {
	beneficariesID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewBeneficariesRepository(DB)
	responseData, err := domain.HideBeneficaries(ctx, &repo, beneficariesID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteBeneficaries(ctx *fasthttp.RequestCtx) {
	beneficariesID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewBeneficariesRepository(DB)
	responseData, err := domain.DeleteBeneficaries(ctx, &repo, beneficariesID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findBeneficariess(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewBeneficariesRepository(DB)

	responseData, count, err := domain.FindBeneficaries(ctx, &repo, &data)

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
	response := []ReadBeneficaries{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getBeneficaries(ctx *fasthttp.RequestCtx) {
	beneficariesID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewBeneficariesRepository(DB)
	responseData, err := domain.GetBeneficaries(ctx, &repo, beneficariesID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
