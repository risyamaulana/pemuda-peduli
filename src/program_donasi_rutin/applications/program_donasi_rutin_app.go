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
	"pemuda-peduli/src/program_donasi_rutin/domain"
	"pemuda-peduli/src/program_donasi_rutin/infrastructure/repository"
	"strconv"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var DB *db.ConnectTo

// ProgramDonasiRutinApp ...
type ProgramDonasiRutinApp struct {
	interfaces.IApplication
}

// NewProgramDonasiRutinApp ...
func NewProgramDonasiRutinApp(db *db.ConnectTo) *ProgramDonasiRutinApp {
	// Place where we init infrastructure, repo etc
	s := ProgramDonasiRutinApp{}
	DB = db
	return &s
}

// Initialize will be called when application run
func (s *ProgramDonasiRutinApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("ProgramDonasiRutin app initialized")
}

// Destroy will be called when app shutdowns
func (s *ProgramDonasiRutinApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("ProgramDonasiRutin app released...")
}

// Route declaration
func (s *ProgramDonasiRutinApp) addRoute(r *router.Router) {
	r.POST("/program-donasi-rutin/create", middleware.CheckAuthToken(DB, createProgramDonasiRutin))

	r.PUT("/program-donasi-rutin/{id}", middleware.CheckAuthToken(DB, updateProgramDonasiRutin))
	r.PUT("/program-donasi-rutin/publish/{id}", middleware.CheckAuthToken(DB, publishProgramDonasiRutin))
	r.PUT("/program-donasi-rutin/hide/{id}", middleware.CheckAuthToken(DB, hideProgramDonasiRutin))

	r.POST("/program-donasi-rutin/list", middleware.CheckAuthToken(DB, findProgramDonasiRutins))
	r.GET("/program-donasi-rutin/{id}", middleware.CheckAuthToken(DB, getProgramDonasiRutin))

	r.DELETE("/program-donasi-rutin/{id}", middleware.CheckAuthToken(DB, deleteProgramDonasiRutin))

	r.POST("/program-donasi-rutin/paket/create/{id}", middleware.CheckAuthToken(DB, createProgramDonasiRutinPaket))
	r.PUT("/program-donasi-rutin/paket/{id}", middleware.CheckAuthToken(DB, updateProgramDonasiRutinPaket))
	r.DELETE("/program-donasi-rutin/paket/{id}", middleware.CheckAuthToken(DB, deleteProgramDonasiRutinPaket))
	r.POST("/program-donasi-rutin/paket/list", middleware.CheckAuthToken(DB, findProgramDonasiRutinPaket))
	r.GET("/program-donasi-rutin/paket/{id}", middleware.CheckAuthToken(DB, getProgramDonasiRutinPaket))

	// Kabar terbaru
	r.POST("/program-donasi-rutin/kabar-terbaru/create", middleware.CheckAdminToken(DB, createKabarTerbaru))
	r.PUT("/program-donasi-rutin/kabar-terbaru/{id}", middleware.CheckAdminToken(DB, updateKabarTerbaru))
	r.POST("/program-donasi-rutin/kabar-terbaru/list", middleware.CheckAuthToken(DB, findKabarTerbaru))
	r.GET("/program-donasi-rutin/kabar-terbaru/{id}", middleware.CheckAuthToken(DB, getKabarTerbaru))
	r.DELETE("/program-donasi-rutin/kabar-terbaru/{id}", middleware.CheckAdminToken(DB, deleteKabarTerbaru))
}

// ============== Handler for each route start here ============

