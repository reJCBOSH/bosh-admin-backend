package ctx

import "github.com/gin-gonic/gin"

type Context struct {
	*gin.Context
}

type HandlerFunc func(ctx *Context)

func Handler(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := new(Context)
		context.Context = c
		h(context)
	}
}
