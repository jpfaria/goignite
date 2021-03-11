package giechocors

import (
	"context"

	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Register(ctx context.Context, instance *echo.Echo) error {
	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("enabling cors middleware in echo")

	instance.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     GetAllowOrigins(),
		AllowMethods:     GetAllowMethods(),
		AllowHeaders:     GetAllowHeaders(),
		AllowCredentials: GetAllowCredentials(),
		ExposeHeaders:    GetExposeHeaders(),
		MaxAge:           GetMaxAge(),
	}))

	logger.Debug("cors middleware successfully enabled in echo")

	return nil
}
