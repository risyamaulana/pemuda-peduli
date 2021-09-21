package applications

import (
	"errors"
	"fmt"
	"log"
	"pemuda-peduli/src/common/handler"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/common/interfaces"
	"pemuda-peduli/src/common/middleware"
	"pemuda-peduli/src/common/utility"

	adminApps "pemuda-peduli/src/admin_user/applications"
	adminUserDom "pemuda-peduli/src/admin_user/domain"

	tokenDom "pemuda-peduli/src/token/domain"
	tokenMod "pemuda-peduli/src/token/domain/entity"
	tokenRepo "pemuda-peduli/src/token/infrastructure/repository"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var DB *db.ConnectTo

// AuthAdminApp ...
type AuthAdminApp struct {
	interfaces.IApplication
}

// NewAuthAdminApp ...
func NewAuthAdminApp(db *db.ConnectTo) *AuthAdminApp {
	// Place where we init infrastructure, repo etc
	s := AuthAdminApp{}
	DB = db
	return &s
}

// Initialize will be called when application run
func (s *AuthAdminApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("AuthAdmin app initialized")
}

// Destroy will be called when app shutdowns
func (s *AuthAdminApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("AuthAdmin app released...")
}

// Route declaration
func (s *AuthAdminApp) addRoute(r *router.Router) {
	r.POST("/auth/admin/login", middleware.CheckAuthToken(DB, login))

	r.POST("/auth/admin/logout", middleware.CheckLoginAdminToken(DB, logout))
}

// ============== Handler for each route start here ============
func login(ctx *fasthttp.RequestCtx) {
	payload, err := GetLoginPayload(ctx.Request.Body())
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Bad JSON Payload"))))
		log.Println("Error Bad Request JSON Payload:", err)
		return
	}

	if err := payload.Validate(ctx); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	data, err := adminUserDom.GetAdminUserByUsername(ctx, DB, payload.Username)
	if err != nil {
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
			fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
			log.Println(err)
			return
		}
	}

	// Check Password
	password := utility.GeneratePass(data.Salt, payload.Password)
	if data.Password != password {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Password not match"))))
		return
	}

	tokenData := tokenMod.TokenEntity{
		DeviceID:   ctx.UserValue("device_id").(string),
		DeviceType: ctx.UserValue("device_type").(string),
		Token:      ctx.UserValue("token").(string),
		IsLogin:    true,
		LoginID:    data.IDPPCPAdminUser,
	}

	tokenRepository := tokenRepo.NewTokenRepository(DB)
	if err = tokenDom.UpdateToken(ctx, &tokenRepository, &tokenData); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(adminApps.ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func logout(ctx *fasthttp.RequestCtx) {
	tokenData := tokenMod.TokenEntity{
		DeviceID:   ctx.UserValue("device_id").(string),
		DeviceType: ctx.UserValue("device_type").(string),
		Token:      ctx.UserValue("token").(string),
		IsLogin:    false,
		LoginID:    "",
	}
	tokenRepository := tokenRepo.NewTokenRepository(DB)
	if err := tokenDom.UpdateToken(ctx, &tokenRepository, &tokenData); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(nil, nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
