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
	"pemuda-peduli/src/penggalang_dana/domain"
	"strconv"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var DB *db.ConnectTo

// PenggalangDanaApp ...
type PenggalangDanaApp struct {
	interfaces.IApplication
}

// NewPenggalangDanaApp ...
func NewPenggalangDanaApp(db *db.ConnectTo) *PenggalangDanaApp {
	// Place where we init infrastructure, repo etc
	s := PenggalangDanaApp{}
	DB = db
	return &s
}

// Initialize will be called when application run
func (s *PenggalangDanaApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("PenggalangDana app initialized")
}

// Destroy will be called when app shutdowns
func (s *PenggalangDanaApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("PenggalangDana app released...")
}

// Route declaration
func (s *PenggalangDanaApp) addRoute(r *router.Router) {
	r.POST("/penggalang-dana/create", middleware.CheckAuthToken(DB, createPenggalangDana))

	r.PUT("/penggalang-dana/{id}", middleware.CheckAuthToken(DB, updatePenggalangDana))
	r.PUT("/penggalang-dana/verified/{id}", middleware.CheckAuthToken(DB, verifiedPenggalangDana))

	r.POST("/penggalang-dana/list", middleware.CheckAuthToken(DB, findPenggalangDanas))
	r.GET("/penggalang-dana/{id}", middleware.CheckAuthToken(DB, getPenggalangDana))

	r.DELETE("/penggalang-dana/{id}", middleware.CheckAuthToken(DB, deletePenggalangDana))
}

// ============== Handler for each route start here ============

func createPenggalangDana(ctx *fasthttp.RequestCtx) {
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
	if err := domain.CreatePenggalangDana(ctx, DB, &data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updatePenggalangDana(ctx *fasthttp.RequestCtx) {
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
	responseData, err := domain.EditPenggalangDana(ctx, DB, data)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func verifiedPenggalangDana(ctx *fasthttp.RequestCtx) {
	penggalangDanaID := fmt.Sprintf("%s", ctx.UserValue("id"))
	responseData, err := domain.ToogleVerified(ctx, DB, penggalangDanaID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deletePenggalangDana(ctx *fasthttp.RequestCtx) {
	penggalangDanaID := fmt.Sprintf("%s", ctx.UserValue("id"))
	responseData, err := domain.DeletePenggalangDana(ctx, DB, penggalangDanaID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findPenggalangDanas(ctx *fasthttp.RequestCtx) {
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
	responseData, count, err := domain.FindPenggalangDana(ctx, DB, data)

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
	response := []ReadPenggalangDana{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp))
	}
	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getPenggalangDana(ctx *fasthttp.RequestCtx) {
	penggalangDanaID := fmt.Sprintf("%s", ctx.UserValue("id"))
	responseData, err := domain.GetPenggalangDana(ctx, DB, penggalangDanaID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
