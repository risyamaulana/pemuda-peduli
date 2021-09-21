package applications

import (
	"errors"
	"fmt"
	"log"
	"math"
	"pemuda-peduli/src/banner/domain"
	"pemuda-peduli/src/banner/infrastructure/repository"
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

// BannerApp ...
type BannerApp struct {
	interfaces.IApplication
}

// NewBannerApp ...
func NewBannerApp(db *db.ConnectTo) *BannerApp {
	// Place where we init infrastructure, repo etc
	s := BannerApp{}
	DB = db
	return &s
}

// Initialize will be called when application run
func (s *BannerApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("Banner app initialized")
}

// Destroy will be called when app shutdowns
func (s *BannerApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("Banner app released...")
}

// Route declaration
func (s *BannerApp) addRoute(r *router.Router) {
	r.POST("/banner/create", middleware.CheckAuthToken(DB, createBanner))

	r.PUT("/banner/{id}", middleware.CheckAuthToken(DB, updateBanner))
	r.PUT("/banner/publish/{id}", middleware.CheckAuthToken(DB, publishBanner))
	r.PUT("/banner/hide/{id}", middleware.CheckAuthToken(DB, hideBanner))

	r.POST("/banner/list", middleware.CheckAuthToken(DB, findBanners))
	r.GET("/banner/{id}", middleware.CheckAuthToken(DB, getBanner))

	r.DELETE("/banner/{id}", middleware.CheckAuthToken(DB, deleteBanner))
}

// ============== Handler for each route start here ============

func createBanner(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewBannerRepository(DB)
	if err := domain.CreateBanner(ctx, &repo, &data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateBanner(ctx *fasthttp.RequestCtx) {
	bannerID := fmt.Sprintf("%s", ctx.UserValue("id"))
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
	repo := repository.NewBannerRepository(DB)
	responseData, err := domain.UpdateBanner(ctx, &repo, data, bannerID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func publishBanner(ctx *fasthttp.RequestCtx) {
	bannerID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewBannerRepository(DB)
	responseData, err := domain.PublishBanner(ctx, &repo, bannerID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func hideBanner(ctx *fasthttp.RequestCtx) {
	bannerID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewBannerRepository(DB)
	responseData, err := domain.HideBanner(ctx, &repo, bannerID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteBanner(ctx *fasthttp.RequestCtx) {
	bannerID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewBannerRepository(DB)
	responseData, err := domain.DeleteBanner(ctx, &repo, bannerID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findBanners(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewBannerRepository(DB)

	responseData, count, err := domain.FindBanner(ctx, &repo, &data)

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
	response := []ReadBanner{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getBanner(ctx *fasthttp.RequestCtx) {
	bannerID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewBannerRepository(DB)
	responseData, err := domain.GetBanner(ctx, &repo, bannerID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
