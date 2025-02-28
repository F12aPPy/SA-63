// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/F12aPPy/app/ent/antenatal"
	"github.com/F12aPPy/app/ent/patient"
	"github.com/F12aPPy/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// PatientQuery is the builder for querying Patient entities.
type PatientQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Patient
	// eager-loading edges.
	withSETMOM *AntenatalQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (pq *PatientQuery) Where(ps ...predicate.Patient) *PatientQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *PatientQuery) Limit(limit int) *PatientQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *PatientQuery) Offset(offset int) *PatientQuery {
	pq.offset = &offset
	return pq
}

// Order adds an order step to the query.
func (pq *PatientQuery) Order(o ...OrderFunc) *PatientQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QuerySETMOM chains the current query on the SETMOM edge.
func (pq *PatientQuery) QuerySETMOM() *AntenatalQuery {
	query := &AntenatalQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(patient.Table, patient.FieldID, pq.sqlQuery()),
			sqlgraph.To(antenatal.Table, antenatal.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, patient.SETMOMTable, patient.SETMOMColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Patient entity in the query. Returns *NotFoundError when no patient was found.
func (pq *PatientQuery) First(ctx context.Context) (*Patient, error) {
	pas, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(pas) == 0 {
		return nil, &NotFoundError{patient.Label}
	}
	return pas[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PatientQuery) FirstX(ctx context.Context) *Patient {
	pa, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return pa
}

// FirstID returns the first Patient id in the query. Returns *NotFoundError when no id was found.
func (pq *PatientQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{patient.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (pq *PatientQuery) FirstXID(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Patient entity in the query, returns an error if not exactly one entity was returned.
func (pq *PatientQuery) Only(ctx context.Context) (*Patient, error) {
	pas, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(pas) {
	case 1:
		return pas[0], nil
	case 0:
		return nil, &NotFoundError{patient.Label}
	default:
		return nil, &NotSingularError{patient.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PatientQuery) OnlyX(ctx context.Context) *Patient {
	pa, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return pa
}

// OnlyID returns the only Patient id in the query, returns an error if not exactly one id was returned.
func (pq *PatientQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{patient.Label}
	default:
		err = &NotSingularError{patient.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PatientQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Patients.
func (pq *PatientQuery) All(ctx context.Context) ([]*Patient, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *PatientQuery) AllX(ctx context.Context) []*Patient {
	pas, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return pas
}

// IDs executes the query and returns a list of Patient ids.
func (pq *PatientQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pq.Select(patient.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PatientQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PatientQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PatientQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PatientQuery) Exist(ctx context.Context) (bool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PatientQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PatientQuery) Clone() *PatientQuery {
	return &PatientQuery{
		config:     pq.config,
		limit:      pq.limit,
		offset:     pq.offset,
		order:      append([]OrderFunc{}, pq.order...),
		unique:     append([]string{}, pq.unique...),
		predicates: append([]predicate.Patient{}, pq.predicates...),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

//  WithSETMOM tells the query-builder to eager-loads the nodes that are connected to
// the "SETMOM" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PatientQuery) WithSETMOM(opts ...func(*AntenatalQuery)) *PatientQuery {
	query := &AntenatalQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withSETMOM = query
	return pq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PatientName string `json:"patient_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Patient.Query().
//		GroupBy(patient.FieldPatientName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pq *PatientQuery) GroupBy(field string, fields ...string) *PatientGroupBy {
	group := &PatientGroupBy{config: pq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		PatientName string `json:"patient_name,omitempty"`
//	}
//
//	client.Patient.Query().
//		Select(patient.FieldPatientName).
//		Scan(ctx, &v)
//
func (pq *PatientQuery) Select(field string, fields ...string) *PatientSelect {
	selector := &PatientSelect{config: pq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(), nil
	}
	return selector
}

func (pq *PatientQuery) prepareQuery(ctx context.Context) error {
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PatientQuery) sqlAll(ctx context.Context) ([]*Patient, error) {
	var (
		nodes       = []*Patient{}
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withSETMOM != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Patient{config: pq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pq.withSETMOM; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Patient)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Antenatal(func(s *sql.Selector) {
			s.Where(sql.InValues(patient.SETMOMColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.patient_setmom
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "patient_setmom" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "patient_setmom" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.SETMOM = append(node.Edges.SETMOM, n)
		}
	}

	return nodes, nil
}

func (pq *PatientQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PatientQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (pq *PatientQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patient.Table,
			Columns: patient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PatientQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(patient.Table)
	selector := builder.Select(t1.Columns(patient.Columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(patient.Columns...)...)
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PatientGroupBy is the builder for group-by Patient entities.
type PatientGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PatientGroupBy) Aggregate(fns ...AggregateFunc) *PatientGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scan the result into the given value.
func (pgb *PatientGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pgb.path(ctx)
	if err != nil {
		return err
	}
	pgb.sql = query
	return pgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pgb *PatientGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PatientGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pgb *PatientGroupBy) StringsX(ctx context.Context) []string {
	v, err := pgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patient.Label}
	default:
		err = fmt.Errorf("ent: PatientGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pgb *PatientGroupBy) StringX(ctx context.Context) string {
	v, err := pgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PatientGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pgb *PatientGroupBy) IntsX(ctx context.Context) []int {
	v, err := pgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patient.Label}
	default:
		err = fmt.Errorf("ent: PatientGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pgb *PatientGroupBy) IntX(ctx context.Context) int {
	v, err := pgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PatientGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pgb *PatientGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patient.Label}
	default:
		err = fmt.Errorf("ent: PatientGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pgb *PatientGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PatientGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pgb *PatientGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patient.Label}
	default:
		err = fmt.Errorf("ent: PatientGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pgb *PatientGroupBy) BoolX(ctx context.Context) bool {
	v, err := pgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pgb *PatientGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pgb.sqlQuery().Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *PatientGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql
	columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
	columns = append(columns, pgb.fields...)
	for _, fn := range pgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(pgb.fields...)
}

// PatientSelect is the builder for select fields of Patient entities.
type PatientSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ps *PatientSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ps.path(ctx)
	if err != nil {
		return err
	}
	ps.sql = query
	return ps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ps *PatientSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ps *PatientSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PatientSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ps *PatientSelect) StringsX(ctx context.Context) []string {
	v, err := ps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ps *PatientSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patient.Label}
	default:
		err = fmt.Errorf("ent: PatientSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ps *PatientSelect) StringX(ctx context.Context) string {
	v, err := ps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ps *PatientSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PatientSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ps *PatientSelect) IntsX(ctx context.Context) []int {
	v, err := ps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ps *PatientSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patient.Label}
	default:
		err = fmt.Errorf("ent: PatientSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ps *PatientSelect) IntX(ctx context.Context) int {
	v, err := ps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ps *PatientSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PatientSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ps *PatientSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ps *PatientSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patient.Label}
	default:
		err = fmt.Errorf("ent: PatientSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ps *PatientSelect) Float64X(ctx context.Context) float64 {
	v, err := ps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ps *PatientSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PatientSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ps *PatientSelect) BoolsX(ctx context.Context) []bool {
	v, err := ps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ps *PatientSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patient.Label}
	default:
		err = fmt.Errorf("ent: PatientSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ps *PatientSelect) BoolX(ctx context.Context) bool {
	v, err := ps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ps *PatientSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ps.sqlQuery().Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ps *PatientSelect) sqlQuery() sql.Querier {
	selector := ps.sql
	selector.Select(selector.Columns(ps.fields...)...)
	return selector
}
