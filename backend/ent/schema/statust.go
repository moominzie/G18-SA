package schema

import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)

// Statust holds the schema definition for the Statust entity.
type Statust struct {
	ent.Schema
}

// Fields of the Statust.
func (Statust) Fields() []ent.Field {
	return []ent.Field{
       field.String("statustname").Unique().NotEmpty(),
   }
}

// Edges of the Statust.
func (Statust) Edges() []ent.Edge {
	 return []ent.Edge{
        edge.To("statusts", Returninvoice.Type).StorageKey(edge.Column("statust_id")),
    }
}
