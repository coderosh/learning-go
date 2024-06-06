package handlers

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func getUserIDFromCtx(ctx echo.Context) string {
	userId, ok := ctx.Get("userId").(int64)

	if !ok {
		return ""
	}

	return fmt.Sprintf("%d", userId)
}

func getBasicTemplateData(ctx echo.Context) map[string]string {
	return map[string]string{
		"userId": getUserIDFromCtx(ctx),
		"path":   ctx.Request().URL.Path,
	}
}
