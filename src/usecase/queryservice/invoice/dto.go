package invoice

import (
	"time"
)

type QueryCondition struct {
	CompanyID           uint64
	Status              []string
	PaymentDueDateStart *time.Time
	PaymentDueDateEnd   *time.Time
}

type Invoice struct {
	ID             uint64
	IssueDate      time.Time
	PaymentAmount  float64
	Fee            float64
	FeeRate        float64
	Tax            float64
	TaxRate        float64
	BillingAmount  float64
	PaymentDueDate time.Time
	Status         string
}
