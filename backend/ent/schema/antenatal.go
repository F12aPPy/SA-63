package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Antenatal holds the schema definition for the Antenatal entity.
type Antenatal struct {
	ent.Schema
}

// Fields of the Antenatal.
func (Antenatal) Fields() []ent.Field {
	return []ent.Field{
		field.Time("ADDED_TIME").Default(time.Now),
	}
}

// Edges of the Antenatal.
func (Antenatal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("GETMOM", Pregnant.Type).Ref("SETMOM").Unique(),
		edge.From("TAKECARE", User.Type).Ref("CARE").Unique(),
		edge.From("GETSTATUS", Babystatus.Type).Ref("SETSTATUS").Unique(),
	}
}
