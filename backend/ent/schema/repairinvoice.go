package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// RepairInvoice holds the schema definition for the RepairInvoice entity.
type RepairInvoice struct {
	ent.Schema
}

// Fields of the RepairInvoice.
func (RepairInvoice) Fields() []ent.Field {
	return []ent.Field{
		field.String("Rename").Unique(),
	}
}

// Edges of the RepairInvoice.
func (RepairInvoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("device", Device.Type).Ref("repair_information").Unique(),
		edge.From("status", StatusR.Type).Ref("repair_information").Unique(),
		edge.From("symptom", Symptom.Type).Ref("repair_information").Unique(),

		edge.From("user", User.Type).Ref("repairinvoice_informations").Unique(),
		edge.To("returninvoice", Returninvoice.Type).
			Unique().StorageKey(edge.Column("returninvoice_id")),
	}
}
