package usecase

import (
	"github.com/google/wire"
	invoiceUC "github.com/kenta1234567893/upsider-coding-test/src/usecase/invoice"
)

var UsecaseSet = wire.NewSet(
	invoiceUC.NewCreateInvoiceUsecase,
)
