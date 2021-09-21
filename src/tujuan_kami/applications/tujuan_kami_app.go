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
	"pemuda-peduli/src/tujuan_kami/domain"
	"pemuda-peduli/src/tujuan_kami/infrastructure/repository"
	"strconv"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var DB *db.ConnectTo

// TujuanKamiApp ...
type TujuanKamiApp struct {
	interfaces.IApplication
}

// NewTujuanKamiApp ...
func NewTujuanKamiApp(db *db.ConnectTo) *TujuanKamiApp {
	// Place where we init infrastructure, repo etc
	s := TujuanKamiApp{}
	DB = db
	return &s
}

// Initialize will be called when application run
func (s *TujuanKamiApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("Tujuan Kami app initialized")
}

// Destroy will be called when app shutdowns
func (s *TujuanKamiApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("Tujuan Kami app released...")
}

// Route declaration
func (s *TujuanKamiApp) addRoute(r *router.Router) {
	r.POST("/tujuan-kami/create", middleware.CheckAuthToken(DB, createTujuanKami))

	r.PUT("/tujuan-kami/{id}", middleware.CheckAuthToken(DB, updateTujuanKami))
	r.PUT("/tujuan-kami/publish/{id}", middleware.CheckAuthToken(DB, publishTujuanKami))
	r.PUT("/tujuan-kami/hide/{id}", middleware.CheckAuthToken(DB, hideTujuanKami))

	r.POST("/tujuan-kami/list", middleware.CheckAuthToken(DB, findTujuanKamis))
	r.GET("/tujuan-kami/{id}", middleware.CheckAuthToken(DB, getTujuanKami))

	r.DELETE("/tujuan-kami/{id}", middleware.CheckAuthToken(DB, deleteTujuanKami))
}

// ============== Handler for each route start here ============

func createTujuanKami(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewTujuanKamiRepository(DB)
	if err := domain.CreateTujuanKami(ctx, &repo, &data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateTujuanKami(ctx *fasthttp.RequestCtx) {
	tujuanKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
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
	repo := repository.NewTujuanKamiRepository(DB)
	responseData, err := domain.UpdateTujuanKami(ctx, &repo, data, tujuanKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func publishTujuanKami(ctx *fasthttp.RequestCtx) {
	tujuanKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewTujuanKamiRepository(DB)
	responseData, err := domain.PublishTujuanKami(ctx, &repo, tujuanKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func hideTujuanKami(ctx *fasthttp.RequestCtx) {
	tujuanKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewTujuanKamiRepository(DB)
	responseData, err := domain.HideTujuanKami(ctx, &repo, tujuanKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteTujuanKami(ctx *fasthttp.RequestCtx) {
	tujuanKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewTujuanKamiRepository(DB)
	responseData, err := domain.DeleteTujuanKami(ctx, &repo, tujuanKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findTujuanKamis(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewTujuanKamiRepository(DB)

	responseData, count, err := domain.FindTujuanKami(ctx, &repo, &data)

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
	response := []ReadTujuanKami{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getTujuanKami(ctx *fasthttp.RequestCtx) {
	tujuanKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewTujuanKamiRepository(DB)
	responseData, err := domain.GetTujuanKami(ctx, &repo, tujuanKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
