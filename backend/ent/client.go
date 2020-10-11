// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/F12aPPy/app/ent/migrate"

	"github.com/F12aPPy/app/ent/antenatal"
	"github.com/F12aPPy/app/ent/babystatus"
	"github.com/F12aPPy/app/ent/patient"
	"github.com/F12aPPy/app/ent/user"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Antenatal is the client for interacting with the Antenatal builders.
	Antenatal *AntenatalClient
	// Babystatus is the client for interacting with the Babystatus builders.
	Babystatus *BabystatusClient
	// Patient is the client for interacting with the Patient builders.
	Patient *PatientClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Antenatal = NewAntenatalClient(c.config)
	c.Babystatus = NewBabystatusClient(c.config)
	c.Patient = NewPatientClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Antenatal:  NewAntenatalClient(cfg),
		Babystatus: NewBabystatusClient(cfg),
		Patient:    NewPatientClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:     cfg,
		Antenatal:  NewAntenatalClient(cfg),
		Babystatus: NewBabystatusClient(cfg),
		Patient:    NewPatientClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Antenatal.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Antenatal.Use(hooks...)
	c.Babystatus.Use(hooks...)
	c.Patient.Use(hooks...)
	c.User.Use(hooks...)
}

// AntenatalClient is a client for the Antenatal schema.
type AntenatalClient struct {
	config
}

