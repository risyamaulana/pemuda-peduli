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
	"pemuda-peduli/src/partner_kami/domain"
	"pemuda-peduli/src/partner_kami/infrastructure/repository"
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

// PartnerKamiApp ...
type PartnerKamiApp struct {
	interfaces.IApplication
}

// NewPartnerKamiApp ...
func NewPartnerKamiApp() *PartnerKamiApp {
	// Place where we init infrastructure, repo etc
	s := PartnerKamiApp{}
	return &s
}

// Initialize will be called when application run
func (s *PartnerKamiApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("Partner Kami app initialized")
}

// Destroy will be called when app shutdowns
func (s *PartnerKamiApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("Partner Kami app released...")
}

// Route declaration
func (s *PartnerKamiApp) addRoute(r *router.Router) {
	r.POST("/partner-kami/create", middleware.CheckAuthToken(createPartnerKami))

	r.PUT("/partner-kami/{id}", middleware.CheckAuthToken(updatePartnerKami))
	r.PUT("/partner-kami/publish/{id}", middleware.CheckAuthToken(publishPartnerKami))
	r.PUT("/partner-kami/hide/{id}", middleware.CheckAuthToken(hidePartnerKami))

	r.POST("/partner-kami/list", middleware.CheckAuthToken(findPartnerKamis))
	r.GET("/partner-kami/{id}", middleware.CheckAuthToken(getPartnerKami))

	r.DELETE("/partner-kami/{id}", middleware.CheckAuthToken(deletePartnerKami))
}

// ============== Handler for each route start here ============

func createPartnerKami(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewPartnerKamiRepository(DB)
	if err := domain.CreatePartnerKami(ctx, &repo, &data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updatePartnerKami(ctx *fasthttp.RequestCtx) {
	partnerKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
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
	repo := repository.NewPartnerKamiRepository(DB)
	responseData, err := domain.UpdatePartnerKami(ctx, &repo, data, partnerKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func publishPartnerKami(ctx *fasthttp.RequestCtx) {
	partnerKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewPartnerKamiRepository(DB)
	responseData, err := domain.PublishPartnerKami(ctx, &repo, partnerKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func hidePartnerKami(ctx *fasthttp.RequestCtx) {
	partnerKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewPartnerKamiRepository(DB)
	responseData, err := domain.HidePartnerKami(ctx, &repo, partnerKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deletePartnerKami(ctx *fasthttp.RequestCtx) {
	partnerKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewPartnerKamiRepository(DB)
	responseData, err := domain.DeletePartnerKami(ctx, &repo, partnerKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findPartnerKamis(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewPartnerKamiRepository(DB)

	responseData, count, err := domain.FindPartnerKami(ctx, &repo, &data)

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
	response := []ReadPartnerKami{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getPartnerKami(ctx *fasthttp.RequestCtx) {
	partnerKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewPartnerKamiRepository(DB)
	responseData, err := domain.GetPartnerKami(ctx, &repo, partnerKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
