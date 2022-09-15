// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/sky0621/cv-admin/src/interfaces/ent/ent/migrate"

	"github.com/sky0621/cv-admin/src/interfaces/ent/ent/careerskill"
	"github.com/sky0621/cv-admin/src/interfaces/ent/ent/careerskillgroup"
	"github.com/sky0621/cv-admin/src/interfaces/ent/ent/careertask"
	"github.com/sky0621/cv-admin/src/interfaces/ent/ent/user"
	"github.com/sky0621/cv-admin/src/interfaces/ent/ent/useractivity"
	"github.com/sky0621/cv-admin/src/interfaces/ent/ent/usercareer"
	"github.com/sky0621/cv-admin/src/interfaces/ent/ent/usercareergroup"
	"github.com/sky0621/cv-admin/src/interfaces/ent/ent/usernote"
	"github.com/sky0621/cv-admin/src/interfaces/ent/ent/usernoteitem"
	"github.com/sky0621/cv-admin/src/interfaces/ent/ent/userqualification"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// CareerSkill is the client for interacting with the CareerSkill builders.
	CareerSkill *CareerSkillClient
	// CareerSkillGroup is the client for interacting with the CareerSkillGroup builders.
	CareerSkillGroup *CareerSkillGroupClient
	// CareerTask is the client for interacting with the CareerTask builders.
	CareerTask *CareerTaskClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserActivity is the client for interacting with the UserActivity builders.
	UserActivity *UserActivityClient
	// UserCareer is the client for interacting with the UserCareer builders.
	UserCareer *UserCareerClient
	// UserCareerGroup is the client for interacting with the UserCareerGroup builders.
	UserCareerGroup *UserCareerGroupClient
	// UserNote is the client for interacting with the UserNote builders.
	UserNote *UserNoteClient
	// UserNoteItem is the client for interacting with the UserNoteItem builders.
	UserNoteItem *UserNoteItemClient
	// UserQualification is the client for interacting with the UserQualification builders.
	UserQualification *UserQualificationClient
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
	c.CareerSkill = NewCareerSkillClient(c.config)
	c.CareerSkillGroup = NewCareerSkillGroupClient(c.config)
	c.CareerTask = NewCareerTaskClient(c.config)
	c.User = NewUserClient(c.config)
	c.UserActivity = NewUserActivityClient(c.config)
	c.UserCareer = NewUserCareerClient(c.config)
	c.UserCareerGroup = NewUserCareerGroupClient(c.config)
	c.UserNote = NewUserNoteClient(c.config)
	c.UserNoteItem = NewUserNoteItemClient(c.config)
	c.UserQualification = NewUserQualificationClient(c.config)
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
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:               ctx,
		config:            cfg,
		CareerSkill:       NewCareerSkillClient(cfg),
		CareerSkillGroup:  NewCareerSkillGroupClient(cfg),
		CareerTask:        NewCareerTaskClient(cfg),
		User:              NewUserClient(cfg),
		UserActivity:      NewUserActivityClient(cfg),
		UserCareer:        NewUserCareerClient(cfg),
		UserCareerGroup:   NewUserCareerGroupClient(cfg),
		UserNote:          NewUserNoteClient(cfg),
		UserNoteItem:      NewUserNoteItemClient(cfg),
		UserQualification: NewUserQualificationClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:               ctx,
		config:            cfg,
		CareerSkill:       NewCareerSkillClient(cfg),
		CareerSkillGroup:  NewCareerSkillGroupClient(cfg),
		CareerTask:        NewCareerTaskClient(cfg),
		User:              NewUserClient(cfg),
		UserActivity:      NewUserActivityClient(cfg),
		UserCareer:        NewUserCareerClient(cfg),
		UserCareerGroup:   NewUserCareerGroupClient(cfg),
		UserNote:          NewUserNoteClient(cfg),
		UserNoteItem:      NewUserNoteItemClient(cfg),
		UserQualification: NewUserQualificationClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		CareerSkill.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
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
	c.CareerSkill.Use(hooks...)
	c.CareerSkillGroup.Use(hooks...)
	c.CareerTask.Use(hooks...)
	c.User.Use(hooks...)
	c.UserActivity.Use(hooks...)
	c.UserCareer.Use(hooks...)
	c.UserCareerGroup.Use(hooks...)
	c.UserNote.Use(hooks...)
	c.UserNoteItem.Use(hooks...)
	c.UserQualification.Use(hooks...)
}

