// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/kenta1234567893/upsider-coding-test/ent/company"
	"github.com/kenta1234567893/upsider-coding-test/ent/customer"
	"github.com/kenta1234567893/upsider-coding-test/ent/invoice"
)

// Invoice is the model entity for the Invoice schema.
type Invoice struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// IssueDate holds the value of the "issue_date" field.
	IssueDate time.Time `json:"issue_date,omitempty"`
	// PaymentAmount holds the value of the "payment_amount" field.
	PaymentAmount float64 `json:"payment_amount,omitempty"`
	// Fee holds the value of the "fee" field.
	Fee float64 `json:"fee,omitempty"`
	// FeeRate holds the value of the "fee_rate" field.
	FeeRate float64 `json:"fee_rate,omitempty"`
	// Tax holds the value of the "tax" field.
	Tax float64 `json:"tax,omitempty"`
	// TaxRate holds the value of the "tax_rate" field.
	TaxRate float64 `json:"tax_rate,omitempty"`
	// BillingAmount holds the value of the "billing_amount" field.
	BillingAmount float64 `json:"billing_amount,omitempty"`
	// PaymentDueDate holds the value of the "payment_due_date" field.
	PaymentDueDate time.Time `json:"payment_due_date,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InvoiceQuery when eager-loading is set.
	Edges             InvoiceEdges `json:"edges"`
	company_invoices  *uint64
	customer_invoices *uint64
	selectValues      sql.SelectValues
}

// InvoiceEdges holds the relations/edges for other nodes in the graph.
type InvoiceEdges struct {
	// Company holds the value of the company edge.
	Company *Company `json:"company,omitempty"`
	// Customer holds the value of the customer edge.
	Customer *Customer `json:"customer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InvoiceEdges) CompanyOrErr() (*Company, error) {
	if e.Company != nil {
		return e.Company, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: company.Label}
	}
	return nil, &NotLoadedError{edge: "company"}
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InvoiceEdges) CustomerOrErr() (*Customer, error) {
	if e.Customer != nil {
		return e.Customer, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: customer.Label}
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Invoice) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case invoice.FieldPaymentAmount, invoice.FieldFee, invoice.FieldFeeRate, invoice.FieldTax, invoice.FieldTaxRate, invoice.FieldBillingAmount:
			values[i] = new(sql.NullFloat64)
		case invoice.FieldID:
			values[i] = new(sql.NullInt64)
		case invoice.FieldStatus:
			values[i] = new(sql.NullString)
		case invoice.FieldCreatedAt, invoice.FieldUpdatedAt, invoice.FieldIssueDate, invoice.FieldPaymentDueDate:
			values[i] = new(sql.NullTime)
		case invoice.ForeignKeys[0]: // company_invoices
			values[i] = new(sql.NullInt64)
		case invoice.ForeignKeys[1]: // customer_invoices
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Invoice fields.
func (i *Invoice) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case invoice.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = uint64(value.Int64)
		case invoice.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case invoice.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		case invoice.FieldIssueDate:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field issue_date", values[j])
			} else if value.Valid {
				i.IssueDate = value.Time
			}
		case invoice.FieldPaymentAmount:
			if value, ok := values[j].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field payment_amount", values[j])
			} else if value.Valid {
				i.PaymentAmount = value.Float64
			}
		case invoice.FieldFee:
			if value, ok := values[j].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field fee", values[j])
			} else if value.Valid {
				i.Fee = value.Float64
			}
		case invoice.FieldFeeRate:
			if value, ok := values[j].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field fee_rate", values[j])
			} else if value.Valid {
				i.FeeRate = value.Float64
			}
		case invoice.FieldTax:
			if value, ok := values[j].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field tax", values[j])
			} else if value.Valid {
				i.Tax = value.Float64
			}
		case invoice.FieldTaxRate:
			if value, ok := values[j].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field tax_rate", values[j])
			} else if value.Valid {
				i.TaxRate = value.Float64
			}
		case invoice.FieldBillingAmount:
			if value, ok := values[j].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field billing_amount", values[j])
			} else if value.Valid {
				i.BillingAmount = value.Float64
			}
		case invoice.FieldPaymentDueDate:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field payment_due_date", values[j])
			} else if value.Valid {
				i.PaymentDueDate = value.Time
			}
		case invoice.FieldStatus:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[j])
			} else if value.Valid {
				i.Status = value.String
			}
		case invoice.ForeignKeys[0]:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field company_invoices", value)
			} else if value.Valid {
				i.company_invoices = new(uint64)
				*i.company_invoices = uint64(value.Int64)
			}
		case invoice.ForeignKeys[1]:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field customer_invoices", value)
			} else if value.Valid {
				i.customer_invoices = new(uint64)
				*i.customer_invoices = uint64(value.Int64)
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Invoice.
// This includes values selected through modifiers, order, etc.
func (i *Invoice) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// QueryCompany queries the "company" edge of the Invoice entity.
func (i *Invoice) QueryCompany() *CompanyQuery {
	return NewInvoiceClient(i.config).QueryCompany(i)
}

// QueryCustomer queries the "customer" edge of the Invoice entity.
func (i *Invoice) QueryCustomer() *CustomerQuery {
	return NewInvoiceClient(i.config).QueryCustomer(i)
}

// Update returns a builder for updating this Invoice.
// Note that you need to call Invoice.Unwrap() before calling this method if this Invoice
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Invoice) Update() *InvoiceUpdateOne {
	return NewInvoiceClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Invoice entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Invoice) Unwrap() *Invoice {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Invoice is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Invoice) String() string {
	var builder strings.Builder
	builder.WriteString("Invoice(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("issue_date=")
	builder.WriteString(i.IssueDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("payment_amount=")
	builder.WriteString(fmt.Sprintf("%v", i.PaymentAmount))
	builder.WriteString(", ")
	builder.WriteString("fee=")
	builder.WriteString(fmt.Sprintf("%v", i.Fee))
	builder.WriteString(", ")
	builder.WriteString("fee_rate=")
	builder.WriteString(fmt.Sprintf("%v", i.FeeRate))
	builder.WriteString(", ")
	builder.WriteString("tax=")
	builder.WriteString(fmt.Sprintf("%v", i.Tax))
	builder.WriteString(", ")
	builder.WriteString("tax_rate=")
	builder.WriteString(fmt.Sprintf("%v", i.TaxRate))
	builder.WriteString(", ")
	builder.WriteString("billing_amount=")
	builder.WriteString(fmt.Sprintf("%v", i.BillingAmount))
	builder.WriteString(", ")
	builder.WriteString("payment_due_date=")
	builder.WriteString(i.PaymentDueDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(i.Status)
	builder.WriteByte(')')
	return builder.String()
}

// Invoices is a parsable slice of Invoice.
type Invoices []*Invoice