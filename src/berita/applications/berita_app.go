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

var (
	DB *db.ConnectTo
)

// db init hardcoded temporary for testing
func init() {
	DB = db.NewDBConnectionFactory(0)
}

// BeritaApp ...
type BeritaApp struct {
	interfaces.IApplication
}

// NewBeritaApp ...
func NewBeritaApp() *BeritaApp {
	// Place where we init infrastructure, repo etc
	s := BeritaApp{}
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
	r.POST("/berita/create", middleware.CheckAuthToken(createBerita))

	r.PUT("/berita/{id}", middleware.CheckAuthToken(updateBerita))
	r.PUT("/berita/publish/{id}", middleware.CheckAuthToken(publishBerita))
	r.PUT("/berita/hide/{id}", middleware.CheckAuthToken(hideBerita))

	r.POST("/berita/list", middleware.CheckAuthToken(findBeritas))
	r.GET("/berita/{id}", middleware.CheckAuthToken(getBerita))

	r.DELETE("/berita/{id}", middleware.CheckAuthToken(deleteBerita))
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

	data := payload.ToEntity()
	repo := repository.NewBeritaRepository(DB)
	if err := domain.CreateBerita(ctx, &repo, &data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
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

	data := payload.ToEntity()
	repo := repository.NewBeritaRepository(DB)
	responseData, err := domain.UpdateBerita(ctx, &repo, data, beritaID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
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
	response := handler.DefaultResponse(ToPayload(responseData), nil)
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
	response := handler.DefaultResponse(ToPayload(responseData), nil)
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
	response := handler.DefaultResponse(ToPayload(responseData), nil)
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
		response = append(response, ToPayload(resp))
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
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
