package backend

import (
	"github.com/MQEnergy/mqshop/internal/app/controller"
	"github.com/MQEnergy/mqshop/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type BrandController struct {
	controller.Controller
}

var Brand = &BrandController{}

// Index ...
func (s *BrandController) Index(ctx *fiber.Ctx) error {
	// Todo implement ...
	return response.SuccessJSON(ctx, "", "")
}
