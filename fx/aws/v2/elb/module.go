package gifxawselb

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	gifxaws "github.com/b2wdigital/goignite/fx/aws/v2"
	"go.uber.org/fx"
)

var once sync.Once

func Module() fx.Option {

	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			gifxaws.AWSModule(),
			fx.Provide(
				elasticloadbalancingv2.NewFromConfig,
			),
		)
	})

	return options
}
