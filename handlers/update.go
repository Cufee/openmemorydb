package handlers

import (
	"byvko.dev/repo/openmemorydb/database/operations"
	"byvko.dev/repo/openmemorydb/types"
	"github.com/gofiber/fiber/v2"
)

func UpdateOneHandler(ctx *fiber.Ctx) error {
	var request types.OperationRequestUpdateOne
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.OperationResult{Error: err.Error()})
	}

	result, err := operations.UpdateOne(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(result)
	}
	return ctx.Status(fiber.StatusOK).JSON(result)
}

func UpdateManyHandler(ctx *fiber.Ctx) error {
	var request types.OperationRequestUpdateMany
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.OperationResult{Error: err.Error()})
	}

	result, err := operations.UpdateMany(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(result)
	}
	return ctx.Status(fiber.StatusOK).JSON(result)
}