// NewAntenatalClient returns a client for the Antenatal from the given config.
func NewAntenatalClient(c config) *AntenatalClient {
	return &AntenatalClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `antenatal.Hooks(f(g(h())))`.
func (c *AntenatalClient) Use(hooks ...Hook) {
	c.hooks.Antenatal = append(c.hooks.Antenatal, hooks...)
}

// Create returns a create builder for Antenatal.
func (c *AntenatalClient) Create() *AntenatalCreate {
	mutation := newAntenatalMutation(c.config, OpCreate)
	return &AntenatalCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Antenatal.
func (c *AntenatalClient) Update() *AntenatalUpdate {
	mutation := newAntenatalMutation(c.config, OpUpdate)
	return &AntenatalUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AntenatalClient) UpdateOne(a *Antenatal) *AntenatalUpdateOne {
	mutation := newAntenatalMutation(c.config, OpUpdateOne, withAntenatal(a))
	return &AntenatalUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AntenatalClient) UpdateOneID(id int) *AntenatalUpdateOne {
	mutation := newAntenatalMutation(c.config, OpUpdateOne, withAntenatalID(id))
	return &AntenatalUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Antenatal.
func (c *AntenatalClient) Delete() *AntenatalDelete {
	mutation := newAntenatalMutation(c.config, OpDelete)
	return &AntenatalDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AntenatalClient) DeleteOne(a *Antenatal) *AntenatalDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AntenatalClient) DeleteOneID(id int) *AntenatalDeleteOne {
	builder := c.Delete().Where(antenatal.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AntenatalDeleteOne{builder}
}

// Create returns a query builder for Antenatal.
func (c *AntenatalClient) Query() *AntenatalQuery {
	return &AntenatalQuery{config: c.config}
}

// Get returns a Antenatal entity by its id.
func (c *AntenatalClient) Get(ctx context.Context, id int) (*Antenatal, error) {
	return c.Query().Where(antenatal.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AntenatalClient) GetX(ctx context.Context, id int) *Antenatal {
	a, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return a
}

// QueryGETMOM queries the GETMOM edge of a Antenatal.
func (c *AntenatalClient) QueryGETMOM(a *Antenatal) *PatientQuery {
	query := &PatientQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(antenatal.Table, antenatal.FieldID, id),
			sqlgraph.To(patient.Table, patient.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, antenatal.GETMOMTable, antenatal.GETMOMColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTAKECARE queries the TAKECARE edge of a Antenatal.
func (c *AntenatalClient) QueryTAKECARE(a *Antenatal) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(antenatal.Table, antenatal.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, antenatal.TAKECARETable, antenatal.TAKECAREColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGETSTATUS queries the GETSTATUS edge of a Antenatal.
func (c *AntenatalClient) QueryGETSTATUS(a *Antenatal) *BabystatusQuery {
	query := &BabystatusQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(antenatal.Table, antenatal.FieldID, id),
			sqlgraph.To(babystatus.Table, babystatus.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, antenatal.GETSTATUSTable, antenatal.GETSTATUSColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AntenatalClient) Hooks() []Hook {
	return c.hooks.Antenatal
}

// BabystatusClient is a client for the Babystatus schema.
type BabystatusClient struct {
	config
}

// NewBabystatusClient returns a client for the Babystatus from the given config.
func NewBabystatusClient(c config) *BabystatusClient {
	return &BabystatusClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `babystatus.Hooks(f(g(h())))`.
func (c *BabystatusClient) Use(hooks ...Hook) {
	c.hooks.Babystatus = append(c.hooks.Babystatus, hooks...)
}

// Create returns a create builder for Babystatus.
func (c *BabystatusClient) Create() *BabystatusCreate {
	mutation := newBabystatusMutation(c.config, OpCreate)
	return &BabystatusCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Babystatus.
func (c *BabystatusClient) Update() *BabystatusUpdate {
	mutation := newBabystatusMutation(c.config, OpUpdate)
	return &BabystatusUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BabystatusClient) UpdateOne(b *Babystatus) *BabystatusUpdateOne {
	mutation := newBabystatusMutation(c.config, OpUpdateOne, withBabystatus(b))
	return &BabystatusUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BabystatusClient) UpdateOneID(id int) *BabystatusUpdateOne {
	mutation := newBabystatusMutation(c.config, OpUpdateOne, withBabystatusID(id))
	return &BabystatusUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Babystatus.
func (c *BabystatusClient) Delete() *BabystatusDelete {
	mutation := newBabystatusMutation(c.config, OpDelete)
	return &BabystatusDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BabystatusClient) DeleteOne(b *Babystatus) *BabystatusDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BabystatusClient) DeleteOneID(id int) *BabystatusDeleteOne {
	builder := c.Delete().Where(babystatus.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BabystatusDeleteOne{builder}
}

// Create returns a query builder for Babystatus.
func (c *BabystatusClient) Query() *BabystatusQuery {
	return &BabystatusQuery{config: c.config}
}

// Get returns a Babystatus entity by its id.
func (c *BabystatusClient) Get(ctx context.Context, id int) (*Babystatus, error) {
	return c.Query().Where(babystatus.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BabystatusClient) GetX(ctx context.Context, id int) *Babystatus {
	b, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return b
}

// QuerySETSTATUS queries the SETSTATUS edge of a Babystatus.
func (c *BabystatusClient) QuerySETSTATUS(b *Babystatus) *AntenatalQuery {
	query := &AntenatalQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(babystatus.Table, babystatus.FieldID, id),
			sqlgraph.To(antenatal.Table, antenatal.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, babystatus.SETSTATUSTable, babystatus.SETSTATUSColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BabystatusClient) Hooks() []Hook {
	return c.hooks.Babystatus
}

// PatientClient is a client for the Patient schema.
type PatientClient struct {
	config
}

// NewPatientClient returns a client for the Patient from the given config.
func NewPatientClient(c config) *PatientClient {
	return &PatientClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `patient.Hooks(f(g(h())))`.
func (c *PatientClient) Use(hooks ...Hook) {
	c.hooks.Patient = append(c.hooks.Patient, hooks...)
}

// Create returns a create builder for Patient.
func (c *PatientClient) Create() *PatientCreate {
	mutation := newPatientMutation(c.config, OpCreate)
	return &PatientCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Patient.
func (c *PatientClient) Update() *PatientUpdate {
	mutation := newPatientMutation(c.config, OpUpdate)
	return &PatientUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PatientClient) UpdateOne(pa *Patient) *PatientUpdateOne {
	mutation := newPatientMutation(c.config, OpUpdateOne, withPatient(pa))
	return &PatientUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PatientClient) UpdateOneID(id int) *PatientUpdateOne {
	mutation := newPatientMutation(c.config, OpUpdateOne, withPatientID(id))
	return &PatientUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Patient.
func (c *PatientClient) Delete() *PatientDelete {
	mutation := newPatientMutation(c.config, OpDelete)
	return &PatientDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PatientClient) DeleteOne(pa *Patient) *PatientDeleteOne {
	return c.DeleteOneID(pa.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PatientClient) DeleteOneID(id int) *PatientDeleteOne {
	builder := c.Delete().Where(patient.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PatientDeleteOne{builder}
}

// Create returns a query builder for Patient.
func (c *PatientClient) Query() *PatientQuery {
	return &PatientQuery{config: c.config}
}

// Get returns a Patient entity by its id.
func (c *PatientClient) Get(ctx context.Context, id int) (*Patient, error) {
	return c.Query().Where(patient.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PatientClient) GetX(ctx context.Context, id int) *Patient {
	pa, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pa
}

// QuerySETMOM queries the SETMOM edge of a Patient.
func (c *PatientClient) QuerySETMOM(pa *Patient) *AntenatalQuery {
	query := &AntenatalQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(patient.Table, patient.FieldID, id),
			sqlgraph.To(antenatal.Table, antenatal.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, patient.SETMOMTable, patient.SETMOMColumn),
		)
		fromV = sqlgraph.Neighbors(pa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PatientClient) Hooks() []Hook {
	return c.hooks.Patient
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryCARE queries the CARE edge of a User.
func (c *UserClient) QueryCARE(u *User) *AntenatalQuery {
	query := &AntenatalQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(antenatal.Table, antenatal.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.CARETable, user.CAREColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
