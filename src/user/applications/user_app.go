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
	"pemuda-peduli/src/user/domain"
	"pemuda-peduli/src/user/infrastructure/repository"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var DB *db.ConnectTo

// db init hardcoded temporary for testing
func init() {
	DB = db.NewDBConnectionFactory(0)
}

// UserApp ...
type UserApp struct {
	interfaces.IApplication
}

// NewUserApp ...
func NewUserApp() *UserApp {
	// Place where we init infrastructure, repo etc
	s := UserApp{}
	return &s
}

// Initialize will be called when application run
func (s *UserApp) Initialize(r *router.Router) {
	s.addRoute(r)
	log.Println("User app initialized")
}

// Destroy will be called when app shutdowns
func (s *UserApp) Destroy() {
	// TODO Do clean up resource here
	log.Println("User app released...")
}

// Route declaration
func (s *UserApp) addRoute(r *router.Router) {
	r.POST("/user/register", middleware.CheckAuthToken(registerUserHandler))

	r.PUT("/user", middleware.CheckUserToken(updateUserHandler))
	r.PUT("/user/change-password", middleware.CheckUserToken(changePasswordHandler))

	r.POST("/user/forgot-password", middleware.CheckUserToken(forgotPasswordHandler))
	r.PUT("/user/reset-password", middleware.CheckAuthToken(resetPasswordHandler))

	r.DELETE("/user", middleware.CheckUserToken(deleteUser))

	r.GET("/user", middleware.CheckUserToken(getUserProfile))
}

// ============== Handler for each route start here ============

func registerUserHandler(ctx *fasthttp.RequestCtx) {
	payload, err := GetRegisterUserPayload(ctx.Request.Body())
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
	repo := repository.NewUserRepository(DB)
	if err := domain.RegisterUser(ctx, &repo, &data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func updateUserHandler(ctx *fasthttp.RequestCtx) {
	payload, err := GetUpdateUserPayload(ctx.Request.Body())
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

	repo := repository.NewUserRepository(DB)
	data, err := domain.UpdateUser(ctx, &repo, payload.ToEntity(ctx))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(ToPayload(data), nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func changePasswordHandler(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewUserRepository(DB)
	err = checkOldPassword(ctx, repo, userID, payload)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		log.Println(err)
		return
	}

	// Update password
	data := payload.ToEntity(ctx)
	if _, err := domain.ChangePassword(ctx, &repo, data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	response := handler.DefaultResponse(nil, nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func forgotPasswordHandler(ctx *fasthttp.RequestCtx) {
	payload, err := GetForgotPasswordPayload(ctx.Request.Body())
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

	repo := repository.NewUserRepository(DB)
	_, err = domain.ForgotPassword(ctx, &repo, payload.Email)
	if err != nil {
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
			fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
			log.Println(err)
			return
		}
	}
	response := handler.DefaultResponse(nil, nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func resetPasswordHandler(ctx *fasthttp.RequestCtx) {
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

	repo := repository.NewUserRepository(DB)
	_, err = domain.ResetPassword(ctx, &repo, payload.Token, payload.NewPassword)
	if err != nil {
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
			fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
			log.Println(err)
			return
		}
	}
	response := handler.DefaultResponse(nil, nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func deleteUser(ctx *fasthttp.RequestCtx) {
	userID := ctx.UserValue("user_id").(string)
	repo := repository.NewUserRepository(DB)
	_, err := domain.RemoveDeleteUser(ctx, &repo, userID)
	if err != nil {
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
			fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
			log.Println(err)
			return
		}
	}
	response := handler.DefaultResponse(nil, nil)
	fmt.Fprintf(ctx, utility.PrettyPrint(response))
}

func getUserProfile(ctx *fasthttp.RequestCtx) {
	userID := ctx.UserValue("user_id").(string)
	repo := repository.NewUserRepository(DB)
	data, err := domain.ReadUser(ctx, &repo, userID)
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

// ============== Local Function ============
func checkOldPassword(ctx *fasthttp.RequestCtx, repo repository.UserRepository, userID string, payload ChangePassword) (err error) {
	userData, err := domain.ReadUser(ctx, &repo, userID)
	if err != nil {
		return
	}

	oldPass := utility.GeneratePass(userData.Salt, payload.OldPassword)
	if oldPass != userData.Password {
		err = errors.New("Unauthorized data: old password is unauthorized")
		return
	}
	return
}
