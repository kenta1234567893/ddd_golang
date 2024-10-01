package infra

import (
	"github.com/google/wire"
	"github.com/kenta1234567893/upsider-coding-test/src/infra/repository"
)

var InfraSet = wire.NewSet(
	repository.NewInvoiceRepository,
)
