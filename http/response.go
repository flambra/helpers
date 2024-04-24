package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type DefaultResponse struct {
	Message string `json:"message,omitempty"`
}

// Returns a 200 OK status.
func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusOK).JSON(data)
}

// Returns a 201 Created status.
func SuccessCreated(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusCreated).JSON(data)
}

// Returns a 400 Bad Request.
func BadRequestResponse(c *fiber.Ctx, message string) error {
	return c.Status(http.StatusBadRequest).JSON(DefaultResponse{
		Message: message,
	})
}

// Returns a 401 Unauthorized status.
func UnauthorizedResponse(c *fiber.Ctx, message string) error {
	return c.Status(http.StatusUnauthorized).JSON(DefaultResponse{
		Message: message,
	})
}

// Returns a 403 Forbidden status.
func ForbiddenResponse(c *fiber.Ctx, message string) error {
	return c.Status(http.StatusForbidden).JSON(DefaultResponse{
		Message: message,
	})
}

// Returns a 404 Not Found status.
func NotFoundResponse(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(http.StatusNotFound).JSON(DefaultResponse{
		Message: message,
	})
}

// Returns a 409 Conflict status.
func StatusConflict(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(http.StatusConflict).JSON(DefaultResponse{
		Message: message,
	})
}

// Returns a 422 Unprocessable Entity.
func UnprocessableResponse(c *fiber.Ctx, message string) error {
	return c.Status(http.StatusUnprocessableEntity).JSON(DefaultResponse{
		Message: message,
	})
}

// Returns a 429 Too Many Requests.
func TooManyRequestResponse(c *fiber.Ctx, message string) error {
	return c.Status(http.StatusTooManyRequests).JSON(DefaultResponse{
		Message: message,
	})
}

// Returns a 500 Internal Server Error.
func InternalServerErrorResponse(c *fiber.Ctx, message string) error {
	return c.Status(http.StatusInternalServerError).JSON(DefaultResponse{
		Message: message,
	})
}
