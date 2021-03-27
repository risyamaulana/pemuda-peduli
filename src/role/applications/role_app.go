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
	"pemuda-peduli/src/role/domain"
	"pemuda-peduli/src/role/infrastructure/repository"
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

// RoleApp ...
type RoleApp struct {
	interfaces.IApplication
}

// NewRoleApp ...
func NewRoleApp() *RoleApp {
	// Place where we init infrastructure, repo etc
	s := RoleApp{}
	return &s
}

// Initialize will be called when application run
func (s *RoleApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("Role app initialized")
}

// Destroy will be called when app shutdowns
func (s *RoleApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("Role app released...")
}

// Route declaration
func (s *RoleApp) addRoute(r *router.Router) {
	r.POST("/role/list", middleware.CheckAdminToken(findRoles))
	r.GET("/role/{id}", middleware.CheckAdminToken(getRole))
}

// ============== Handler for each route start here ============
func findRoles(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewRoleRepository(DB)

	responseData, count, err := domain.FindRole(ctx, &repo, &data)

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
	response := []ReadRole{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getRole(ctx *fasthttp.RequestCtx) {
	roleID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewRoleRepository(DB)
	responseData, err := domain.GetRole(ctx, &repo, roleID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
