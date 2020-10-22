package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Building holds the schema definition for the Building entity.
type Building struct {
	ent.Schema
}

// Fields of the Building.
func (Building) Fields() []ent.Field {
	return []ent.Field{
		field.String("buname").Unique(),
	}
}

// Edges of the Building.
func (Building) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rooms", Room.Type).StorageKey(edge.Column("building_id")),
		edge.To("user_informations", User.Type).StorageKey(edge.Column("building_id")),
	}
}
