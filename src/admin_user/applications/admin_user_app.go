package applications

import (
	"errors"
	"fmt"
	"log"
	"math"
	"pemuda-peduli/src/admin_user/domain"
	"pemuda-peduli/src/admin_user/infrastructure/repository"
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

// AdminUserApp ...
type AdminUserApp struct {
	interfaces.IApplication
}

// NewAdminUserApp ...
func NewAdminUserApp(db *db.ConnectTo) *AdminUserApp {
	// Place where we init infrastructure, repo etc
	s := AdminUserApp{}
	DB = db

	return &s
}

// Initialize will be called when application run
func (s *AdminUserApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("AdminUser app initialized")
}

// Destroy will be called when app shutdowns
func (s *AdminUserApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("AdminUser app released...")
}

// Route declaration
func (s *AdminUserApp) addRoute(r *router.Router) {
	r.POST("/admin/create", middleware.CheckAdminToken(DB, createAdminUser))

	r.POST("/admin/check-username", middleware.CheckAdminToken(DB, checkAdminUsername))

	r.POST("/admin/list", middleware.CheckAdminToken(DB, findAdminUsers))
	r.GET("/admin/{id}", middleware.CheckAdminToken(DB, getAdminUser))
	r.GET("/admin", middleware.CheckAdminToken(DB, getAdminProfile))

	r.PUT("/admin", middleware.CheckAdminToken(DB, updateAdminUser))
	r.PUT("/admin/change-password", middleware.CheckAdminToken(DB, changePassword))
	r.PUT("/admin/reset-password", middleware.CheckAdminToken(DB, resetPassword))
	r.PUT("/admin/change-role", middleware.CheckAdminToken(DB, changeRole))

	r.DELETE("/admin/{id}", middleware.CheckAdminToken(DB, deleteAdminUser))
}

// ============== Handler for each route start here ============

func createAdminUser(ctx *fasthttp.RequestCtx) {
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
	if err := domain.CreateAdminUser(ctx, DB, &data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateAdminUser(ctx *fasthttp.RequestCtx) {
	userID := ctx.UserValue("user_id").(string)
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
	repo := repository.NewAdminUserRepository(DB)
	responseData, err := domain.UpdateAdminUser(ctx, &repo, data, userID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func changePassword(ctx *fasthttp.RequestCtx) {
	userID := ctx.UserValue("user_id").(string)
	payload, err := GetChangePasswordPayload(ctx.Request.Body())
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

	// Check old password
	userData, err := domain.GetAdminUser(ctx, DB, userID)
	if err != nil {
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
			fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
			log.Println(err)
			return
		}
	}

	oldPass := utility.GeneratePass(userData.Salt, payload.OldPassword)
	if oldPass != userData.Password {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Unauthorized data: old password is unauthorized"))))
		return
	}

	data := payload.ToEntity()
	repo := repository.NewAdminUserRepository(DB)
	if err := domain.UpdatePassword(ctx, &repo, data, userID); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(nil, nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

// TODO: Reset password
func resetPassword(ctx *fasthttp.RequestCtx) {
	payload, err := GetResetPasswordPayload(ctx.Request.Body())
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

	newPassword, err := domain.ResetPassword(ctx, DB, payload.ID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ResetPasswordResponse{
		NewPassword: newPassword,
	}, nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

// TODO: Change role
func changeRole(ctx *fasthttp.RequestCtx) {
	payload, err := GetChangeRolePayload(ctx.Request.Body())
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

	data, err := domain.ChangeRole(ctx, DB, payload.ID, payload.Role)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteAdminUser(ctx *fasthttp.RequestCtx) {
	adminID := fmt.Sprintf("%s", ctx.UserValue("id"))
	repo := repository.NewAdminUserRepository(DB)
	responseData, err := domain.DeleteAdminUser(ctx, &repo, adminID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func findAdminUsers(ctx *fasthttp.RequestCtx) {
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

	responseData, count, err := domain.FindAdminUser(ctx, DB, &data)

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
	response := []ReadAdminUser{}
	for _, resp := range responseData {
		response = append(response, ToPayload(resp))
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.PaginationResponse(response, nil, page, limit, int(pageTotal), count)))
}

func getAdminUser(ctx *fasthttp.RequestCtx) {
	adminID := fmt.Sprintf("%s", ctx.UserValue("id"))
	responseData, err := domain.GetAdminUser(ctx, DB, adminID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}
	response := handler.DefaultResponse(ToPayload(responseData), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func getAdminProfile(ctx *fasthttp.RequestCtx) {
	userID := ctx.UserValue("user_id").(string)
	data, err := domain.GetAdminUser(ctx, DB, userID)
	if err != nil {
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
			fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
			log.Println(err)
			return
		}
	}
	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func checkAdminUsername(ctx *fasthttp.RequestCtx) {
	payload, err := GetUsernameAdminPayload(ctx.Request.Body())
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

	_, err = domain.GetAdminUserByUsername(ctx, DB, payload.Username)
	if err == nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Username is already used"))))
		return
	} else {
		response := handler.DefaultResponse("Username available", nil)
		fmt.Fprintf(ctx, utility.PrettyPrint(response))
	}
}
