package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CustomerBank holds the schema definition for the CustomerBank entity.
type CustomerBank struct {
	ent.Schema
}

// Fields of the CustomerBank.
func (CustomerBank) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("bank_name"),
		field.String("branch_name"),
		field.String("bank_account_number"),
		field.String("bank_account_name"),
	}
}

// Edges of the CustomerBank.
func (CustomerBank) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).Ref("customer_banks").Unique().Required(),
	}
}

func (CustomerBank) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
