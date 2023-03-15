package controllerUtils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SendError(ctx *fiber.Ctx, errStr string, statusCode int16) error {
	ctx.Status(fiber.StatusInternalServerError).SendString(errStr)
	return fmt.Errorf(errStr)
}
