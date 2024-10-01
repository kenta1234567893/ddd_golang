package model

type InvoiceStatus struct {
	ValueObject[string]
}

func InvoiceStatusPending() InvoiceStatus {
	return InvoiceStatus{NewValueObject("pending")}
}

func InvoiceStatusProgress() InvoiceStatus {
	return InvoiceStatus{NewValueObject("progress")}
}

func InvoiceStatusPaid() InvoiceStatus {
	return InvoiceStatus{NewValueObject("paid")}
}

func InvoiceStatusError() InvoiceStatus {
	return InvoiceStatus{NewValueObject("error")}
}
