package handlers

import (
	"gshort/views/page"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PageHandler struct{}

func (ph *PageHandler) HandleHomePage(ctx echo.Context) error {
	return Render(ctx, http.StatusOK, page.Home(getBasicTemplateData(ctx)))
}

func (ph *PageHandler) HandleAboutPage(ctx echo.Context) error {
	return Render(ctx, http.StatusOK, page.About(getBasicTemplateData(ctx)))
}
