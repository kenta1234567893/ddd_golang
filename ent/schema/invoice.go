package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Invoice holds the schema definition for the Invoice entity.
type Invoice struct {
	ent.Schema
}

// Fields of the Invoice.
func (Invoice) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Time("issue_date").
			SchemaType(map[string]string{
				dialect.MySQL: "DATETIME(6)",
			}),
		field.Float("payment_amount"),
		field.Float("fee"),
		field.Float("fee_rate"),
		field.Float("tax"),
		field.Float("tax_rate"),
		field.Float("billing_amount"),
		field.Time("payment_due_date").
			SchemaType(map[string]string{
				dialect.MySQL: "DATETIME(6)",
			}),
		field.String("status"),
	}
}

// Edges of the Invoice.
func (Invoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("invoices").Unique().Required(),
		edge.From("customer", Customer.Type).Ref("invoices").Unique().Required(),
	}
}

func (Invoice) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
