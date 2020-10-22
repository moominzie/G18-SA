package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Faculty holds the schema definition for the Faculty entity.
type Faculty struct {
	ent.Schema
}

// Fields of the Faculty.
func (Faculty) Fields() []ent.Field {
	return []ent.Field{
		field.String("fname").Unique(),
	}
}

// Edges of the Faculty.
func (Faculty) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("branchs", Branch.Type).StorageKey(edge.Column("faculty_id")),
		edge.To("user_informations", User.Type).StorageKey(edge.Column("faculty_id")),
	}
}
