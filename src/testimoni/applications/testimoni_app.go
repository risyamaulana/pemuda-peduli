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
	"pemuda-peduli/src/testimoni/domain"
	"pemuda-peduli/src/testimoni/infrastructure/repository"
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

// TestimoniApp ...
type TestimoniApp struct {
	interfaces.IApplication
}

// NewTestimoniApp ...
func NewTestimoniApp() *TestimoniApp {
	// Place where we init infrastructure, repo etc
	s := TestimoniApp{}
	return &s
}

// Initialize will be called when application run
func (s *TestimoniApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("Testimoni app initialized")
}

// Destroy will be called when app shutdowns
func (s *TestimoniApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("Testimoni app released...")
}

// Route declaration
func (s *TestimoniApp) addRoute(r *router.Router) {
	r.POST("/testimoni/create", middleware.CheckAuthToken(createTestimoni))

	r.PUT("/testimoni/{id}", middleware.CheckAuthToken(updateTestimoni))
	r.PUT("/testimoni/publish/{id}", middleware.CheckAuthToken(publishTestimoni))
	r.PUT("/testimoni/hide/{id}", middleware.CheckAuthToken(hideTestimoni))

	r.POST("/testimoni/list", middleware.CheckAuthToken(findTestimonis))
	r.GET("/testimoni/{id}", middleware.CheckAuthToken(getTestimoni))

	r.DELETE("/testimoni/{id}", middleware.CheckAuthToken(deleteTestimoni))
}

// ============== Handler for each route start here ============

func createTestimoni(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewTestimoniRepository(DB)
	if err := domain.CreateTestimoni(ctx, &repo, &data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateTestimoni(ctx *fasthttp.RequestCtx) {
	testimoniID := fmt.Sprintf("%s", ctx.UserValue("id"))
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
	repo := repository.NewTestimoniRepository(DB)
	responseData, err := domain.UpdateTestimoni(ctx, &repo, data, testimoniID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func publishTestimoni(ctx *fasthttp.RequestCtx) {
	testimoniID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewTestimoniRepository(DB)
	responseData, err := domain.PublishTestimoni(ctx, &repo, testimoniID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func hideTestimoni(ctx *fasthttp.RequestCtx) {
	testimoniID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewTestimoniRepository(DB)
	responseData, err := domain.HideTestimoni(ctx, &repo, testimoniID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteTestimoni(ctx *fasthttp.RequestCtx) {
	testimoniID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewTestimoniRepository(DB)
	responseData, err := domain.DeleteTestimoni(ctx, &repo, testimoniID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findTestimonis(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewTestimoniRepository(DB)

	responseData, count, err := domain.FindTestimoni(ctx, &repo, &data)

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
	response := []ReadTestimoni{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getTestimoni(ctx *fasthttp.RequestCtx) {
	testimoniID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewTestimoniRepository(DB)
	responseData, err := domain.GetTestimoni(ctx, &repo, testimoniID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
