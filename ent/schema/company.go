package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Company holds the schema definition for the Company entity.
type Company struct {
	ent.Schema
}

// Fields of the Company.
func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("company_name"),
		field.String("ceo_name"),
		field.String("phone_number"),
		field.String("zip_code"),
		field.String("address"),
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
		edge.To("invoices", Invoice.Type),
	}
}

func (Company) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
