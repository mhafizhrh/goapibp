package response

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type CustomResponse struct {
	status     string
	statusCode int
	data       any
}

func (c *CustomResponse) Error() string {
	return fmt.Sprint(c.statusCode)
}
func (c *CustomResponse) Status() string {
	return c.status
}
func (c *CustomResponse) StatusCode() int {
	return c.statusCode
}
func (c *CustomResponse) Data() any {
	return c.data
}
func (c *CustomResponse) ToFiberResponse(fiberCtx *fiber.Ctx) error {
	return fiberCtx.Status(c.statusCode).JSON(c.data)
}

func NewResponse(statusCode int, data any) CustomResponse {
	return CustomResponse{
		status:     http.StatusText(statusCode),
		statusCode: statusCode,
		data:       data,
	}
}
