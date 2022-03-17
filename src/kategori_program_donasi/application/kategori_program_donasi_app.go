package application

import (
	"errors"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
	"math"
	"pemuda-peduli/src/common/handler"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/common/interfaces"
	"pemuda-peduli/src/common/middleware"
	"pemuda-peduli/src/common/utility"
	"pemuda-peduli/src/kategori_program_donasi/domain"
	"pemuda-peduli/src/kategori_program_donasi/infrastructure/repository"
	"strconv"
)

var DB *db.ConnectTo

// KategoriProgramDonasiApp ...
type KategoriProgramDonasiApp struct {
	interfaces.IApplication
}

// NewKategoriProgramDonasiApp ...
func NewKategoriProgramDonasiApp(db *db.ConnectTo) *KategoriProgramDonasiApp {
	// Place where we init infrastructure, repo etc
	s := KategoriProgramDonasiApp{}
	DB = db
	return &s
}

// Initialize will be called when application run
func (s *KategoriProgramDonasiApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("Kategori Program Donasi app initialized")
}

// Destroy will be called when app shutdowns
func (s *KategoriProgramDonasiApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("Kategori Program Donasi app released...")
}

// Route declaration
func (s *KategoriProgramDonasiApp) addRoute(r *router.Router) {
	r.POST("/kategori-program-donasi/create", middleware.CheckAuthToken(DB, createKategoriProgramDonasi))

	r.PUT("/kategori-program-donasi/{id}", middleware.CheckAuthToken(DB, updateKategoriProgramDonasi))

	r.POST("/kategori-program-donasi/list", middleware.CheckAuthToken(DB, findKategoriProgramDonasis))
	r.GET("/kategori-program-donasi/{id}", middleware.CheckAuthToken(DB, getKategoriProgramDonasi))

	r.DELETE("/kategori-program-donasi/{id}", middleware.CheckAuthToken(DB, deleteKategoriProgramDonasi))
}

// ============== Handler for each route start here ============

func createKategoriProgramDonasi(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewKategoriProgramDonasiRepository(DB)
	if err := domain.CreateKategoriProgramDonasi(ctx, &repo, &data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateKategoriProgramDonasi(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewKategoriProgramDonasiRepository(DB)
	responseData, err := domain.UpdateKategoriProgramDonasi(ctx, &repo, data, kontakKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteKategoriProgramDonasi(ctx *fasthttp.RequestCtx) {
	kontakKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewKategoriProgramDonasiRepository(DB)
	responseData, err := domain.DeleteKategoriProgramDonasi(ctx, &repo, kontakKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findKategoriProgramDonasis(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewKategoriProgramDonasiRepository(DB)

	responseData, count, err := domain.FindKategoriProgramDonasi(ctx, &repo, &data)

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
	response := []ReadKategoriProgramDonasi{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getKategoriProgramDonasi(ctx *fasthttp.RequestCtx) {
	kontakKamiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewKategoriProgramDonasiRepository(DB)
	responseData, err := domain.GetKategoriProgramDonasi(ctx, &repo, kontakKamiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
