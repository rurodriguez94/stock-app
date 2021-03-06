package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/stock-app/stock-api/internal/pkg/apierror"
	"github.com/stock-app/stock-api/internal/pkg/logs"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		err := c.Errors.Last()
		logs.Log().Error(err.Error())

		apiErr := parseAPIError(c.Errors.Last())

		c.JSON(apiErr.Status, apiErr)
	}
}

func parseAPIError(err *gin.Error) apierror.APIError {
	switch err.Err.(type) {
	case apierror.APIError:
		return err.Err.(apierror.APIError)
	default:
		return apierror.NewStatusInternalServerError("Internal Server Error")
	}
}
