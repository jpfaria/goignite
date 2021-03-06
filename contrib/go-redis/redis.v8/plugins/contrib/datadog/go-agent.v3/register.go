package newrelic

import (
	"context"

	datadog "github.com/b2wdigital/goignite/v2/contrib/datadog/dd-trace-go.v1"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/go-redis/redis/v8"
	redistrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis.v8"
)

func Register(ctx context.Context, client *redis.Client) error {

	if !IsEnabled() || !datadog.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating redis in datadog")

	redistrace.WrapClient(client)

	logger.Debug("redis successfully integrated in datadog")

	return nil
}
