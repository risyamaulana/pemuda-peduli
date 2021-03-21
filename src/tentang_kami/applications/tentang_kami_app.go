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
	"pemuda-peduli/src/tentang_kami/domain"
	"pemuda-peduli/src/tentang_kami/infrastructure/repository"
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

// TentangKamiApp ...
type TentangKamiApp struct {
	interfaces.IApplication
}

// NewTentangKamiApp ...
func NewTentangKamiApp() *TentangKamiApp {
	// Place where we init infrastructure, repo etc
	s := TentangKamiApp{}
	return &s
}

// Initialize will be called when application run
func (s *TentangKamiApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("Tentang Kami app initialized")
}

// Destroy will be called when app shutdowns
func (s *TentangKamiApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("Tentang Kami app released...")
}

// Route declaration
func (s *TentangKamiApp) addRoute(r *router.Router) {
	r.POST("/tentang-kami/create", middleware.CheckAuthToken(createTentangKami))

	r.PUT("/tentang-kami/{id}", middleware.CheckAuthToken(updateTentangKami))

	r.POST("/tentang-kami/list", middleware.CheckAuthToken(findTentangKamis))
	r.GET("/tentang-kami/{id}", middleware.CheckAuthToken(getTentangKami))

	r.DELETE("/tentang-kami/{id}", middleware.CheckAuthToken(deleteTentangKami))
}

// ============== Handler for each route start here ============

func createTentangKami(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewTentangKamiRepository(DB)
	if err := domain.CreateTentangKami(ctx, &repo, &data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateTentangKami(ctx *fasthttp.RequestCtx) {
	tentangKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
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
	repo := repository.NewTentangKamiRepository(DB)
	responseData, err := domain.UpdateTentangKami(ctx, &repo, data, tentangKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteTentangKami(ctx *fasthttp.RequestCtx) {
	tentangKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewTentangKamiRepository(DB)
	responseData, err := domain.DeleteTentangKami(ctx, &repo, tentangKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findTentangKamis(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewTentangKamiRepository(DB)

	responseData, count, err := domain.FindTentangKami(ctx, &repo, &data)

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
	response := []ReadTentangKami{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getTentangKami(ctx *fasthttp.RequestCtx) {
	tentangKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewTentangKamiRepository(DB)
	responseData, err := domain.GetTentangKami(ctx, &repo, tentangKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
