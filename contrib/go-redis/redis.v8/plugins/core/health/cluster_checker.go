package health

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type ClusterClientChecker struct {
	client *redis.ClusterClient
}

func (c *ClusterClientChecker) Check(ctx context.Context) error {
	return c.client.Ping(ctx).Err()
}

func NewClusterClientChecker(client *redis.ClusterClient) *ClusterClientChecker {
	return &ClusterClientChecker{client: client}
}
