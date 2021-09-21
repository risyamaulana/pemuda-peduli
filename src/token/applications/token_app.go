package applications

import (
	"errors"
	"fmt"
	"log"
	"os"
	"pemuda-peduli/src/common/handler"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/common/interfaces"
	"pemuda-peduli/src/common/utility"
	"pemuda-peduli/src/token/domain"
	"pemuda-peduli/src/token/infrastructure/repository"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var DB *db.ConnectTo

// TokenApp ...
type TokenApp struct {
	interfaces.IApplication
}

// NewTokenApp ...
func NewTokenApp(db *db.ConnectTo) *TokenApp {
	// Place where we init infrastructure, repo etc
	s := TokenApp{}
	DB = db
	return &s
}

// Initialize will be called when application run
func (s *TokenApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("Token app initialized")
}

// Route declaration
func (s *TokenApp) addRoute(r *router.Router) {
	// r.POST("/token/auth", middleware.CheckLoginToken(validateToken))
	r.POST("/token", getToken)
	r.POST("/refresh-token", refreshToken)
}

// ============== Handler for each route start here ============
func getToken(ctx *fasthttp.RequestCtx) {
	payload, err := GetPayload(ctx.Request.Body())
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Bad JSON Payload"))))
		log.Println("Error Bad Request JSON Payload:", err)
		return
	}
	ctx.SetUserValue("TOKEN_SECRET_KEY", payload.SecretKey)

	if err := payload.Validate(ctx); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	data := payload.ToEntity()

	if err := domain.GenerateToken(ctx, &data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	repo := repository.NewTokenRepository(DB)
	if err := domain.CreateOrUpdateToken(ctx, &repo, &data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func refreshToken(ctx *fasthttp.RequestCtx) {
	refreshToken := ctx.Request.Header.Peek("pp-refresh-token")
	if string(refreshToken) == "" {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Failed auth: Header pp_refresh_token is required"))))
		return
	}
	ctx.SetUserValue("TOKEN_SECRET_KEY", os.Getenv("TOKEN_SECRET_KEY"))

	data, err := domain.RefreshToken(ctx, string(refreshToken))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	repo := repository.NewTokenRepository(DB)
	if err := domain.CreateOrUpdateToken(ctx, &repo, &data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}
