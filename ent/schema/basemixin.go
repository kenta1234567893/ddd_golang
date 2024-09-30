package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type BaseMixin struct {
	mixin.Schema
}

func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Optional().
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "DATETIME(6)",
			}).
			Annotations(entsql.Default("CURRENT_TIMESTAMP(6)")),
		field.Time("updated_at").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "DATETIME(6)",
			}).
			Annotations(entsql.Default("CURRENT_TIMESTAMP(6)")),
	}
}
