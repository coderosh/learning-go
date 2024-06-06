package main

import (
	"gshort/handlers"
	"gshort/middlewares"

	"github.com/labstack/echo/v4"
)

func Routes(app *echo.Echo) {

	pageRoutes(app)

	authRoutes(app)

	dashboardRoutes(app)

	urlRoutes(app)
}

func pageRoutes(app *echo.Echo) {
	pageHandler := handlers.PageHandler{}

	app.GET("/", pageHandler.HandleHomePage, middlewares.OptionalAuth)
	app.GET("/about/", pageHandler.HandleAboutPage, middlewares.OptionalAuth)
}

func authRoutes(app *echo.Echo) {
	authHandler := handlers.AuthHandler{}

	app.GET("/login/", authHandler.HandleLoginPage, middlewares.OptionalAuth)
	app.POST("/login/", authHandler.HandleLogin)

	app.GET("/signup/", authHandler.HandleSignupPage, middlewares.OptionalAuth)
	app.POST("/signup/", authHandler.HandleSignup)

	app.POST("/logout/", authHandler.HandleLogout)
}

func dashboardRoutes(app *echo.Echo) {
	dashboardHandler := handlers.DashboardHanlder{}

	dashboardRoutes := app.Group("/dashboard")
	dashboardRoutes.Use(middlewares.Authenticate)

	dashboardRoutes.GET("/", dashboardHandler.HandleDashboardPage)
	dashboardRoutes.GET("/analytics/", dashboardHandler.HandleAnalyticsPage)
}

func urlRoutes(app *echo.Echo) {
	urlHandler := handlers.UrlHandler{}

	urlRoutes := app.Group("/urls")
	urlRoutes.POST("/", urlHandler.HandleCreateUrl, middlewares.Authenticate)
	urlRoutes.GET("/:id/", urlHandler.HandleShortUrl, middlewares.Authenticate)
}
