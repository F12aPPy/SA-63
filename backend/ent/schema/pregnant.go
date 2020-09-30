package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Pregnant holds the schema definition for the Pregnant entity.
type Pregnant struct {
	ent.Schema
}

// Fields of the Pregnant.
func (Pregnant) Fields() []ent.Field {
	return []ent.Field{
		field.String("PREGNANT_NAME").NotEmpty(),
		field.Int("PREGNANT_AGE").Positive(),
	}
}

// Edges of the Pregnant.
func (Pregnant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("SETMOM", Antenatal.Type),
	}
}
