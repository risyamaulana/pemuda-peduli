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
	"pemuda-peduli/src/kontak_kami/domain"
	"pemuda-peduli/src/kontak_kami/infrastructure/repository"
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

// KontakKamiApp ...
type KontakKamiApp struct {
	interfaces.IApplication
}

// NewKontakKamiApp ...
func NewKontakKamiApp() *KontakKamiApp {
	// Place where we init infrastructure, repo etc
	s := KontakKamiApp{}
	return &s
}

// Initialize will be called when application run
func (s *KontakKamiApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("Kontak Kami app initialized")
}

// Destroy will be called when app shutdowns
func (s *KontakKamiApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("Kontak Kami app released...")
}

// Route declaration
func (s *KontakKamiApp) addRoute(r *router.Router) {
	r.POST("/kontak-kami/create", middleware.CheckAuthToken(createKontakKami))

	r.PUT("/kontak-kami/{id}", middleware.CheckAuthToken(updateKontakKami))
	r.PUT("/kontak-kami/publish/{id}", middleware.CheckAuthToken(publishKontakKami))
	r.PUT("/kontak-kami/hide/{id}", middleware.CheckAuthToken(hideKontakKami))

	r.POST("/kontak-kami/list", middleware.CheckAuthToken(findKontakKamis))
	r.GET("/kontak-kami/{id}", middleware.CheckAuthToken(getKontakKami))

	r.DELETE("/kontak-kami/{id}", middleware.CheckAuthToken(deleteKontakKami))
}

// ============== Handler for each route start here ============

func createKontakKami(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewKontakKamiRepository(DB)
	if err := domain.CreateKontakKami(ctx, &repo, &data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateKontakKami(ctx *fasthttp.RequestCtx) {
	kontakKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
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
	repo := repository.NewKontakKamiRepository(DB)
	responseData, err := domain.UpdateKontakKami(ctx, &repo, data, kontakKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func publishKontakKami(ctx *fasthttp.RequestCtx) {
	kontakKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewKontakKamiRepository(DB)
	responseData, err := domain.PublishKontakKami(ctx, &repo, kontakKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func hideKontakKami(ctx *fasthttp.RequestCtx) {
	kontakKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewKontakKamiRepository(DB)
	responseData, err := domain.HideKontakKami(ctx, &repo, kontakKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteKontakKami(ctx *fasthttp.RequestCtx) {
	kontakKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewKontakKamiRepository(DB)
	responseData, err := domain.DeleteKontakKami(ctx, &repo, kontakKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findKontakKamis(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewKontakKamiRepository(DB)

	responseData, count, err := domain.FindKontakKami(ctx, &repo, &data)

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
	response := []ReadKontakKami{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getKontakKami(ctx *fasthttp.RequestCtx) {
	kontakKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewKontakKamiRepository(DB)
	responseData, err := domain.GetKontakKami(ctx, &repo, kontakKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
