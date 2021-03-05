package giecho

import (
	"context"
	"strconv"

	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/labstack/echo/v4"
)

var (
	instance *echo.Echo
)

type Ext func(context.Context, *echo.Echo) error

func New(ctx context.Context, exts ...Ext) *echo.Echo {

	instance = echo.New()

	instance.HideBanner = GetHideBanner()
	instance.Logger = WrapLogger(gilog.GetLogger())

	for _, ext := range exts {
		if err := ext(ctx, instance); err != nil {
			panic(err)
		}
	}

	return instance
}

func Serve(ctx context.Context) {
	logger := gilog.FromContext(ctx)
	logger.Infof("starting echo server. https://echo.labstack.com/")
	if err := instance.Start(serverPort()); err != nil {
		instance.Logger.Fatalf(err.Error())
	}
}

func serverPort() string {
	return ":" + strconv.Itoa(GetPort())
}
