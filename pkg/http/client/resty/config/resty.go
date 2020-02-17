package config

import (
	"github.com/jpfaria/goignite/pkg/config"

	"log"
)

const Debug = "resty.debug"
const RequestTimeout = "resty.request.timeout"
const RetryCount = "resty.retry.count"
const RetryWaitTime = "resty.retry.waittime"
const RetryMaxWaitTime = "resty.retry.maxwaittime"
const HealthEnabled = "resty.health.enabled"
const HealthDescription = "resty.health.description"
const HealthRequired = "resty.health.required"


func init() {
	log.Println("getting configurations for resty")

	config.Add(Debug, false, "defines global debug request")
	config.Add(RequestTimeout, 2000, "defines global http request timeout (ms)")
	config.Add(RetryCount, 0, "defines global max http retries")
	config.Add(RetryWaitTime, 200, "defines global retry wait time (ms)")
	config.Add(RetryMaxWaitTime, 2000, "defines global max retry wait time (ms)")
	config.Add(HealthEnabled, true, "enabled/disable health check")
	config.Add(HealthDescription, "default connection", "define health description")
	config.Add(HealthRequired, false, "define health description")


}
