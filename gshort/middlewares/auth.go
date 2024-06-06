package middlewares

import (
	"gshort/utils"

	"github.com/labstack/echo/v4"
)

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		tokenCookie, err := ctx.Cookie("session_token")
		if err != nil {
			return err
		}

		userId, err := utils.VerifyJWT(tokenCookie.Value)
		if err != nil {
			return err
		}

		ctx.Set("userId", userId)

		return next(ctx)
	}
}

func OptionalAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		tokenCookie, err := ctx.Cookie("session_token")

		if err == nil {
			userId, err := utils.VerifyJWT(tokenCookie.Value)
			if err == nil {
				ctx.Set("userId", userId)
			}
		}

		return next(ctx)
	}
}
