package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("company_name"),
		field.String("ceo_name"),
		field.String("phone_number"),
		field.String("zip_code"),
		field.String("address"),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("customer_banks", CustomerBank.Type),
		edge.To("invoices", Invoice.Type),
	}
}

func (Customer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
