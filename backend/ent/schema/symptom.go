package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Symptom holds the schema definition for the Symptom entity.
type Symptom struct {
	ent.Schema
}

// Fields of the Symptom.
func (Symptom) Fields() []ent.Field {
	return []ent.Field{
		field.String("Syname").Unique(),
	}
}

// Edges of the Symptom.
func (Symptom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("repair_information", RepairInvoice.Type).StorageKey(edge.Column("symptom_id")),
	}
}
