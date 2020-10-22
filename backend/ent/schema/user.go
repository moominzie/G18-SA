package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("personalID").Unique(),
		field.String("personalName"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("faculty", Faculty.Type).Ref("user_informations").Unique(),
		edge.From("branch", Branch.Type).Ref("user_informations").Unique(),
		edge.From("building", Building.Type).Ref("user_informations").Unique(),
		edge.From("room", Room.Type).Ref("user_informations").Unique(),
	}
}
