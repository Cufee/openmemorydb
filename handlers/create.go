package handlers

import (
	"byvko.dev/repo/openmemorydb/database/operations"
	"byvko.dev/repo/openmemorydb/types"
	"github.com/gofiber/fiber/v2"
)

func CreateOneHandler(ctx *fiber.Ctx) error {
	var request types.OperationRequestCreateOne
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.OperationResult{Error: err.Error()})
	}

	result, err := operations.CreateOne(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(result)
	}
	return ctx.Status(fiber.StatusCreated).JSON(result)
}

func CreateManyHandler(ctx *fiber.Ctx) error {
	var request types.OperationRequestCreateMany
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.OperationResult{Error: err.Error()})
	}

	result, err := operations.CreateMany(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(result)
	}
	return ctx.Status(fiber.StatusCreated).JSON(result)
}
