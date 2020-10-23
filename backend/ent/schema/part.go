package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Part holds the schema definition for the Part entity.
type Part struct {
	ent.Schema
}

// Fields of the Part.
func (Part) Fields() []ent.Field {
	return []ent.Field{
		field.String("Pname").Unique(),
	}
}

// Edges of the Part.
func (Part) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("part_informations", Partorder.Type).StorageKey(edge.Column("part_id")),
	}
}
