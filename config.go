package idempotency

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

// IdempotencyConfig represents all available options for the middleware
type IdempotencyConfig struct {
	// HeaderName is to be passed when use want to use a custom key name for your
	// requestid header. By default it looks for the key name X-REQUEST-ID
	HeaderName string

	// WhitelistHTTPMethods are The HTTP methods to whitelist without requestid to be passed
	WhitelistHTTPMethods []string

	// ContextKeyName The requestid middleware takes requestid from the header validates it
	// if any pattern is given for validation and sets the value in your gin context
	// it generally sets the requestid with the following key name "RequestID" if you
	// want a different key name to be used you can pass the value here
	ContextKeyName string

	// Response The response that you want to return to the user when the requestid is not present in the headers
	// the default response will be {"error":"${{HeaderName}} is missing"}
	// you can pass any serializable value to pass in the response
	Response interface{}

	// StatusCode The httpStatusCode to return when the requestid is not present
	// by default it will be 403 , For more reference regarding status codes please check this out
	// https://go.dev/src/net/http/status.go
	StatusCode int
}

// DefaultConfig Values

const (
	DefaultHeaderName     = "Idempotency-Key"
	DefaultContextKeyName = "IdempotencyKey"
)

func getDefaultResponse(headerName string) gin.H {
	responseMessage := fmt.Sprintf("%s is missing", headerName)
	return gin.H{"error": responseMessage}
}

// Default returns a generic default configuration.
func Default() IdempotencyConfig {
	return IdempotencyConfig{
		HeaderName:     DefaultHeaderName,
		ContextKeyName: DefaultContextKeyName,
		StatusCode:     http.StatusForbidden,
		Response:       getDefaultResponse(DefaultContextKeyName),
	}
}

func newIdempotency(config IdempotencyConfig) IdempotencyConfig {
	cfg := IdempotencyConfig{}

	if cfg.HeaderName == "" {
		cfg.HeaderName = DefaultHeaderName
	} else {
		cfg.HeaderName = config.HeaderName
	}

	if cfg.ContextKeyName == "" {
		cfg.ContextKeyName = DefaultContextKeyName
	} else {
		cfg.ContextKeyName = config.ContextKeyName
	}

	if isNil(cfg.Response) {
		cfg.Response = getDefaultResponse(cfg.HeaderName)
	} else {
		cfg.Response = config.Response
	}

	if cfg.StatusCode == 0 {
		cfg.StatusCode = http.StatusForbidden
	} else {
		cfg.StatusCode = config.StatusCode
	}

	cfg.WhitelistHTTPMethods = config.WhitelistHTTPMethods

	return cfg
}

func isNil(c interface{}) bool {
	return c == nil || (reflect.ValueOf(c).Kind() == reflect.Ptr && reflect.ValueOf(c).IsNil())

}
