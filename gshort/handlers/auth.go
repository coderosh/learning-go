package handlers

import (
	"gshort/models"
	"gshort/utils"
	"gshort/views/auth"
	"maps"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
}

func (uh *AuthHandler) HandleLoginPage(ctx echo.Context) error {
	tplData := getBasicTemplateData(ctx)
	defaultData := getDefaultAuthFormData("", "", "")

	maps.Copy(tplData, defaultData)

	return Render(ctx, 200, auth.Login(tplData))
}

func (uh *AuthHandler) HandleSignupPage(ctx echo.Context) error {
	tplData := getBasicTemplateData(ctx)
	defaultData := getDefaultAuthFormData("", "", "")

	maps.Copy(tplData, defaultData)

	return Render(ctx, 200, auth.Signup(tplData))
}

func (uh *AuthHandler) HandleLogin(ctx echo.Context) error {
	var user models.User

	err := ctx.Bind(&user)
	if err != nil {
		return err
	}

	err = user.VerifyEmailAndPassword()
	if err != nil {
		if err.Error() == "invalid credentials" {
			tplData := getBasicTemplateData(ctx)
			defaultData := getDefaultAuthFormData(user.Email, user.Password, err.Error())

			maps.Copy(tplData, defaultData)

			return Render(ctx, http.StatusUnauthorized, auth.Login(tplData))
		}

		return err
	}

	token, exp, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return err
	}

	sessionCookie := http.Cookie{
		Name:     "session_token",
		Value:    token,
		HttpOnly: true,
		Secure:   true,
		Expires:  exp,
	}

	ctx.SetCookie(&sessionCookie)

	return ctx.Redirect(http.StatusSeeOther, "/dashboard")
}

func (uh *AuthHandler) HandleSignup(ctx echo.Context) error {
	var user models.User
	ctx.Bind(&user)

	err := user.Save()
	if err != nil {
		return err
	}

	return ctx.Redirect(http.StatusSeeOther, "/login")
}

func (uh *AuthHandler) HandleLogout(ctx echo.Context) error {
	sessionCookie := http.Cookie{
		Name:     "session_token",
		Value:    "",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
		Secure:   true,
	}

	ctx.SetCookie(&sessionCookie)

	return ctx.Redirect(http.StatusSeeOther, "/login")
}

func getDefaultAuthFormData(email string, password string, err string) map[string]string {
	return map[string]string{
		"email":    email,
		"password": password,
		"error":    err,
	}
}
