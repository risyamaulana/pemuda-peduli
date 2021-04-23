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
	"pemuda-peduli/src/menu_extras/domain"
	"pemuda-peduli/src/menu_extras/infrastructure/repository"
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

// MenuExtrasApp ...
type MenuExtrasApp struct {
	interfaces.IApplication
}

// NewMenuExtrasApp ...
func NewMenuExtrasApp() *MenuExtrasApp {
	// Place where we init infrastructure, repo etc
	s := MenuExtrasApp{}
	return &s
}

// Initialize will be called when application run
func (s *MenuExtrasApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("MenuExtras app initialized")
}

// Destroy will be called when app shutdowns
func (s *MenuExtrasApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("MenuExtras app released...")
}

// Route declaration
func (s *MenuExtrasApp) addRoute(r *router.Router) {
	r.POST("/menu-extras/create", middleware.CheckAuthToken(createMenuExtras))

	r.PUT("/menu-extras/{id}", middleware.CheckAuthToken(updateMenuExtras))
	r.PUT("/menu-extras/publish/{id}", middleware.CheckAuthToken(publishMenuExtras))
	r.PUT("/menu-extras/hide/{id}", middleware.CheckAuthToken(hideMenuExtras))

	r.POST("/menu-extras/list", middleware.CheckAuthToken(findMenuExtrass))
	r.GET("/menu-extras/{id}", middleware.CheckAuthToken(getMenuExtras))

	r.DELETE("/menu-extras/{id}", middleware.CheckAuthToken(deleteMenuExtras))
}

// ============== Handler for each route start here ============

func createMenuExtras(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewMenuExtrasRepository(DB)
	if err := domain.CreateMenuExtras(ctx, &repo, &data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateMenuExtras(ctx *fasthttp.RequestCtx) {
	menuExtrasID := fmt.Sprintf("%s", ctx.UserValue("id"))
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
	repo := repository.NewMenuExtrasRepository(DB)
	responseData, err := domain.UpdateMenuExtras(ctx, &repo, data, menuExtrasID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func publishMenuExtras(ctx *fasthttp.RequestCtx) {
	menuExtrasID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewMenuExtrasRepository(DB)
	responseData, err := domain.PublishMenuExtras(ctx, &repo, menuExtrasID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func hideMenuExtras(ctx *fasthttp.RequestCtx) {
	menuExtrasID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewMenuExtrasRepository(DB)
	responseData, err := domain.HideMenuExtras(ctx, &repo, menuExtrasID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteMenuExtras(ctx *fasthttp.RequestCtx) {
	menuExtrasID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewMenuExtrasRepository(DB)
	responseData, err := domain.DeleteMenuExtras(ctx, &repo, menuExtrasID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findMenuExtrass(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewMenuExtrasRepository(DB)

	responseData, count, err := domain.FindMenuExtras(ctx, &repo, &data)

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
	response := []ReadMenuExtras{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getMenuExtras(ctx *fasthttp.RequestCtx) {
	menuExtrasID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewMenuExtrasRepository(DB)
	responseData, err := domain.GetMenuExtras(ctx, &repo, menuExtrasID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
