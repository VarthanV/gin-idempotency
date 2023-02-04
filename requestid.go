package requestid

import "github.com/gin-gonic/gin"

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (r *RequestID) applyToContext(ctx *gin.Context) {
	if contains(r.WhitelistHTTPsMethods, ctx.Request.Method) {
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

func New(config RequestID) gin.HandlerFunc {
	r := newRequestID(config)
	return func(ctx *gin.Context) {
		r.applyToContext(ctx)
	}
}
