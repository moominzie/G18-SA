package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Returninvoice holds the schema definition for the Returninvoice entity.
type Returninvoice struct {
	ent.Schema
}

// Fields of the Returninvoice.
func (Returninvoice) Fields() []ent.Field {
	return []ent.Field{
		field.Time("addedtime"),
	}
}

// Edges of the Returninvoice.
func (Returninvoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Repairinvoice", RepairInvoice.Type).Ref("returninvoice").Unique().Required(),
		edge.From("Employee", Employee.Type).Ref("employees").Unique(),
		edge.From("Statust", Statust.Type).Ref("statusts").Unique(),
	}
}
