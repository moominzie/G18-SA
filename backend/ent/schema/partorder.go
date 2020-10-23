package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

// Partorder holds the schema definition for the Partorder entity.
type Partorder struct {
	ent.Schema
}

// Fields of the Partorder.
func (Partorder) Fields() []ent.Field {
	return nil
}

// Edges of the Partorder.
func (Partorder) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("repairinvoice", RepairInvoice.Type).Ref("part_informations").Unique(),
		edge.From("employee", Employee.Type).Ref("employeepart").Unique(),
		edge.From("part", Part.Type).Ref("part_informations").Unique(),
	}
}
