package handlers

import (
	"fmt"

	"byvko.dev/repo/openmemorydb/database/operations"
	"byvko.dev/repo/openmemorydb/types"
	"github.com/gofiber/fiber/v2"
)

func UpdateOneHandler(ctx *fiber.Ctx) error {
	var request types.OperationRequestUpdateOne
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.OperationResult{Error: err.Error()})
	}

	if request.Operation != types.UpdateOperationName {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.OperationResult{Error: "invalid operation name"})
	}

	if !request.Update.IsValid() {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.OperationResult{Error: "invalid request"})
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

	if request.Operation != types.UpdateOperationName {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.OperationResult{Error: "invalid operation name"})
	}

	for i, update := range request.Updates {
		if !update.IsValid() {
			return ctx.Status(fiber.StatusBadRequest).JSON(types.OperationResult{Error: fmt.Sprintf("invalid request with index %v", i)})
		}
	}

	result, err := operations.UpdateMany(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(result)
	}
	return ctx.Status(fiber.StatusOK).JSON(result)
}
