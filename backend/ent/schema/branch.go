package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Branch holds the schema definition for the Branch entity.
type Branch struct {
	ent.Schema
}

// Fields of the Branch.
func (Branch) Fields() []ent.Field {
	return []ent.Field{
		field.String("brname").Unique(),
	}
}

// Edges of the Branch.
func (Branch) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("faculty", Faculty.Type).Ref("branchs").Unique(),
		edge.To("user_informations", User.Type).StorageKey(edge.Column("branch_id")),
	}
}
