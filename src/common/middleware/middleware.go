package middleware

import (
	"errors"
	"fmt"

	"pemuda-peduli/src/common/handler"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/common/utility"

	tokenDom "pemuda-peduli/src/token/domain"

	adminUserDom "pemuda-peduli/src/admin_user/domain"

	userDom "pemuda-peduli/src/user/domain"
	userRep "pemuda-peduli/src/user/infrastructure/repository"

	roleDom "pemuda-peduli/src/role/domain"
	roleRepo "pemuda-peduli/src/role/infrastructure/repository"

	"github.com/valyala/fasthttp"
)

var (
	// authorization
	corsAllowHeaders     = "Origin, X-Request-With, Content-Type, Accept, pp-token, pp-refresh-token"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
	corsAllow            = "DELETE, GET, OPTIONS, POST, PUT"

	DB *db.ConnectTo
)

// db init hardcoded temporary for testing
func init() {
	DB = db.NewDBConnectionFactory(0)
}

func Cors(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)
		ctx.Response.Header.Set("Allow", corsAllow)
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.Response.Header.Set("Server", "pp-service")

		next(ctx)
	})
}

func CheckAuthToken(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		method := ctx.Request.Header.Method()
		if string(method) != "OPTIONS" {
			token := ctx.Request.Header.Peek("pp-token")
			if string(token) == "" {
				ctx.SetStatusCode(fasthttp.StatusUnauthorized)
				fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Failed auth: Header pp-token is required"))))
				return
			}
			// Check validation token
			err := tokenDom.Validate(ctx, string(token), DB)
			if err != nil {
				ctx.SetStatusCode(fasthttp.StatusUnauthorized)
				fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
				return
			}
		}
		next(ctx)
	})
}

func CheckLoginAdminToken(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		method := ctx.Request.Header.Method()
		if string(method) != "OPTIONS" {
			token := ctx.Request.Header.Peek("pp-token")
			if string(token) == "" {
				ctx.SetStatusCode(fasthttp.StatusUnauthorized)
				fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Failed auth: Header pp-token is required"))))
				return
			}
			// Check validation token
			err := tokenDom.ValidateAdminLogin(ctx, string(token), DB)
			if err != nil {
				ctx.SetStatusCode(fasthttp.StatusUnauthorized)
				fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
				return
			}
		}
		next(ctx)
	})
}

func CheckAdminToken(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		method := ctx.Request.Header.Method()
		if string(method) != "OPTIONS" {
			token := ctx.Request.Header.Peek("pp-token")
			if string(token) == "" {
				ctx.SetStatusCode(fasthttp.StatusUnauthorized)
				fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Failed auth: Header pp-token is required"))))
				return
			}
			// Check validation token
			err := tokenDom.ValidateAdminLogin(ctx, string(token), DB)
			if err != nil {
				ctx.SetStatusCode(fasthttp.StatusUnauthorized)
				fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Failed, user is unauthorized"))))
				return
			}

			// Get User Data
			userID := ctx.UserValue("user_id").(string)
			dataUser, err := adminUserDom.GetAdminUser(ctx, DB, userID)
			if err != nil {
				ctx.SetStatusCode(fasthttp.StatusUnauthorized)
				fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Failed, user is unauthorized"))))
				return
			}

			// Get Role Data
			roleRepository := roleRepo.NewRoleRepository(DB)
			roleData, err := roleDom.GetRole(ctx, &roleRepository, dataUser.Role)
			if err != nil {
				ctx.SetStatusCode(fasthttp.StatusUnauthorized)
				fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
				return
			}
			ctx.SetUserValue("user_role_level", roleData.RoleLevel)
		}
		next(ctx)
	})
}

func CheckLoginuserToken(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		method := ctx.Request.Header.Method()
		if string(method) != "OPTIONS" {
			token := ctx.Request.Header.Peek("pp-token")
			if string(token) == "" {
				ctx.SetStatusCode(fasthttp.StatusUnauthorized)
				fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Failed auth: Header pp-token is required"))))
				return
			}
			// Check validation token
			err := tokenDom.ValidateUserLogin(ctx, string(token), DB)
			if err != nil {
				ctx.SetStatusCode(fasthttp.StatusUnauthorized)
				fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
				return
			}
		}
		next(ctx)
	})
}

func CheckUserToken(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		method := ctx.Request.Header.Method()
		if string(method) != "OPTIONS" {
			token := ctx.Request.Header.Peek("pp-token")
			if string(token) == "" {
				ctx.SetStatusCode(fasthttp.StatusUnauthorized)
				fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Failed auth: Header pp-token is required"))))
				return
			}
			// Check validation token
			err := tokenDom.ValidateUserLogin(ctx, string(token), DB)
			if err != nil {
				ctx.SetStatusCode(fasthttp.StatusUnauthorized)
				fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
				return
			}

			// Get User Data
			userID := ctx.UserValue("user_id").(string)
			repo := userRep.NewUserRepository(DB)
			_, err = userDom.ReadUser(ctx, &repo, userID)
			if err != nil {
				ctx.SetStatusCode(fasthttp.StatusUnauthorized)
				fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Failed, user is unauthorized, user not found"))))
				return
			}
		}
		next(ctx)
	})
}
