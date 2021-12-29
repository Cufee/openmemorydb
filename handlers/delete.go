package handlers

import (
	"byvko.dev/repo/openmemorydb/database/operations"
	"byvko.dev/repo/openmemorydb/types"
	"github.com/gofiber/fiber/v2"
)

func DeleteOneHandler(ctx *fiber.Ctx) error {
	var request types.OperationRequestDeleteOne
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.OperationResult{Error: err.Error()})
	}

	if request.Operation != types.DeleteOperationName {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.OperationResult{Error: "invalid operation name"})
	}

	if !request.IsValid() {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.OperationResult{Error: "invalid request"})
	}

	result, err := operations.DeleteOne(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(result)
	}
	return ctx.Status(fiber.StatusOK).JSON(result)
}

func DeleteManyHandler(ctx *fiber.Ctx) error {
	var request types.OperationRequestDeleteMany
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.OperationResult{Error: err.Error()})
	}

	if request.Operation != types.DeleteOperationName {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.OperationResult{Error: "invalid operation name"})
	}

	if !request.IsValid() {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.OperationResult{Error: "invalid request"})
	}

	result, err := operations.DeleteMany(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(result)
	}
	return ctx.Status(fiber.StatusOK).JSON(result)
}
