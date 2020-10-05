// Code generated by entc, DO NOT EDIT.

package ent

import (
	"baritone/bot/commands/guildconfig"
	"baritone/ent/guild"
	"baritone/ent/warning"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/diamondburned/arikawa/discord"
	"github.com/facebook/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeGuild   = "Guild"
	TypeUser    = "User"
	TypeWarning = "Warning"
)

// GuildMutation represents an operation that mutate the Guilds
// nodes in the graph.
type GuildMutation struct {
	config
	op              Op
	typ             string
	id              *discord.GuildID
	_config         *guildconfig.GuildConfig
	clearedFields   map[string]struct{}
	warnings        map[int]struct{}
	removedwarnings map[int]struct{}
	clearedwarnings bool
	done            bool
	oldValue        func(context.Context) (*Guild, error)
}

var _ ent.Mutation = (*GuildMutation)(nil)

// guildOption allows to manage the mutation configuration using functional options.
type guildOption func(*GuildMutation)

// newGuildMutation creates new mutation for $n.Name.
func newGuildMutation(c config, op Op, opts ...guildOption) *GuildMutation {
	m := &GuildMutation{
		config:        c,
		op:            op,
		typ:           TypeGuild,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withGuildID sets the id field of the mutation.
func withGuildID(id discord.GuildID) guildOption {
	return func(m *GuildMutation) {
		var (
			err   error
			once  sync.Once
			value *Guild
		)
		m.oldValue = func(ctx context.Context) (*Guild, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Guild.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withGuild sets the old Guild of the mutation.
func withGuild(node *Guild) guildOption {
	return func(m *GuildMutation) {
		m.oldValue = func(context.Context) (*Guild, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m GuildMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m GuildMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on Guild creation.
func (m *GuildMutation) SetID(id discord.GuildID) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *GuildMutation) ID() (id discord.GuildID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetConfig sets the config field.
func (m *GuildMutation) SetConfig(gc guildconfig.GuildConfig) {
	m._config = &gc
}

// Config returns the config value in the mutation.
func (m *GuildMutation) Config() (r guildconfig.GuildConfig, exists bool) {
	v := m._config
	if v == nil {
		return
	}
	return *v, true
}

// OldConfig returns the old config value of the Guild.
// If the Guild object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *GuildMutation) OldConfig(ctx context.Context) (v guildconfig.GuildConfig, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldConfig is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldConfig requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldConfig: %w", err)
	}
	return oldValue.Config, nil
}

// ResetConfig reset all changes of the "config" field.
func (m *GuildMutation) ResetConfig() {
	m._config = nil
}

// AddWarningIDs adds the warnings edge to Warning by ids.
func (m *GuildMutation) AddWarningIDs(ids ...int) {
	if m.warnings == nil {
		m.warnings = make(map[int]struct{})
	}
	for i := range ids {
		m.warnings[ids[i]] = struct{}{}
	}
}

// ClearWarnings clears the warnings edge to Warning.
func (m *GuildMutation) ClearWarnings() {
	m.clearedwarnings = true
}

// WarningsCleared returns if the edge warnings was cleared.
func (m *GuildMutation) WarningsCleared() bool {
	return m.clearedwarnings
}

// RemoveWarningIDs removes the warnings edge to Warning by ids.
func (m *GuildMutation) RemoveWarningIDs(ids ...int) {
	if m.removedwarnings == nil {
		m.removedwarnings = make(map[int]struct{})
	}
	for i := range ids {
		m.removedwarnings[ids[i]] = struct{}{}
	}
}

// RemovedWarnings returns the removed ids of warnings.
func (m *GuildMutation) RemovedWarningsIDs() (ids []int) {
	for id := range m.removedwarnings {
		ids = append(ids, id)
	}
	return
}

// WarningsIDs returns the warnings ids in the mutation.
func (m *GuildMutation) WarningsIDs() (ids []int) {
	for id := range m.warnings {
		ids = append(ids, id)
	}
	return
}

// ResetWarnings reset all changes of the "warnings" edge.
func (m *GuildMutation) ResetWarnings() {
	m.warnings = nil
	m.clearedwarnings = false
	m.removedwarnings = nil
}

// Op returns the operation name.
func (m *GuildMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Guild).
func (m *GuildMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *GuildMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m._config != nil {
		fields = append(fields, guild.FieldConfig)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *GuildMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case guild.FieldConfig:
		return m.Config()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *GuildMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case guild.FieldConfig:
		return m.OldConfig(ctx)
	}
	return nil, fmt.Errorf("unknown Guild field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *GuildMutation) SetField(name string, value ent.Value) error {
	switch name {
	case guild.FieldConfig:
		v, ok := value.(guildconfig.GuildConfig)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetConfig(v)
		return nil
	}
	return fmt.Errorf("unknown Guild field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *GuildMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *GuildMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *GuildMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Guild numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *GuildMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *GuildMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *GuildMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Guild nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *GuildMutation) ResetField(name string) error {
	switch name {
	case guild.FieldConfig:
		m.ResetConfig()
		return nil
	}
	return fmt.Errorf("unknown Guild field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *GuildMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.warnings != nil {
		edges = append(edges, guild.EdgeWarnings)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *GuildMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case guild.EdgeWarnings:
		ids := make([]ent.Value, 0, len(m.warnings))
		for id := range m.warnings {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *GuildMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedwarnings != nil {
		edges = append(edges, guild.EdgeWarnings)
	}
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *GuildMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case guild.EdgeWarnings:
		ids := make([]ent.Value, 0, len(m.removedwarnings))
		for id := range m.removedwarnings {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *GuildMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedwarnings {
		edges = append(edges, guild.EdgeWarnings)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *GuildMutation) EdgeCleared(name string) bool {
	switch name {
	case guild.EdgeWarnings:
		return m.clearedwarnings
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *GuildMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Guild unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *GuildMutation) ResetEdge(name string) error {
	switch name {
	case guild.EdgeWarnings:
		m.ResetWarnings()
		return nil
	}
	return fmt.Errorf("unknown Guild edge %s", name)
}

// UserMutation represents an operation that mutate the Users
// nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *discord.UserID
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows to manage the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for $n.Name.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the id field of the mutation.
func withUserID(id discord.UserID) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on User creation.
func (m *UserMutation) SetID(id discord.UserID) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *UserMutation) ID() (id discord.UserID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}

// WarningMutation represents an operation that mutate the Warnings
// nodes in the graph.
type WarningMutation struct {
	config
	op              Op
	typ             string
	id              *int
	reason          *string
	date            *time.Time
	clearedFields   map[string]struct{}
	user            *discord.UserID
	cleareduser     bool
	issuedBy        *discord.UserID
	clearedissuedBy bool
	guild           *discord.GuildID
	clearedguild    bool
	done            bool
	oldValue        func(context.Context) (*Warning, error)
}

var _ ent.Mutation = (*WarningMutation)(nil)

// warningOption allows to manage the mutation configuration using functional options.
type warningOption func(*WarningMutation)

// newWarningMutation creates new mutation for $n.Name.
func newWarningMutation(c config, op Op, opts ...warningOption) *WarningMutation {
	m := &WarningMutation{
		config:        c,
		op:            op,
		typ:           TypeWarning,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withWarningID sets the id field of the mutation.
func withWarningID(id int) warningOption {
	return func(m *WarningMutation) {
		var (
			err   error
			once  sync.Once
			value *Warning
		)
		m.oldValue = func(ctx context.Context) (*Warning, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Warning.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withWarning sets the old Warning of the mutation.
func withWarning(node *Warning) warningOption {
	return func(m *WarningMutation) {
		m.oldValue = func(context.Context) (*Warning, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m WarningMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m WarningMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *WarningMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetReason sets the reason field.
func (m *WarningMutation) SetReason(s string) {
	m.reason = &s
}

// Reason returns the reason value in the mutation.
func (m *WarningMutation) Reason() (r string, exists bool) {
	v := m.reason
	if v == nil {
		return
	}
	return *v, true
}

// OldReason returns the old reason value of the Warning.
// If the Warning object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *WarningMutation) OldReason(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldReason is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldReason requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldReason: %w", err)
	}
	return oldValue.Reason, nil
}

// ResetReason reset all changes of the "reason" field.
func (m *WarningMutation) ResetReason() {
	m.reason = nil
}

// SetDate sets the date field.
func (m *WarningMutation) SetDate(t time.Time) {
	m.date = &t
}

// Date returns the date value in the mutation.
func (m *WarningMutation) Date() (r time.Time, exists bool) {
	v := m.date
	if v == nil {
		return
	}
	return *v, true
}

// OldDate returns the old date value of the Warning.
// If the Warning object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *WarningMutation) OldDate(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDate is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDate: %w", err)
	}
	return oldValue.Date, nil
}

// ResetDate reset all changes of the "date" field.
func (m *WarningMutation) ResetDate() {
	m.date = nil
}

// SetUserID sets the user edge to User by id.
func (m *WarningMutation) SetUserID(id discord.UserID) {
	m.user = &id
}

// ClearUser clears the user edge to User.
func (m *WarningMutation) ClearUser() {
	m.cleareduser = true
}

// UserCleared returns if the edge user was cleared.
func (m *WarningMutation) UserCleared() bool {
	return m.cleareduser
}

// UserID returns the user id in the mutation.
func (m *WarningMutation) UserID() (id discord.UserID, exists bool) {
	if m.user != nil {
		return *m.user, true
	}
	return
}

// UserIDs returns the user ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// UserID instead. It exists only for internal usage by the builders.
func (m *WarningMutation) UserIDs() (ids []discord.UserID) {
	if id := m.user; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetUser reset all changes of the "user" edge.
func (m *WarningMutation) ResetUser() {
	m.user = nil
	m.cleareduser = false
}

// SetIssuedByID sets the issuedBy edge to User by id.
func (m *WarningMutation) SetIssuedByID(id discord.UserID) {
	m.issuedBy = &id
}

// ClearIssuedBy clears the issuedBy edge to User.
func (m *WarningMutation) ClearIssuedBy() {
	m.clearedissuedBy = true
}

// IssuedByCleared returns if the edge issuedBy was cleared.
func (m *WarningMutation) IssuedByCleared() bool {
	return m.clearedissuedBy
}

// IssuedByID returns the issuedBy id in the mutation.
func (m *WarningMutation) IssuedByID() (id discord.UserID, exists bool) {
	if m.issuedBy != nil {
		return *m.issuedBy, true
	}
	return
}

// IssuedByIDs returns the issuedBy ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// IssuedByID instead. It exists only for internal usage by the builders.
func (m *WarningMutation) IssuedByIDs() (ids []discord.UserID) {
	if id := m.issuedBy; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetIssuedBy reset all changes of the "issuedBy" edge.
func (m *WarningMutation) ResetIssuedBy() {
	m.issuedBy = nil
	m.clearedissuedBy = false
}

// SetGuildID sets the guild edge to Guild by id.
func (m *WarningMutation) SetGuildID(id discord.GuildID) {
	m.guild = &id
}

// ClearGuild clears the guild edge to Guild.
func (m *WarningMutation) ClearGuild() {
	m.clearedguild = true
}

// GuildCleared returns if the edge guild was cleared.
func (m *WarningMutation) GuildCleared() bool {
	return m.clearedguild
}

// GuildID returns the guild id in the mutation.
func (m *WarningMutation) GuildID() (id discord.GuildID, exists bool) {
	if m.guild != nil {
		return *m.guild, true
	}
	return
}

// GuildIDs returns the guild ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// GuildID instead. It exists only for internal usage by the builders.
func (m *WarningMutation) GuildIDs() (ids []discord.GuildID) {
	if id := m.guild; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetGuild reset all changes of the "guild" edge.
func (m *WarningMutation) ResetGuild() {
	m.guild = nil
	m.clearedguild = false
}

// Op returns the operation name.
func (m *WarningMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Warning).
func (m *WarningMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *WarningMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.reason != nil {
		fields = append(fields, warning.FieldReason)
	}
	if m.date != nil {
		fields = append(fields, warning.FieldDate)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *WarningMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case warning.FieldReason:
		return m.Reason()
	case warning.FieldDate:
		return m.Date()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *WarningMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case warning.FieldReason:
		return m.OldReason(ctx)
	case warning.FieldDate:
		return m.OldDate(ctx)
	}
	return nil, fmt.Errorf("unknown Warning field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *WarningMutation) SetField(name string, value ent.Value) error {
	switch name {
	case warning.FieldReason:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetReason(v)
		return nil
	case warning.FieldDate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDate(v)
		return nil
	}
	return fmt.Errorf("unknown Warning field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *WarningMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *WarningMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *WarningMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Warning numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *WarningMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *WarningMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *WarningMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Warning nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *WarningMutation) ResetField(name string) error {
	switch name {
	case warning.FieldReason:
		m.ResetReason()
		return nil
	case warning.FieldDate:
		m.ResetDate()
		return nil
	}
	return fmt.Errorf("unknown Warning field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *WarningMutation) AddedEdges() []string {
	edges := make([]string, 0, 3)
	if m.user != nil {
		edges = append(edges, warning.EdgeUser)
	}
	if m.issuedBy != nil {
		edges = append(edges, warning.EdgeIssuedBy)
	}
	if m.guild != nil {
		edges = append(edges, warning.EdgeGuild)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *WarningMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case warning.EdgeUser:
		if id := m.user; id != nil {
			return []ent.Value{*id}
		}
	case warning.EdgeIssuedBy:
		if id := m.issuedBy; id != nil {
			return []ent.Value{*id}
		}
	case warning.EdgeGuild:
		if id := m.guild; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *WarningMutation) RemovedEdges() []string {
	edges := make([]string, 0, 3)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *WarningMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *WarningMutation) ClearedEdges() []string {
	edges := make([]string, 0, 3)
	if m.cleareduser {
		edges = append(edges, warning.EdgeUser)
	}
	if m.clearedissuedBy {
		edges = append(edges, warning.EdgeIssuedBy)
	}
	if m.clearedguild {
		edges = append(edges, warning.EdgeGuild)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *WarningMutation) EdgeCleared(name string) bool {
	switch name {
	case warning.EdgeUser:
		return m.cleareduser
	case warning.EdgeIssuedBy:
		return m.clearedissuedBy
	case warning.EdgeGuild:
		return m.clearedguild
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *WarningMutation) ClearEdge(name string) error {
	switch name {
	case warning.EdgeUser:
		m.ClearUser()
		return nil
	case warning.EdgeIssuedBy:
		m.ClearIssuedBy()
		return nil
	case warning.EdgeGuild:
		m.ClearGuild()
		return nil
	}
	return fmt.Errorf("unknown Warning unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *WarningMutation) ResetEdge(name string) error {
	switch name {
	case warning.EdgeUser:
		m.ResetUser()
		return nil
	case warning.EdgeIssuedBy:
		m.ResetIssuedBy()
		return nil
	case warning.EdgeGuild:
		m.ResetGuild()
		return nil
	}
	return fmt.Errorf("unknown Warning edge %s", name)
}
