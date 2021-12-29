package handlers

import (
	"fmt"

	"byvko.dev/repo/openmemorydb/database/operations"
	"byvko.dev/repo/openmemorydb/types"
	"github.com/gofiber/fiber/v2"
)

func CreateOneHandler(ctx *fiber.Ctx) error {
	var request types.OperationRequestCreateOne
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.OperationResult{Error: err.Error()})
	}

	if request.Operation != types.CreateOperationName {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.OperationResult{Error: "invalid operation name"})
	}

	if !request.Document.IsValid() {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.OperationResult{Error: "invalid document"})
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

	if request.Operation != types.CreateOperationName {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.OperationResult{Error: "invalid operation name"})
	}

	for i, document := range request.Documents {
		if !document.IsValid() {
			return ctx.Status(fiber.StatusBadRequest).JSON(types.OperationResult{Error: fmt.Sprintf("invalid document with index %v", i)})
		}
	}

	result, err := operations.CreateMany(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(result)
	}
	return ctx.Status(fiber.StatusCreated).JSON(result)
}
