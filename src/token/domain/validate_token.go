package domain

import (
	"errors"
	"fmt"
	"log"
	"os"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/token/infrastructure/repository"

	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
)

func Validate(ctx *fasthttp.RequestCtx, token string, db *db.ConnectTo) (err error) {
	authToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET_KEY")), nil
	})
	if authToken != nil {
		if !authToken.Valid || err != nil {
			log.Println(err)
			err = errors.New("Failed auth : Please provide correct token!")
		} else {
			if claims, ok := authToken.Claims.(jwt.MapClaims); ok && authToken.Valid {
				repo := repository.NewTokenRepository(db)
				data, errCheckDevice := repo.CheckTokenDevice(token, claims["device_id"].(string), claims["device_type"].(string))
				if errCheckDevice != nil {
					err = errors.New("Failed auth : Please provide correct token!")
					return

				}
				// Get the user record from database or
				ctx.SetUserValue("device_id", claims["device_id"].(string))
				ctx.SetUserValue("device_type", claims["device_type"].(string))
				ctx.SetUserValue("token", token)
				ctx.SetUserValue("user_id", data.LoginID)
			}
			return
		}
	} else {
		err = errors.New("Failed auth : Please provide correct token!")
	}

	return
}

func ValidateAdminLogin(ctx *fasthttp.RequestCtx, token string, db *db.ConnectTo) (err error) {
	authToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET_KEY")), nil
	})
	if authToken != nil {
		if !authToken.Valid || err != nil {
			log.Println(err)
			err = errors.New("Failed auth : Please provide correct token!")
			return
		} else {
			if claims, ok := authToken.Claims.(jwt.MapClaims); ok && authToken.Valid {
				repo := repository.NewTokenRepository(db)
				data, errCheckDevice := repo.CheckTokenDevice(token, claims["device_id"].(string), claims["device_type"].(string))
				if errCheckDevice != nil {
					err = errors.New("Failed auth : Please provide correct token!")
					return

				}
				if !data.IsLogin {
					err = errors.New("Failed auth : Please login first.")
					return
				}
				// Get the user record from database or
				ctx.SetUserValue("device_id", claims["device_id"].(string))
				ctx.SetUserValue("device_type", claims["device_type"].(string))
				ctx.SetUserValue("token", token)
				ctx.SetUserValue("user_id", data.LoginID)
			}
			return
		}
	} else {
		err = errors.New("Failed auth : Please provide correct token!")
		return
	}

}

func ValidateUserLogin(ctx *fasthttp.RequestCtx, token string, db *db.ConnectTo) (err error) {
	authToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET_KEY")), nil
	})
	if authToken != nil {
		if !authToken.Valid || err != nil {
			log.Println(err)
			err = errors.New("Failed auth : Please provide correct token!")
			return
		} else {
			if claims, ok := authToken.Claims.(jwt.MapClaims); ok && authToken.Valid {
				repo := repository.NewTokenRepository(db)
				data, errCheckDevice := repo.CheckTokenDevice(token, claims["device_id"].(string), claims["device_type"].(string))
				if errCheckDevice != nil {
					err = errors.New("Failed auth : Please provide correct token!")
					return

				}
				if !data.IsLogin {
					err = errors.New("Failed auth : Please login first.")
					return
				}
				// Get the user record from database or
				ctx.SetUserValue("device_id", claims["device_id"].(string))
				ctx.SetUserValue("device_type", claims["device_type"].(string))
				ctx.SetUserValue("token", token)
				ctx.SetUserValue("user_id", data.LoginID)
			}
			return
		}
	} else {
		err = errors.New("Failed auth : Please provide correct token!")
		return
	}

}
