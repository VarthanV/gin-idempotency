# gin-idempotency


## Introduction
gin-idempotency is a middleware for the [Gin](https://gin-gonic.com/) webframework. It checks whether the Idempotency key is passed in the headers


An idempotency key is a unique value generated by you. Our servers use this key to recognize subsequent retries of the same request. 

It is used in places where we need to prevent the impact recognize subsequent retries of the same request. It is mainly used in payment/transaction API's. More about idempotency

- [Stripe Reference](http://bit.ly/3WYB3Wc)
- [Razorpay Reference](http://bit.ly/3latbDV)

## Usage

Download and install it using go get

```
go get github.com/VarthanV/gin-idempotency
```

Import it in your code

```golang
import "github.com/VarthanV/gin-idempotency"

```

## Examples

**using default config**

```golang
package main

import (
  "github.com/Varthan/idempotency"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.Use(idempotency.Default())

  r.POST("/transfer", func(c *gin.Context) {
    var existingIdempotentKey = "foo"
    // The value in the header will parsed and will be set in the context with the following key name by default 
    var idempotencyKeyFromCtx = c.GetString("IdempotencyKey")
    if idempotencyKeyFromCtx == existingIdempotentKey { 
        c.JSON(403, gin.H{"message":"send an new idempotency key"})
        return
    }
  })
  r.Run(":8000")
}
```

**using custom config**

```golang
package main

import (
  "github.com/Varthan/idempotency"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.Use(idempotency.New(idempotency.IdempotencyConfig{
    HeaderName: "foo" // The middleware by default looks for the header with the key name  ``Idempotency-Key``
    ContextKeyName: "foo-ctx" // The value in the header will parsed and will be set in the context with the  key name IdempotencyKey by default, You can customise it based on your needs
    StatusCode: 418 // The httpStatusCode which you want to return incase the idempotencykey is not present default is 403
    Response: map[string]string{"message":"idempotency-key-doesnt-exist"} //Response payload which you want to send to the client when the key is not present. Default response is {"error":"${{HeaderName}} is missing"}
    
  }))

  r.POST("/fund-transfer", func(c *gin.Context) {
    var existingIdempotentKey = "foo"
    // The value in the header will parsed and will be set in the context with the following key name by default 
    var idempotencyKeyFromCtx = c.GetString("IdempotencyKey")
    if idempotencyKeyFromCtx == existingIdempotentKey { 
        c.JSON(403, gin.H{"message":"send an new idempotency key"})
        return
    }
  })
  r.Run(":8000")
}
```