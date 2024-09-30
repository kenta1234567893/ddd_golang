// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/kenta1234567893/upsider-coding-test/ent/customer"
	"github.com/kenta1234567893/upsider-coding-test/ent/customerbank"
)

// CustomerBank is the model entity for the CustomerBank schema.
type CustomerBank struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// BankName holds the value of the "bank_name" field.
	BankName string `json:"bank_name,omitempty"`
	// BranchName holds the value of the "branch_name" field.
	BranchName string `json:"branch_name,omitempty"`
	// BankAccountNumber holds the value of the "bank_account_number" field.
	BankAccountNumber string `json:"bank_account_number,omitempty"`
	// BankAccountName holds the value of the "bank_account_name" field.
	BankAccountName string `json:"bank_account_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CustomerBankQuery when eager-loading is set.
	Edges                   CustomerBankEdges `json:"edges"`
	customer_customer_banks *uint64
	selectValues            sql.SelectValues
}

// CustomerBankEdges holds the relations/edges for other nodes in the graph.
type CustomerBankEdges struct {
	// Customer holds the value of the customer edge.
	Customer *Customer `json:"customer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerBankEdges) CustomerOrErr() (*Customer, error) {
	if e.Customer != nil {
		return e.Customer, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: customer.Label}
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CustomerBank) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case customerbank.FieldID:
			values[i] = new(sql.NullInt64)
		case customerbank.FieldBankName, customerbank.FieldBranchName, customerbank.FieldBankAccountNumber, customerbank.FieldBankAccountName:
			values[i] = new(sql.NullString)
		case customerbank.FieldCreatedAt, customerbank.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case customerbank.ForeignKeys[0]: // customer_customer_banks
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CustomerBank fields.
func (cb *CustomerBank) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case customerbank.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cb.ID = uint64(value.Int64)
		case customerbank.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cb.CreatedAt = value.Time
			}
		case customerbank.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cb.UpdatedAt = value.Time
			}
		case customerbank.FieldBankName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bank_name", values[i])
			} else if value.Valid {
				cb.BankName = value.String
			}
		case customerbank.FieldBranchName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field branch_name", values[i])
			} else if value.Valid {
				cb.BranchName = value.String
			}
		case customerbank.FieldBankAccountNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bank_account_number", values[i])
			} else if value.Valid {
				cb.BankAccountNumber = value.String
			}
		case customerbank.FieldBankAccountName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bank_account_name", values[i])
			} else if value.Valid {
				cb.BankAccountName = value.String
			}
		case customerbank.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field customer_customer_banks", value)
			} else if value.Valid {
				cb.customer_customer_banks = new(uint64)
				*cb.customer_customer_banks = uint64(value.Int64)
			}
		default:
			cb.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CustomerBank.
// This includes values selected through modifiers, order, etc.
func (cb *CustomerBank) Value(name string) (ent.Value, error) {
	return cb.selectValues.Get(name)
}

// QueryCustomer queries the "customer" edge of the CustomerBank entity.
func (cb *CustomerBank) QueryCustomer() *CustomerQuery {
	return NewCustomerBankClient(cb.config).QueryCustomer(cb)
}

// Update returns a builder for updating this CustomerBank.
// Note that you need to call CustomerBank.Unwrap() before calling this method if this CustomerBank
// was returned from a transaction, and the transaction was committed or rolled back.
func (cb *CustomerBank) Update() *CustomerBankUpdateOne {
	return NewCustomerBankClient(cb.config).UpdateOne(cb)
}

// Unwrap unwraps the CustomerBank entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cb *CustomerBank) Unwrap() *CustomerBank {
	_tx, ok := cb.config.driver.(*txDriver)
	if !ok {
		panic("ent: CustomerBank is not a transactional entity")
	}
	cb.config.driver = _tx.drv
	return cb
}

// String implements the fmt.Stringer.
func (cb *CustomerBank) String() string {
	var builder strings.Builder
	builder.WriteString("CustomerBank(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cb.ID))
	builder.WriteString("created_at=")
	builder.WriteString(cb.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(cb.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("bank_name=")
	builder.WriteString(cb.BankName)
	builder.WriteString(", ")
	builder.WriteString("branch_name=")
	builder.WriteString(cb.BranchName)
	builder.WriteString(", ")
	builder.WriteString("bank_account_number=")
	builder.WriteString(cb.BankAccountNumber)
	builder.WriteString(", ")
	builder.WriteString("bank_account_name=")
	builder.WriteString(cb.BankAccountName)
	builder.WriteByte(')')
	return builder.String()
}

// CustomerBanks is a parsable slice of CustomerBank.
type CustomerBanks []*CustomerBank
