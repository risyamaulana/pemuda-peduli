package domain

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"pemuda-peduli/src/token/domain/entity"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(ctx context.Context, data *entity.TokenEntity) (err error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// This is the information which frontend can use
	expiredToken := time.Now().Add(time.Hour * 1)
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = data.Name
	claims["device_id"] = data.DeviceID
	claims["device_type"] = data.DeviceType
	claims["created_at"] = data.CreatedAt
	claims["exp"] = expiredToken.Unix()

	// Generate encoded token and send it as response.
	expiredRefreshToken := time.Now().Add(time.Hour * 24)
	// The signing string should be secret (a generated UUID works too)
	t, err := token.SignedString([]byte(ctx.Value("TOKEN_SECRET_KEY").(string)))
	if err != nil {
		return err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["name"] = data.Name
	rtClaims["device_id"] = data.DeviceID
	rtClaims["device_type"] = data.DeviceType
	rtClaims["created_at"] = data.CreatedAt
	rtClaims["exp"] = expiredRefreshToken.Unix()

	rt, err := refreshToken.SignedString([]byte(ctx.Value("TOKEN_SECRET_KEY").(string)))
	if err != nil {
		return err
	}

	data.Token = t
	data.TokenExpired = expiredToken
	data.RefreshToken = rt
	data.RefreshTokenExpired = expiredRefreshToken

	return
}

func RefreshToken(ctx context.Context, refreshToken string) (response entity.TokenEntity, err error) {
	data := entity.TokenEntity{}
	authToken, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
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
				// Get the user record from database or
				data.Name = claims["name"].(string)
				data.DeviceID = claims["device_id"].(string)
				data.DeviceType = claims["device_type"].(string)
				data.CreatedAt = time.Now()

				if err = GenerateToken(ctx, &data); err != nil {
					return
				}
				response = data
				return
			}
			return
		}
	} else {
		err = errors.New("Failed auth : Please provide correct token!")
		return
	}
}
