package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// StatusR holds the schema definition for the StatusR entity.
type StatusR struct {
	ent.Schema
}

// Fields of the StatusR.
func (StatusR) Fields() []ent.Field {
	return []ent.Field{
		field.String("Sname").Unique(),
	}
}

// Edges of the StatusR.
func (StatusR) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("repair_information", RepairInvoice.Type).StorageKey(edge.Column("statusr_id")),
	}
}
