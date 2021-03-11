package giechobodylimit

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

	logger.Trace("enabling body limit middleware in echo")

	instance.Use(middleware.BodyLimit(GetSize()))

	logger.Debug("body limit middleware successfully enabled in echo")

	return nil
}
