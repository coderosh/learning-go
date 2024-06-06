package handlers

import (
	"errors"
	"net/http"

	"gshort/models"
	"gshort/views/dashboard"

	"github.com/labstack/echo/v4"
)

type DashboardHanlder struct {
}

func (dh *DashboardHanlder) HandleDashboardPage(ctx echo.Context) error {
	userId, ok := ctx.Get("userId").(int64)
	if !ok {
		return errors.New("something went wrong")
	}

	urls, err := models.FindAllUrlsByUserId(userId)
	if err != nil {
		return err
	}

	return Render(ctx, http.StatusOK, dashboard.Dashboard(getBasicTemplateData(ctx), urls))

}

func (dh *DashboardHanlder) HandleAnalyticsPage(ctx echo.Context) error {
	return Render(ctx, http.StatusOK, dashboard.Analytics(getBasicTemplateData(ctx)))
}
