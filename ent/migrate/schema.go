// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CompaniesColumns holds the columns for the "companies" table.
	CompaniesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, Default: "CURRENT_TIMESTAMP(6)", SchemaType: map[string]string{"mysql": "DATETIME(6)"}},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, Default: "CURRENT_TIMESTAMP(6)", SchemaType: map[string]string{"mysql": "DATETIME(6)"}},
		{Name: "company_name", Type: field.TypeString},
		{Name: "ceo_name", Type: field.TypeString},
		{Name: "phone_number", Type: field.TypeString},
		{Name: "zip_code", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
	}
	// CompaniesTable holds the schema information for the "companies" table.
	CompaniesTable = &schema.Table{
		Name:       "companies",
		Columns:    CompaniesColumns,
		PrimaryKey: []*schema.Column{CompaniesColumns[0]},
	}
	// CustomersColumns holds the columns for the "customers" table.
	CustomersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, Default: "CURRENT_TIMESTAMP(6)", SchemaType: map[string]string{"mysql": "DATETIME(6)"}},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, Default: "CURRENT_TIMESTAMP(6)", SchemaType: map[string]string{"mysql": "DATETIME(6)"}},
		{Name: "company_name", Type: field.TypeString},
		{Name: "ceo_name", Type: field.TypeString},
		{Name: "phone_number", Type: field.TypeString},
		{Name: "zip_code", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
	}
	// CustomersTable holds the schema information for the "customers" table.
	CustomersTable = &schema.Table{
		Name:       "customers",
		Columns:    CustomersColumns,
		PrimaryKey: []*schema.Column{CustomersColumns[0]},
	}
	// CustomerBanksColumns holds the columns for the "customer_banks" table.
	CustomerBanksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, Default: "CURRENT_TIMESTAMP(6)", SchemaType: map[string]string{"mysql": "DATETIME(6)"}},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, Default: "CURRENT_TIMESTAMP(6)", SchemaType: map[string]string{"mysql": "DATETIME(6)"}},
		{Name: "bank_name", Type: field.TypeString},
		{Name: "branch_name", Type: field.TypeString},
		{Name: "bank_account_number", Type: field.TypeString},
		{Name: "bank_account_name", Type: field.TypeString},
		{Name: "customer_customer_banks", Type: field.TypeUint64},
	}
	// CustomerBanksTable holds the schema information for the "customer_banks" table.
	CustomerBanksTable = &schema.Table{
		Name:       "customer_banks",
		Columns:    CustomerBanksColumns,
		PrimaryKey: []*schema.Column{CustomerBanksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "customer_banks_customers_customer_banks",
				Columns:    []*schema.Column{CustomerBanksColumns[7]},
				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// InvoicesColumns holds the columns for the "invoices" table.
	InvoicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, Default: "CURRENT_TIMESTAMP(6)", SchemaType: map[string]string{"mysql": "DATETIME(6)"}},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, Default: "CURRENT_TIMESTAMP(6)", SchemaType: map[string]string{"mysql": "DATETIME(6)"}},
		{Name: "issue_date", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "DATETIME(6)"}},
		{Name: "payment_amount", Type: field.TypeFloat64},
		{Name: "fee", Type: field.TypeFloat64},
		{Name: "fee_rate", Type: field.TypeFloat64},
		{Name: "tax", Type: field.TypeFloat64},
		{Name: "tax_rate", Type: field.TypeFloat64},
		{Name: "billing_amount", Type: field.TypeFloat64},
		{Name: "payment_due_date", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "DATETIME(6)"}},
		{Name: "status", Type: field.TypeString},
		{Name: "company_invoices", Type: field.TypeUint64},
		{Name: "customer_invoices", Type: field.TypeUint64},
	}
	// InvoicesTable holds the schema information for the "invoices" table.
	InvoicesTable = &schema.Table{
		Name:       "invoices",
		Columns:    InvoicesColumns,
		PrimaryKey: []*schema.Column{InvoicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "invoices_companies_invoices",
				Columns:    []*schema.Column{InvoicesColumns[12]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "invoices_customers_invoices",
				Columns:    []*schema.Column{InvoicesColumns[13]},
				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, Default: "CURRENT_TIMESTAMP(6)", SchemaType: map[string]string{"mysql": "DATETIME(6)"}},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, Default: "CURRENT_TIMESTAMP(6)", SchemaType: map[string]string{"mysql": "DATETIME(6)"}},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeFloat64},
		{Name: "password", Type: field.TypeFloat64},
		{Name: "company_users", Type: field.TypeUint64},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_companies_users",
				Columns:    []*schema.Column{UsersColumns[6]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CompaniesTable,
		CustomersTable,
		CustomerBanksTable,
		InvoicesTable,
		UsersTable,
	}
)

func init() {
	CustomerBanksTable.ForeignKeys[0].RefTable = CustomersTable
	InvoicesTable.ForeignKeys[0].RefTable = CompaniesTable
	InvoicesTable.ForeignKeys[1].RefTable = CustomersTable
	UsersTable.ForeignKeys[0].RefTable = CompaniesTable
}
