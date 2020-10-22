package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Billingstatus holds the schema definition for the Billingstatus entity.
type Billingstatus struct {
	ent.Schema
}

// Fields of the Billingstatus.
func (Billingstatus) Fields() []ent.Field {
	return []ent.Field{
		field.String("billingstatusname").Unique(),
	}
}

// Edges of the Billingstatus.
func (Billingstatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("billingstatuss", Bill.Type).StorageKey(edge.Column("billingstatus_id")),
	}
}
