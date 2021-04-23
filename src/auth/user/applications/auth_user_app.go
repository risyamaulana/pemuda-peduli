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

	tokenDom "pemuda-peduli/src/token/domain"
	tokenMod "pemuda-peduli/src/token/domain/entity"
	tokenRep "pemuda-peduli/src/token/infrastructure/repository"

	userApp "pemuda-peduli/src/user/applications"
	userDom "pemuda-peduli/src/user/domain"
	userRep "pemuda-peduli/src/user/infrastructure/repository"

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

// AuthUserApp ...
type AuthUserApp struct {
	interfaces.IApplication
}

// NewAuthUserApp ...
func NewAuthUserApp() *AuthUserApp {
	// Place where we init infrastructure, repo etc
	s := AuthUserApp{}
	return &s
}

// Initialize will be called when application run
func (s *AuthUserApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("AuthUser app initialized")
}

// Destroy will be called when app shutdowns
func (s *AuthUserApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("AuthUser app released...")
}

// Route declaration
func (s *AuthUserApp) addRoute(r *router.Router) {
	r.POST("/auth/user/login", middleware.CheckAuthToken(login))

	r.POST("/auth/user/logout", middleware.CheckLoginuserToken(logout))
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
	repo := userRep.NewUserRepository(DB)
	data, err := userDom.ReadLoginuser(ctx, &repo, payload.Username)
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
		LoginID:    data.IDUser,
	}

	tokenRepository := tokenRep.NewTokenRepository(DB)
	if err = tokenDom.UpdateToken(ctx, &tokenRepository, &tokenData); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(userApp.ToPayload(data), nil)
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
	tokenRepository := tokenRep.NewTokenRepository(DB)
	if err := tokenDom.UpdateToken(ctx, &tokenRepository, &tokenData); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(nil, nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
