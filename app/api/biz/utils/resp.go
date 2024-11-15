package utils

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/czczcz831/tiktok-mall/app/api/biz/utils/packer"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	c.JSON(code, packer.NewErrorResponse(err))
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, packer.NewSuccessResponse(data))
}
