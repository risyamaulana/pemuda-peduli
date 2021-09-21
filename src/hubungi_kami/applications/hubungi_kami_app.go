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
	"pemuda-peduli/src/hubungi_kami/domain"
	"pemuda-peduli/src/hubungi_kami/infrastructure/repository"
	"strconv"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var DB *db.ConnectTo

// HubungiKamiApp ...
type HubungiKamiApp struct {
	interfaces.IApplication
}

// NewHubungiKamiApp ...
func NewHubungiKamiApp(db *db.ConnectTo) *HubungiKamiApp {
	// Place where we init infrastructure, repo etc
	s := HubungiKamiApp{}
	DB = db
	return &s
}

// Initialize will be called when application run
func (s *HubungiKamiApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("Hubungi Kami app initialized")
}

// Destroy will be called when app shutdowns
func (s *HubungiKamiApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("Hubungi Kami app released...")
}

// Route declaration
func (s *HubungiKamiApp) addRoute(r *router.Router) {
	r.POST("/hubungi-kami/create", middleware.CheckAuthToken(DB, createHubungiKami))

	r.PUT("/hubungi-kami/{id}", middleware.CheckAuthToken(DB, updateHubungiKami))
	r.PUT("/hubungi-kami/publish/{id}", middleware.CheckAuthToken(DB, publishHubungiKami))
	r.PUT("/hubungi-kami/hide/{id}", middleware.CheckAuthToken(DB, hideHubungiKami))

	r.POST("/hubungi-kami/list", middleware.CheckAuthToken(DB, findHubungiKamis))
	r.GET("/hubungi-kami/{id}", middleware.CheckAuthToken(DB, getHubungiKami))

	r.DELETE("/hubungi-kami/{id}", middleware.CheckAuthToken(DB, deleteHubungiKami))
}

// ============== Handler for each route start here ============

func createHubungiKami(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewHubungiKamiRepository(DB)
	if err := domain.CreateHubungiKami(ctx, &repo, &data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateHubungiKami(ctx *fasthttp.RequestCtx) {
	hubungiKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
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
	repo := repository.NewHubungiKamiRepository(DB)
	responseData, err := domain.UpdateHubungiKami(ctx, &repo, data, hubungiKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func publishHubungiKami(ctx *fasthttp.RequestCtx) {
	hubungiKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewHubungiKamiRepository(DB)
	responseData, err := domain.PublishHubungiKami(ctx, &repo, hubungiKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func hideHubungiKami(ctx *fasthttp.RequestCtx) {
	hubungiKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewHubungiKamiRepository(DB)
	responseData, err := domain.HideHubungiKami(ctx, &repo, hubungiKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteHubungiKami(ctx *fasthttp.RequestCtx) {
	hubungiKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewHubungiKamiRepository(DB)
	responseData, err := domain.DeleteHubungiKami(ctx, &repo, hubungiKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findHubungiKamis(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewHubungiKamiRepository(DB)

	responseData, count, err := domain.FindHubungiKami(ctx, &repo, &data)

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
	response := []ReadHubungiKami{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getHubungiKami(ctx *fasthttp.RequestCtx) {
	hubungiKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewHubungiKamiRepository(DB)
	responseData, err := domain.GetHubungiKami(ctx, &repo, hubungiKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
