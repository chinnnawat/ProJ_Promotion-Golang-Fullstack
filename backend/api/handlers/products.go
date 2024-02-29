package handlers

import (
	"context"

	"github.com/chinnnawat/ProJ_Promotion-Golang/backend/models/productsuc"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type CreateProductUseCase interface {
	CreateProduct(ctx context.Context, request productsuc.CreateProductRequest) (*productsuc.CreateProductResponse, error)
}

func CreateProductHandler(useCase CreateProductUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ctx = c.UserContext()

		var request = productsuc.CreateProductRequest{}

		err := c.BodyParser(&request)
		if err != nil {
			return errors.Wrap(err, "unable to parse incoming request")
		}

		response, err := useCase.CreateProduct(ctx, request)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(response)
	}
}
