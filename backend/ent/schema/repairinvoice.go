package schema

import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)

// Repairinvoice holds the schema definition for the Repairinvoice entity.
type Repairinvoice struct {
	ent.Schema
}

// Fields of the Repairinvoice.
func (Repairinvoice) Fields() []ent.Field {
	return []ent.Field{
	   field.Int("symptomid"),
       field.Int("deviceid"),
	   field.Int("userid"),
	   field.Int("statusrepairid"),
   }
}

// Edges of the Repairinvoice.
func (Repairinvoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("repairinvoices", Returninvoice.Type).
			Unique().StorageKey(edge.Column("reparinvoice_id")),
	}
}
