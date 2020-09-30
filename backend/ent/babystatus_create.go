// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/F12aPPy/app/ent/antenatal"
	"github.com/F12aPPy/app/ent/babystatus"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// BabystatusCreate is the builder for creating a Babystatus entity.
type BabystatusCreate struct {
	config
	mutation *BabystatusMutation
	hooks    []Hook
}

// SetSTATUSBABYNAME sets the STATUS_BABY_NAME field.
func (bc *BabystatusCreate) SetSTATUSBABYNAME(s string) *BabystatusCreate {
	bc.mutation.SetSTATUSBABYNAME(s)
	return bc
}

// AddSETSTATUSIDs adds the SETSTATUS edge to Antenatal by ids.
func (bc *BabystatusCreate) AddSETSTATUSIDs(ids ...int) *BabystatusCreate {
	bc.mutation.AddSETSTATUSIDs(ids...)
	return bc
}

// AddSETSTATUS adds the SETSTATUS edges to Antenatal.
func (bc *BabystatusCreate) AddSETSTATUS(a ...*Antenatal) *BabystatusCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return bc.AddSETSTATUSIDs(ids...)
}

// Mutation returns the BabystatusMutation object of the builder.
func (bc *BabystatusCreate) Mutation() *BabystatusMutation {
	return bc.mutation
}

// Save creates the Babystatus in the database.
func (bc *BabystatusCreate) Save(ctx context.Context) (*Babystatus, error) {
	if _, ok := bc.mutation.STATUSBABYNAME(); !ok {
		return nil, &ValidationError{Name: "STATUS_BABY_NAME", err: errors.New("ent: missing required field \"STATUS_BABY_NAME\"")}
	}
	if v, ok := bc.mutation.STATUSBABYNAME(); ok {
		if err := babystatus.STATUSBABYNAMEValidator(v); err != nil {
			return nil, &ValidationError{Name: "STATUS_BABY_NAME", err: fmt.Errorf("ent: validator failed for field \"STATUS_BABY_NAME\": %w", err)}
		}
	}
	var (
		err  error
		node *Babystatus
	)
	if len(bc.hooks) == 0 {
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BabystatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bc.mutation = mutation
			node, err = bc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BabystatusCreate) SaveX(ctx context.Context) *Babystatus {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bc *BabystatusCreate) sqlSave(ctx context.Context) (*Babystatus, error) {
	b, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	b.ID = int(id)
	return b, nil
}

func (bc *BabystatusCreate) createSpec() (*Babystatus, *sqlgraph.CreateSpec) {
	var (
		b     = &Babystatus{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: babystatus.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: babystatus.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.STATUSBABYNAME(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: babystatus.FieldSTATUSBABYNAME,
		})
		b.STATUSBABYNAME = value
	}
	if nodes := bc.mutation.SETSTATUSIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   babystatus.SETSTATUSTable,
			Columns: []string{babystatus.SETSTATUSColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: antenatal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return b, _spec
}