func createProgramDonasiRutin(ctx *fasthttp.RequestCtx) {
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
	responseData, err := domain.CreateProgramDonasiRutin(ctx, DB, &data)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func createProgramDonasiRutinPaket(ctx *fasthttp.RequestCtx) {
	payload, err := GetCreatePaketPayload(ctx.Request.Body())
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

	data := payload.ToEntity(ctx)
	err = domain.CreateProgramDonasiRutinPaket(ctx, DB, &data)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayloadPaket(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateProgramDonasiRutin(ctx *fasthttp.RequestCtx) {
	programDonasiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	payload, err := GetUpdatePayload(ctx.Request.Body())
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
	responseData, err := domain.EditProgramDonasiRutin(ctx, DB, data, programDonasiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateProgramDonasiRutinPaket(ctx *fasthttp.RequestCtx) {
	paketID := fmt.Sprintf("%s", ctx.UserValue("id"))
	payload, err := GetUpdatePaketPayload(ctx.Request.Body())
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
	responseData, err := domain.EditProgramDonasiRutinPaket(ctx, DB, data, paketID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayloadPaket(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func publishProgramDonasiRutin(ctx *fasthttp.RequestCtx) {
	programDonasiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewProgramDonasiRutinRepository(DB)
	responseData, err := domain.PublishProgramDonasiRutin(ctx, &repo, programDonasiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func hideProgramDonasiRutin(ctx *fasthttp.RequestCtx) {
	programDonasiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewProgramDonasiRutinRepository(DB)
	responseData, err := domain.HideProgramDonasiRutin(ctx, &repo, programDonasiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteProgramDonasiRutin(ctx *fasthttp.RequestCtx) {
	programDonasiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewProgramDonasiRutinRepository(DB)
	responseData, err := domain.DeleteProgramDonasiRutin(ctx, &repo, programDonasiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteProgramDonasiRutinPaket(ctx *fasthttp.RequestCtx) {
	paketID := fmt.Sprintf("%s", ctx.UserValue("id"))
	responseData, err := domain.DeleteProgramDonasiRutinPaket(ctx, DB, paketID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayloadPaket(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findProgramDonasiRutins(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewProgramDonasiRutinRepository(DB)

	responseData, count, err := domain.FindProgramDonasiRutin(ctx, &repo, &data)

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
	response := []ReadProgramDonasiRutin{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func findProgramDonasiRutinPaket(ctx *fasthttp.RequestCtx) {
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
	responseData, count, err := domain.FindProgramDonasiRutinPaket(ctx, DB, &data)

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
	response := []ReadProgramDonasiRutinPaket{}
	for _, resp := range responseData {
		response = append(response, ToPayloadPaket(resp))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getProgramDonasiRutin(ctx *fasthttp.RequestCtx) {
	programDonasiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewProgramDonasiRutinRepository(DB)
	responseData, err := domain.GetProgramDonasiRutin(ctx, &repo, programDonasiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func getProgramDonasiRutinPaket(ctx *fasthttp.RequestCtx) {
	programDonasiID := fmt.Sprintf("%s", ctx.UserValue("id"))
	responseData, err := domain.GetProgramDonasiRutinPaket(ctx, DB, programDonasiID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayloadPaket(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

// Kabar terbaru
func createKabarTerbaru(ctx *fasthttp.RequestCtx) {
	payload, err := GetCreateNewsPayload(ctx.Request.Body())
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
	err = domain.CreateProgramDonasiNews(ctx, DB, &data)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayloadNews(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateKabarTerbaru(ctx *fasthttp.RequestCtx) {
	kabarTerbaruID, err := strconv.ParseInt(ctx.UserValue("id").(string), 10, 64)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("id not valid"))))
		log.Println("Error id not valid:", err)
		return
	}
	payload, err := GetUpdateNewsPayload(ctx.Request.Body())
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
	responseData, err := domain.UpdateProgramDonasiNews(ctx, DB, data, kabarTerbaruID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayloadNews(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteKabarTerbaru(ctx *fasthttp.RequestCtx) {
	kabarTerbaruID, err := strconv.ParseInt(ctx.UserValue("id").(string), 10, 64)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("id not valid"))))
		log.Println("Error id not valid:", err)
		return
	}
	repo := repository.NewProgramDonasiRutinRepository(DB)
	responseData, err := domain.DeleteProgramDonasiNews(ctx, &repo, kabarTerbaruID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayloadNews(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findKabarTerbaru(ctx *fasthttp.RequestCtx) {
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

	responseData, count, err := domain.FindProgramDonasiNews(ctx, DB, &data)

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
	response := []ReadProgramDonasiNews{}
	for _, resp := range responseData {
		response = append(response, ToPayloadNews(resp))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getKabarTerbaru(ctx *fasthttp.RequestCtx) {
	kabarTerbaruID, err := strconv.ParseInt(ctx.UserValue("id").(string), 10, 64)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("id not valid"))))
		log.Println("Error id not valid:", err)
		return
	}
	responseData, err := domain.GetProgramDonasiNews(ctx, DB, kabarTerbaruID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayloadNews(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
