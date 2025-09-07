package rate_limiter

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func UseRateLimiter(requestLimit int) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before request")

		if true {
			c.Next()
		} else {
			c.Abort()
		}

		fmt.Println("After request")
	}
}
