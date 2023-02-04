package idempotency

import "github.com/gin-gonic/gin"

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (r *IdempotencyConfig) applyToContext(ctx *gin.Context) {
	if contains(r.WhitelistHTTPMethods, ctx.Request.Method) {
		ctx.Next()
	}
	requestID := ctx.GetHeader(r.HeaderName)
	if requestID == "" {
		ctx.AbortWithStatusJSON(r.StatusCode, r.Response)
		return
	}
	ctx.Set(r.HeaderName, requestID)
	ctx.Next()
}

func New(config IdempotencyConfig) gin.HandlerFunc {
	r := newIdempotency(config)
	return func(ctx *gin.Context) {
		r.applyToContext(ctx)
	}
}
