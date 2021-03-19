package applications

import (
	"errors"
	"fmt"
	"log"
	"math"
	"pemuda-peduli/src/album/domain"
	"pemuda-peduli/src/album/infrastructure/repository"
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

// AlbumApp ...
type AlbumApp struct {
	interfaces.IApplication
}

// NewAlbumApp ...
func NewAlbumApp() *AlbumApp {
	// Place where we init infrastructure, repo etc
	s := AlbumApp{}
	return &s
}

// Initialize will be called when application run
func (s *AlbumApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("Album app initialized")
}

// Destroy will be called when app shutdowns
func (s *AlbumApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("Album app released...")
}

// Route declaration
func (s *AlbumApp) addRoute(r *router.Router) {
	r.POST("/album/create", middleware.CheckAuthToken(createAlbum))

	r.PUT("/album/{id}", middleware.CheckAuthToken(updateAlbum))
	r.PUT("/album/publish/{id}", middleware.CheckAuthToken(publishAlbum))
	r.PUT("/album/hide/{id}", middleware.CheckAuthToken(hideAlbum))

	r.POST("/album/list", middleware.CheckAuthToken(findAlbums))
	r.GET("/album/{id}", middleware.CheckAuthToken(getAlbum))

	r.DELETE("/album/{id}", middleware.CheckAuthToken(deleteAlbum))
}

// ============== Handler for each route start here ============

func createAlbum(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewAlbumRepository(DB)
	if err := domain.CreateAlbum(ctx, &repo, &data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateAlbum(ctx *fasthttp.RequestCtx) {
	albumID := fmt.Sprintf("%s", ctx.UserValue("id"))
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
	repo := repository.NewAlbumRepository(DB)
	responseData, err := domain.UpdateAlbum(ctx, &repo, data, albumID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func publishAlbum(ctx *fasthttp.RequestCtx) {
	albumID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewAlbumRepository(DB)
	responseData, err := domain.PublishAlbum(ctx, &repo, albumID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func hideAlbum(ctx *fasthttp.RequestCtx) {
	albumID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewAlbumRepository(DB)
	responseData, err := domain.HideAlbum(ctx, &repo, albumID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteAlbum(ctx *fasthttp.RequestCtx) {
	albumID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewAlbumRepository(DB)
	responseData, err := domain.DeleteAlbum(ctx, &repo, albumID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findAlbums(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewAlbumRepository(DB)

	responseData, count, err := domain.FindAlbum(ctx, &repo, &data)

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
	response := []ReadAlbum{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getAlbum(ctx *fasthttp.RequestCtx) {
	albumID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewAlbumRepository(DB)
	responseData, err := domain.GetAlbum(ctx, &repo, albumID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
