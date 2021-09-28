package applications

import (
	"errors"
	"fmt"
	"log"
	"math"
	"pemuda-peduli/src/berita/domain"
	"pemuda-peduli/src/berita/infrastructure/repository"
	"pemuda-peduli/src/common/handler"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/common/interfaces"
	"pemuda-peduli/src/common/middleware"
	"pemuda-peduli/src/common/utility"
	"strconv"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var DB *db.ConnectTo

// BeritaApp ...
type BeritaApp struct {
	interfaces.IApplication
}

// NewBeritaApp ...
func NewBeritaApp(db *db.ConnectTo) *BeritaApp {
	// Place where we init infrastructure, repo etc
	s := BeritaApp{}
	DB = db
	return &s
}

// Initialize will be called when application run
func (s *BeritaApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("Berita app initialized")
}

// Destroy will be called when app shutdowns
func (s *BeritaApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("Berita app released...")
}

// Route declaration
func (s *BeritaApp) addRoute(r *router.Router) {
	r.POST("/berita/create", middleware.CheckAuthToken(DB, createBerita))

	r.PUT("/berita/{id}", middleware.CheckAuthToken(DB, updateBerita))
	r.PUT("/berita/publish/{id}", middleware.CheckAuthToken(DB, publishBerita))
	r.PUT("/berita/hide/{id}", middleware.CheckAuthToken(DB, hideBerita))

	r.POST("/berita/list", middleware.CheckAuthToken(DB, findBeritas))
	r.GET("/berita/{id}", middleware.CheckAuthToken(DB, getBerita))

	r.DELETE("/berita/{id}", middleware.CheckAuthToken(DB, deleteBerita))

	r.GET("/berita/list-tag", getListTag)
}

// ============== Handler for each route start here ============

func createBerita(ctx *fasthttp.RequestCtx) {
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

	data, detail := payload.ToEntity()
	repo := repository.NewBeritaRepository(DB)
	responseData, err := domain.CreateBerita(ctx, &repo, &data, &detail)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(responseData, true), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateBerita(ctx *fasthttp.RequestCtx) {
	beritaID := fmt.Sprintf("%s", ctx.UserValue("id"))
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

	data, detail := payload.ToEntity()
	repo := repository.NewBeritaRepository(DB)
	responseData, err := domain.UpdateBerita(ctx, &repo, data, detail, beritaID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData, true), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func publishBerita(ctx *fasthttp.RequestCtx) {
	beritaID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewBeritaRepository(DB)
	responseData, err := domain.PublishBerita(ctx, &repo, beritaID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData, false), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func hideBerita(ctx *fasthttp.RequestCtx) {
	beritaID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewBeritaRepository(DB)
	responseData, err := domain.HideBerita(ctx, &repo, beritaID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData, false), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteBerita(ctx *fasthttp.RequestCtx) {
	beritaID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewBeritaRepository(DB)
	responseData, err := domain.DeleteBerita(ctx, &repo, beritaID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData, false), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findBeritas(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewBeritaRepository(DB)

	responseData, count, err := domain.FindBerita(ctx, &repo, &data)

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
	response := []ReadBerita{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp, false))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getBerita(ctx *fasthttp.RequestCtx) {
	beritaID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewBeritaRepository(DB)
	responseData, err := domain.GetBerita(ctx, &repo, beritaID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData, true), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func getListTag(ctx *fasthttp.RequestCtx) {
	repo := repository.NewBeritaRepository(DB)
	responseData, err := domain.GetListTag(ctx, &repo)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(responseData, nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
