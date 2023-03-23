package controllerUtils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SendError(ctx *fiber.Ctx, errStr string, statusCode int) error {
	ctx.Status(statusCode).SendString(errStr)
	return fmt.Errorf(errStr)
}
