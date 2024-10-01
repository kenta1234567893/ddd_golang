//go:build wireinject
// +build wireinject

package libs

import (
	"github.com/google/wire"
	"github.com/kenta1234567893/upsider-coding-test/ent"
	"github.com/kenta1234567893/upsider-coding-test/src/controller"
	"github.com/kenta1234567893/upsider-coding-test/src/infra"
	"github.com/kenta1234567893/upsider-coding-test/src/usecase"
)

type InitializeControllers struct {
	InvoiceController *controller.InvoiceController
}

func InitializeController(client *ent.Client) *InitializeControllers {
	wire.Build(
		controller.ControllerSet,
		infra.InfraSet,
		usecase.UsecaseSet,
		wire.Struct(new(InitializeControllers), "*"),
	)
	return &InitializeControllers{}
}
