package handlers

import (
	"errors"
	"gshort/models"
	"gshort/views/dashboard"
	"math/rand"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type UrlHandler struct {
}

func (uh *UrlHandler) HandleCreateUrl(ctx echo.Context) error {
	var url models.Url
	err := ctx.Bind(&url)
	if err != nil {
		return err
	}

	url.ShortCode = createShortCode()
	url.DateTime = time.Now()
	curId, ok := ctx.Get("userId").(int64)

	if !ok {
		return errors.New("something went wrong")
	}

	url.UserID = curId

	err = url.Save()
	if err != nil {
		return err
	}

	return Render(ctx, http.StatusCreated, dashboard.UrlItem(templ.URL(url.LongUrl), templ.URL(url.ShortCode)))
}

func (uh *UrlHandler) HandleShortUrl(ctx echo.Context) error {
	shortCode := ctx.Param("id")
	url, err := models.FindUrlByShortCode(shortCode)
	if err != nil {
		return err
	}

	return ctx.Redirect(http.StatusFound, url.LongUrl)
}

func createShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}
