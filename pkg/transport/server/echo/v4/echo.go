package echo

import (
	"context"
	"strconv"

	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	instance *echo.Echo
)

func Start(ctx context.Context) *echo.Echo {

	instance = echo.New()

	instance.HideBanner = GetHideBanner()
	instance.Logger = Wrap(log.GetLogger())

	setDefaultMiddlewares(ctx, instance)
	setDefaultRouters(ctx, instance)

	return instance
}

func setDefaultMiddlewares(ctx context.Context, instance *echo.Echo) {

	if GetMiddlewareLogEnabled() {
		instance.Use(Logger())
	}

	if GetMiddlewareRecoverEnabled() {
		instance.Use(middleware.Recover())
	}
}

func setDefaultRouters(ctx context.Context, instance *echo.Echo) {

	l := log.FromContext(ctx)

	statusRoute := GetStatusRoute()

	l.Infof("configuring status router on %s", statusRoute)

	statusHandler := NewResourceStatusHandler()
	instance.GET(statusRoute, statusHandler.Get)

	healthRoute := GetHealthRoute()

	l.Infof("configuring health router on %s", healthRoute)

	healthHandler := NewHealthHandler()
	instance.GET(healthRoute, healthHandler.Get)
}

func Serve(ctx context.Context) {
	l := log.FromContext(ctx)
	l.Infof("starting echo server. https://echo.labstack.com/")
	err := instance.Start(getServerPort())
	instance.Logger.Fatalf(err.Error())
}

func getServerPort() string {
	return ":" + strconv.Itoa(GetPort())
}
