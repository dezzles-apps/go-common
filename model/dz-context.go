package model

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DAContext struct {
	Logger   *zap.Logger
	Ctx      *context.Context
	UserId   *string
	UserType *string
}

// DAContextKey is the gin.Context key used to store a DAContext.
const DAContextKey = "daContext"

func GetContext(c *gin.Context) *DAContext {
	if value, exists := c.Get(DAContextKey); exists {
		if ctx, ok := value.(*DAContext); ok && ctx != nil {
			return ctx
		}
	}

	var logger *zap.Logger
	loggerVal, exists := c.Get("logger")
	if exists {
		logger = loggerVal.(*zap.Logger)
	} else {
		logger = zap.NewNop()
	}
	ctx := c.Request.Context()
	daContext := &DAContext{
		Logger: logger,
		Ctx:    &ctx,
	}
	c.Set(DAContextKey, daContext)
	return daContext
}
