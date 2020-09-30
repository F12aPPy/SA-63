package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Babystatus holds the schema definition for the Babystatus entity.
type Babystatus struct {
	ent.Schema
}

// Fields of the Status.
func (Babystatus) Fields() []ent.Field {
	return []ent.Field{
		field.String("STATUS_BABY_NAME").NotEmpty(),
	}
}

// Edges of the Status.
func (Babystatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("SETSTATUS", Antenatal.Type),
	}
}