// CareerSkillClient is a client for the CareerSkill schema.
type CareerSkillClient struct {
	config
}

// NewCareerSkillClient returns a client for the CareerSkill from the given config.
func NewCareerSkillClient(c config) *CareerSkillClient {
	return &CareerSkillClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `careerskill.Hooks(f(g(h())))`.
func (c *CareerSkillClient) Use(hooks ...Hook) {
	c.hooks.CareerSkill = append(c.hooks.CareerSkill, hooks...)
}

// Create returns a builder for creating a CareerSkill entity.
func (c *CareerSkillClient) Create() *CareerSkillCreate {
	mutation := newCareerSkillMutation(c.config, OpCreate)
	return &CareerSkillCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CareerSkill entities.
func (c *CareerSkillClient) CreateBulk(builders ...*CareerSkillCreate) *CareerSkillCreateBulk {
	return &CareerSkillCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CareerSkill.
func (c *CareerSkillClient) Update() *CareerSkillUpdate {
	mutation := newCareerSkillMutation(c.config, OpUpdate)
	return &CareerSkillUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CareerSkillClient) UpdateOne(cs *CareerSkill) *CareerSkillUpdateOne {
	mutation := newCareerSkillMutation(c.config, OpUpdateOne, withCareerSkill(cs))
	return &CareerSkillUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CareerSkillClient) UpdateOneID(id int) *CareerSkillUpdateOne {
	mutation := newCareerSkillMutation(c.config, OpUpdateOne, withCareerSkillID(id))
	return &CareerSkillUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CareerSkill.
func (c *CareerSkillClient) Delete() *CareerSkillDelete {
	mutation := newCareerSkillMutation(c.config, OpDelete)
	return &CareerSkillDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CareerSkillClient) DeleteOne(cs *CareerSkill) *CareerSkillDeleteOne {
	return c.DeleteOneID(cs.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CareerSkillClient) DeleteOneID(id int) *CareerSkillDeleteOne {
	builder := c.Delete().Where(careerskill.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CareerSkillDeleteOne{builder}
}

// Query returns a query builder for CareerSkill.
func (c *CareerSkillClient) Query() *CareerSkillQuery {
	return &CareerSkillQuery{
		config: c.config,
	}
}

// Get returns a CareerSkill entity by its id.
func (c *CareerSkillClient) Get(ctx context.Context, id int) (*CareerSkill, error) {
	return c.Query().Where(careerskill.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CareerSkillClient) GetX(ctx context.Context, id int) *CareerSkill {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CareerSkillClient) Hooks() []Hook {
	return c.hooks.CareerSkill
}

// CareerSkillGroupClient is a client for the CareerSkillGroup schema.
type CareerSkillGroupClient struct {
	config
}

// NewCareerSkillGroupClient returns a client for the CareerSkillGroup from the given config.
func NewCareerSkillGroupClient(c config) *CareerSkillGroupClient {
	return &CareerSkillGroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `careerskillgroup.Hooks(f(g(h())))`.
func (c *CareerSkillGroupClient) Use(hooks ...Hook) {
	c.hooks.CareerSkillGroup = append(c.hooks.CareerSkillGroup, hooks...)
}

// Create returns a builder for creating a CareerSkillGroup entity.
func (c *CareerSkillGroupClient) Create() *CareerSkillGroupCreate {
	mutation := newCareerSkillGroupMutation(c.config, OpCreate)
	return &CareerSkillGroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CareerSkillGroup entities.
func (c *CareerSkillGroupClient) CreateBulk(builders ...*CareerSkillGroupCreate) *CareerSkillGroupCreateBulk {
	return &CareerSkillGroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CareerSkillGroup.
func (c *CareerSkillGroupClient) Update() *CareerSkillGroupUpdate {
	mutation := newCareerSkillGroupMutation(c.config, OpUpdate)
	return &CareerSkillGroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CareerSkillGroupClient) UpdateOne(csg *CareerSkillGroup) *CareerSkillGroupUpdateOne {
	mutation := newCareerSkillGroupMutation(c.config, OpUpdateOne, withCareerSkillGroup(csg))
	return &CareerSkillGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CareerSkillGroupClient) UpdateOneID(id int) *CareerSkillGroupUpdateOne {
	mutation := newCareerSkillGroupMutation(c.config, OpUpdateOne, withCareerSkillGroupID(id))
	return &CareerSkillGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CareerSkillGroup.
func (c *CareerSkillGroupClient) Delete() *CareerSkillGroupDelete {
	mutation := newCareerSkillGroupMutation(c.config, OpDelete)
	return &CareerSkillGroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CareerSkillGroupClient) DeleteOne(csg *CareerSkillGroup) *CareerSkillGroupDeleteOne {
	return c.DeleteOneID(csg.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CareerSkillGroupClient) DeleteOneID(id int) *CareerSkillGroupDeleteOne {
	builder := c.Delete().Where(careerskillgroup.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CareerSkillGroupDeleteOne{builder}
}

// Query returns a query builder for CareerSkillGroup.
func (c *CareerSkillGroupClient) Query() *CareerSkillGroupQuery {
	return &CareerSkillGroupQuery{
		config: c.config,
	}
}

// Get returns a CareerSkillGroup entity by its id.
func (c *CareerSkillGroupClient) Get(ctx context.Context, id int) (*CareerSkillGroup, error) {
	return c.Query().Where(careerskillgroup.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CareerSkillGroupClient) GetX(ctx context.Context, id int) *CareerSkillGroup {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CareerSkillGroupClient) Hooks() []Hook {
	return c.hooks.CareerSkillGroup
}

// CareerTaskClient is a client for the CareerTask schema.
type CareerTaskClient struct {
	config
}

// NewCareerTaskClient returns a client for the CareerTask from the given config.
func NewCareerTaskClient(c config) *CareerTaskClient {
	return &CareerTaskClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `careertask.Hooks(f(g(h())))`.
func (c *CareerTaskClient) Use(hooks ...Hook) {
	c.hooks.CareerTask = append(c.hooks.CareerTask, hooks...)
}

// Create returns a builder for creating a CareerTask entity.
func (c *CareerTaskClient) Create() *CareerTaskCreate {
	mutation := newCareerTaskMutation(c.config, OpCreate)
	return &CareerTaskCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CareerTask entities.
func (c *CareerTaskClient) CreateBulk(builders ...*CareerTaskCreate) *CareerTaskCreateBulk {
	return &CareerTaskCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CareerTask.
func (c *CareerTaskClient) Update() *CareerTaskUpdate {
	mutation := newCareerTaskMutation(c.config, OpUpdate)
	return &CareerTaskUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CareerTaskClient) UpdateOne(ct *CareerTask) *CareerTaskUpdateOne {
	mutation := newCareerTaskMutation(c.config, OpUpdateOne, withCareerTask(ct))
	return &CareerTaskUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CareerTaskClient) UpdateOneID(id int) *CareerTaskUpdateOne {
	mutation := newCareerTaskMutation(c.config, OpUpdateOne, withCareerTaskID(id))
	return &CareerTaskUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CareerTask.
func (c *CareerTaskClient) Delete() *CareerTaskDelete {
	mutation := newCareerTaskMutation(c.config, OpDelete)
	return &CareerTaskDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CareerTaskClient) DeleteOne(ct *CareerTask) *CareerTaskDeleteOne {
	return c.DeleteOneID(ct.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CareerTaskClient) DeleteOneID(id int) *CareerTaskDeleteOne {
	builder := c.Delete().Where(careertask.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CareerTaskDeleteOne{builder}
}

// Query returns a query builder for CareerTask.
func (c *CareerTaskClient) Query() *CareerTaskQuery {
	return &CareerTaskQuery{
		config: c.config,
	}
}

// Get returns a CareerTask entity by its id.
func (c *CareerTaskClient) Get(ctx context.Context, id int) (*CareerTask, error) {
	return c.Query().Where(careertask.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CareerTaskClient) GetX(ctx context.Context, id int) *CareerTask {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CareerTaskClient) Hooks() []Hook {
	return c.hooks.CareerTask
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

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
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

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// UserActivityClient is a client for the UserActivity schema.
type UserActivityClient struct {
	config
}

// NewUserActivityClient returns a client for the UserActivity from the given config.
func NewUserActivityClient(c config) *UserActivityClient {
	return &UserActivityClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `useractivity.Hooks(f(g(h())))`.
func (c *UserActivityClient) Use(hooks ...Hook) {
	c.hooks.UserActivity = append(c.hooks.UserActivity, hooks...)
}

// Create returns a builder for creating a UserActivity entity.
func (c *UserActivityClient) Create() *UserActivityCreate {
	mutation := newUserActivityMutation(c.config, OpCreate)
	return &UserActivityCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserActivity entities.
func (c *UserActivityClient) CreateBulk(builders ...*UserActivityCreate) *UserActivityCreateBulk {
	return &UserActivityCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserActivity.
func (c *UserActivityClient) Update() *UserActivityUpdate {
	mutation := newUserActivityMutation(c.config, OpUpdate)
	return &UserActivityUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserActivityClient) UpdateOne(ua *UserActivity) *UserActivityUpdateOne {
	mutation := newUserActivityMutation(c.config, OpUpdateOne, withUserActivity(ua))
	return &UserActivityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserActivityClient) UpdateOneID(id int) *UserActivityUpdateOne {
	mutation := newUserActivityMutation(c.config, OpUpdateOne, withUserActivityID(id))
	return &UserActivityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserActivity.
func (c *UserActivityClient) Delete() *UserActivityDelete {
	mutation := newUserActivityMutation(c.config, OpDelete)
	return &UserActivityDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserActivityClient) DeleteOne(ua *UserActivity) *UserActivityDeleteOne {
	return c.DeleteOneID(ua.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserActivityClient) DeleteOneID(id int) *UserActivityDeleteOne {
	builder := c.Delete().Where(useractivity.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserActivityDeleteOne{builder}
}

// Query returns a query builder for UserActivity.
func (c *UserActivityClient) Query() *UserActivityQuery {
	return &UserActivityQuery{
		config: c.config,
	}
}

// Get returns a UserActivity entity by its id.
func (c *UserActivityClient) Get(ctx context.Context, id int) (*UserActivity, error) {
	return c.Query().Where(useractivity.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserActivityClient) GetX(ctx context.Context, id int) *UserActivity {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserActivityClient) Hooks() []Hook {
	return c.hooks.UserActivity
}

// UserCareerClient is a client for the UserCareer schema.
type UserCareerClient struct {
	config
}

// NewUserCareerClient returns a client for the UserCareer from the given config.
func NewUserCareerClient(c config) *UserCareerClient {
	return &UserCareerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `usercareer.Hooks(f(g(h())))`.
func (c *UserCareerClient) Use(hooks ...Hook) {
	c.hooks.UserCareer = append(c.hooks.UserCareer, hooks...)
}

// Create returns a builder for creating a UserCareer entity.
func (c *UserCareerClient) Create() *UserCareerCreate {
	mutation := newUserCareerMutation(c.config, OpCreate)
	return &UserCareerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserCareer entities.
func (c *UserCareerClient) CreateBulk(builders ...*UserCareerCreate) *UserCareerCreateBulk {
	return &UserCareerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserCareer.
func (c *UserCareerClient) Update() *UserCareerUpdate {
	mutation := newUserCareerMutation(c.config, OpUpdate)
	return &UserCareerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserCareerClient) UpdateOne(uc *UserCareer) *UserCareerUpdateOne {
	mutation := newUserCareerMutation(c.config, OpUpdateOne, withUserCareer(uc))
	return &UserCareerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserCareerClient) UpdateOneID(id int) *UserCareerUpdateOne {
	mutation := newUserCareerMutation(c.config, OpUpdateOne, withUserCareerID(id))
	return &UserCareerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserCareer.
func (c *UserCareerClient) Delete() *UserCareerDelete {
	mutation := newUserCareerMutation(c.config, OpDelete)
	return &UserCareerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserCareerClient) DeleteOne(uc *UserCareer) *UserCareerDeleteOne {
	return c.DeleteOneID(uc.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserCareerClient) DeleteOneID(id int) *UserCareerDeleteOne {
	builder := c.Delete().Where(usercareer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserCareerDeleteOne{builder}
}

// Query returns a query builder for UserCareer.
func (c *UserCareerClient) Query() *UserCareerQuery {
	return &UserCareerQuery{
		config: c.config,
	}
}

// Get returns a UserCareer entity by its id.
func (c *UserCareerClient) Get(ctx context.Context, id int) (*UserCareer, error) {
	return c.Query().Where(usercareer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserCareerClient) GetX(ctx context.Context, id int) *UserCareer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserCareerClient) Hooks() []Hook {
	return c.hooks.UserCareer
}

// UserCareerGroupClient is a client for the UserCareerGroup schema.
type UserCareerGroupClient struct {
	config
}

// NewUserCareerGroupClient returns a client for the UserCareerGroup from the given config.
func NewUserCareerGroupClient(c config) *UserCareerGroupClient {
	return &UserCareerGroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `usercareergroup.Hooks(f(g(h())))`.
func (c *UserCareerGroupClient) Use(hooks ...Hook) {
	c.hooks.UserCareerGroup = append(c.hooks.UserCareerGroup, hooks...)
}

// Create returns a builder for creating a UserCareerGroup entity.
func (c *UserCareerGroupClient) Create() *UserCareerGroupCreate {
	mutation := newUserCareerGroupMutation(c.config, OpCreate)
	return &UserCareerGroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserCareerGroup entities.
func (c *UserCareerGroupClient) CreateBulk(builders ...*UserCareerGroupCreate) *UserCareerGroupCreateBulk {
	return &UserCareerGroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserCareerGroup.
func (c *UserCareerGroupClient) Update() *UserCareerGroupUpdate {
	mutation := newUserCareerGroupMutation(c.config, OpUpdate)
	return &UserCareerGroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserCareerGroupClient) UpdateOne(ucg *UserCareerGroup) *UserCareerGroupUpdateOne {
	mutation := newUserCareerGroupMutation(c.config, OpUpdateOne, withUserCareerGroup(ucg))
	return &UserCareerGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserCareerGroupClient) UpdateOneID(id int) *UserCareerGroupUpdateOne {
	mutation := newUserCareerGroupMutation(c.config, OpUpdateOne, withUserCareerGroupID(id))
	return &UserCareerGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserCareerGroup.
func (c *UserCareerGroupClient) Delete() *UserCareerGroupDelete {
	mutation := newUserCareerGroupMutation(c.config, OpDelete)
	return &UserCareerGroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserCareerGroupClient) DeleteOne(ucg *UserCareerGroup) *UserCareerGroupDeleteOne {
	return c.DeleteOneID(ucg.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserCareerGroupClient) DeleteOneID(id int) *UserCareerGroupDeleteOne {
	builder := c.Delete().Where(usercareergroup.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserCareerGroupDeleteOne{builder}
}

// Query returns a query builder for UserCareerGroup.
func (c *UserCareerGroupClient) Query() *UserCareerGroupQuery {
	return &UserCareerGroupQuery{
		config: c.config,
	}
}

// Get returns a UserCareerGroup entity by its id.
func (c *UserCareerGroupClient) Get(ctx context.Context, id int) (*UserCareerGroup, error) {
	return c.Query().Where(usercareergroup.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserCareerGroupClient) GetX(ctx context.Context, id int) *UserCareerGroup {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserCareerGroupClient) Hooks() []Hook {
	return c.hooks.UserCareerGroup
}

// UserNoteClient is a client for the UserNote schema.
type UserNoteClient struct {
	config
}

// NewUserNoteClient returns a client for the UserNote from the given config.
func NewUserNoteClient(c config) *UserNoteClient {
	return &UserNoteClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `usernote.Hooks(f(g(h())))`.
func (c *UserNoteClient) Use(hooks ...Hook) {
	c.hooks.UserNote = append(c.hooks.UserNote, hooks...)
}

// Create returns a builder for creating a UserNote entity.
func (c *UserNoteClient) Create() *UserNoteCreate {
	mutation := newUserNoteMutation(c.config, OpCreate)
	return &UserNoteCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserNote entities.
func (c *UserNoteClient) CreateBulk(builders ...*UserNoteCreate) *UserNoteCreateBulk {
	return &UserNoteCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserNote.
func (c *UserNoteClient) Update() *UserNoteUpdate {
	mutation := newUserNoteMutation(c.config, OpUpdate)
	return &UserNoteUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserNoteClient) UpdateOne(un *UserNote) *UserNoteUpdateOne {
	mutation := newUserNoteMutation(c.config, OpUpdateOne, withUserNote(un))
	return &UserNoteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserNoteClient) UpdateOneID(id int) *UserNoteUpdateOne {
	mutation := newUserNoteMutation(c.config, OpUpdateOne, withUserNoteID(id))
	return &UserNoteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserNote.
func (c *UserNoteClient) Delete() *UserNoteDelete {
	mutation := newUserNoteMutation(c.config, OpDelete)
	return &UserNoteDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserNoteClient) DeleteOne(un *UserNote) *UserNoteDeleteOne {
	return c.DeleteOneID(un.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserNoteClient) DeleteOneID(id int) *UserNoteDeleteOne {
	builder := c.Delete().Where(usernote.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserNoteDeleteOne{builder}
}

// Query returns a query builder for UserNote.
func (c *UserNoteClient) Query() *UserNoteQuery {
	return &UserNoteQuery{
		config: c.config,
	}
}

// Get returns a UserNote entity by its id.
func (c *UserNoteClient) Get(ctx context.Context, id int) (*UserNote, error) {
	return c.Query().Where(usernote.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserNoteClient) GetX(ctx context.Context, id int) *UserNote {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserNoteClient) Hooks() []Hook {
	return c.hooks.UserNote
}

// UserNoteItemClient is a client for the UserNoteItem schema.
type UserNoteItemClient struct {
	config
}

// NewUserNoteItemClient returns a client for the UserNoteItem from the given config.
func NewUserNoteItemClient(c config) *UserNoteItemClient {
	return &UserNoteItemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `usernoteitem.Hooks(f(g(h())))`.
func (c *UserNoteItemClient) Use(hooks ...Hook) {
	c.hooks.UserNoteItem = append(c.hooks.UserNoteItem, hooks...)
}

// Create returns a builder for creating a UserNoteItem entity.
func (c *UserNoteItemClient) Create() *UserNoteItemCreate {
	mutation := newUserNoteItemMutation(c.config, OpCreate)
	return &UserNoteItemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserNoteItem entities.
func (c *UserNoteItemClient) CreateBulk(builders ...*UserNoteItemCreate) *UserNoteItemCreateBulk {
	return &UserNoteItemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserNoteItem.
func (c *UserNoteItemClient) Update() *UserNoteItemUpdate {
	mutation := newUserNoteItemMutation(c.config, OpUpdate)
	return &UserNoteItemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserNoteItemClient) UpdateOne(uni *UserNoteItem) *UserNoteItemUpdateOne {
	mutation := newUserNoteItemMutation(c.config, OpUpdateOne, withUserNoteItem(uni))
	return &UserNoteItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserNoteItemClient) UpdateOneID(id int) *UserNoteItemUpdateOne {
	mutation := newUserNoteItemMutation(c.config, OpUpdateOne, withUserNoteItemID(id))
	return &UserNoteItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserNoteItem.
func (c *UserNoteItemClient) Delete() *UserNoteItemDelete {
	mutation := newUserNoteItemMutation(c.config, OpDelete)
	return &UserNoteItemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserNoteItemClient) DeleteOne(uni *UserNoteItem) *UserNoteItemDeleteOne {
	return c.DeleteOneID(uni.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserNoteItemClient) DeleteOneID(id int) *UserNoteItemDeleteOne {
	builder := c.Delete().Where(usernoteitem.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserNoteItemDeleteOne{builder}
}

// Query returns a query builder for UserNoteItem.
func (c *UserNoteItemClient) Query() *UserNoteItemQuery {
	return &UserNoteItemQuery{
		config: c.config,
	}
}

// Get returns a UserNoteItem entity by its id.
func (c *UserNoteItemClient) Get(ctx context.Context, id int) (*UserNoteItem, error) {
	return c.Query().Where(usernoteitem.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserNoteItemClient) GetX(ctx context.Context, id int) *UserNoteItem {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserNoteItemClient) Hooks() []Hook {
	return c.hooks.UserNoteItem
}

// UserQualificationClient is a client for the UserQualification schema.
type UserQualificationClient struct {
	config
}

// NewUserQualificationClient returns a client for the UserQualification from the given config.
func NewUserQualificationClient(c config) *UserQualificationClient {
	return &UserQualificationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `userqualification.Hooks(f(g(h())))`.
func (c *UserQualificationClient) Use(hooks ...Hook) {
	c.hooks.UserQualification = append(c.hooks.UserQualification, hooks...)
}

// Create returns a builder for creating a UserQualification entity.
func (c *UserQualificationClient) Create() *UserQualificationCreate {
	mutation := newUserQualificationMutation(c.config, OpCreate)
	return &UserQualificationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserQualification entities.
func (c *UserQualificationClient) CreateBulk(builders ...*UserQualificationCreate) *UserQualificationCreateBulk {
	return &UserQualificationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserQualification.
func (c *UserQualificationClient) Update() *UserQualificationUpdate {
	mutation := newUserQualificationMutation(c.config, OpUpdate)
	return &UserQualificationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserQualificationClient) UpdateOne(uq *UserQualification) *UserQualificationUpdateOne {
	mutation := newUserQualificationMutation(c.config, OpUpdateOne, withUserQualification(uq))
	return &UserQualificationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserQualificationClient) UpdateOneID(id int) *UserQualificationUpdateOne {
	mutation := newUserQualificationMutation(c.config, OpUpdateOne, withUserQualificationID(id))
	return &UserQualificationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserQualification.
func (c *UserQualificationClient) Delete() *UserQualificationDelete {
	mutation := newUserQualificationMutation(c.config, OpDelete)
	return &UserQualificationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserQualificationClient) DeleteOne(uq *UserQualification) *UserQualificationDeleteOne {
	return c.DeleteOneID(uq.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserQualificationClient) DeleteOneID(id int) *UserQualificationDeleteOne {
	builder := c.Delete().Where(userqualification.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserQualificationDeleteOne{builder}
}

// Query returns a query builder for UserQualification.
func (c *UserQualificationClient) Query() *UserQualificationQuery {
	return &UserQualificationQuery{
		config: c.config,
	}
}

// Get returns a UserQualification entity by its id.
func (c *UserQualificationClient) Get(ctx context.Context, id int) (*UserQualification, error) {
	return c.Query().Where(userqualification.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserQualificationClient) GetX(ctx context.Context, id int) *UserQualification {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserQualificationClient) Hooks() []Hook {
	return c.hooks.UserQualification
}
