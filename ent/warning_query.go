// Code generated by entc, DO NOT EDIT.

package ent

import (
	"baritone/ent/guild"
	"baritone/ent/predicate"
	"baritone/ent/user"
	"baritone/ent/warning"
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/diamondburned/arikawa/discord"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// WarningQuery is the builder for querying Warning entities.
type WarningQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Warning
	// eager-loading edges.
	withUser     *UserQuery
	withIssuedBy *UserQuery
	withGuild    *GuildQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (wq *WarningQuery) Where(ps ...predicate.Warning) *WarningQuery {
	wq.predicates = append(wq.predicates, ps...)
	return wq
}

// Limit adds a limit step to the query.
func (wq *WarningQuery) Limit(limit int) *WarningQuery {
	wq.limit = &limit
	return wq
}

// Offset adds an offset step to the query.
func (wq *WarningQuery) Offset(offset int) *WarningQuery {
	wq.offset = &offset
	return wq
}

// Order adds an order step to the query.
func (wq *WarningQuery) Order(o ...OrderFunc) *WarningQuery {
	wq.order = append(wq.order, o...)
	return wq
}

// QueryUser chains the current query on the user edge.
func (wq *WarningQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: wq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(warning.Table, warning.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, warning.UserTable, warning.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryIssuedBy chains the current query on the issuedBy edge.
func (wq *WarningQuery) QueryIssuedBy() *UserQuery {
	query := &UserQuery{config: wq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(warning.Table, warning.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, warning.IssuedByTable, warning.IssuedByColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGuild chains the current query on the guild edge.
func (wq *WarningQuery) QueryGuild() *GuildQuery {
	query := &GuildQuery{config: wq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(warning.Table, warning.FieldID, selector),
			sqlgraph.To(guild.Table, guild.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, warning.GuildTable, warning.GuildColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Warning entity in the query. Returns *NotFoundError when no warning was found.
func (wq *WarningQuery) First(ctx context.Context) (*Warning, error) {
	nodes, err := wq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{warning.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wq *WarningQuery) FirstX(ctx context.Context) *Warning {
	node, err := wq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Warning id in the query. Returns *NotFoundError when no id was found.
func (wq *WarningQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{warning.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (wq *WarningQuery) FirstXID(ctx context.Context) int {
	id, err := wq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Warning entity in the query, returns an error if not exactly one entity was returned.
func (wq *WarningQuery) Only(ctx context.Context) (*Warning, error) {
	nodes, err := wq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{warning.Label}
	default:
		return nil, &NotSingularError{warning.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wq *WarningQuery) OnlyX(ctx context.Context) *Warning {
	node, err := wq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only Warning id in the query, returns an error if not exactly one id was returned.
func (wq *WarningQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{warning.Label}
	default:
		err = &NotSingularError{warning.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wq *WarningQuery) OnlyIDX(ctx context.Context) int {
	id, err := wq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Warnings.
func (wq *WarningQuery) All(ctx context.Context) ([]*Warning, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wq *WarningQuery) AllX(ctx context.Context) []*Warning {
	nodes, err := wq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Warning ids.
func (wq *WarningQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := wq.Select(warning.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wq *WarningQuery) IDsX(ctx context.Context) []int {
	ids, err := wq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wq *WarningQuery) Count(ctx context.Context) (int, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wq *WarningQuery) CountX(ctx context.Context) int {
	count, err := wq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wq *WarningQuery) Exist(ctx context.Context) (bool, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wq *WarningQuery) ExistX(ctx context.Context) bool {
	exist, err := wq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wq *WarningQuery) Clone() *WarningQuery {
	return &WarningQuery{
		config:     wq.config,
		limit:      wq.limit,
		offset:     wq.offset,
		order:      append([]OrderFunc{}, wq.order...),
		unique:     append([]string{}, wq.unique...),
		predicates: append([]predicate.Warning{}, wq.predicates...),
		// clone intermediate query.
		sql:  wq.sql.Clone(),
		path: wq.path,
	}
}

//  WithUser tells the query-builder to eager-loads the nodes that are connected to
// the "user" edge. The optional arguments used to configure the query builder of the edge.
func (wq *WarningQuery) WithUser(opts ...func(*UserQuery)) *WarningQuery {
	query := &UserQuery{config: wq.config}
	for _, opt := range opts {
		opt(query)
	}
	wq.withUser = query
	return wq
}

//  WithIssuedBy tells the query-builder to eager-loads the nodes that are connected to
// the "issuedBy" edge. The optional arguments used to configure the query builder of the edge.
func (wq *WarningQuery) WithIssuedBy(opts ...func(*UserQuery)) *WarningQuery {
	query := &UserQuery{config: wq.config}
	for _, opt := range opts {
		opt(query)
	}
	wq.withIssuedBy = query
	return wq
}

//  WithGuild tells the query-builder to eager-loads the nodes that are connected to
// the "guild" edge. The optional arguments used to configure the query builder of the edge.
func (wq *WarningQuery) WithGuild(opts ...func(*GuildQuery)) *WarningQuery {
	query := &GuildQuery{config: wq.config}
	for _, opt := range opts {
		opt(query)
	}
	wq.withGuild = query
	return wq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Reason string `json:"reason,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Warning.Query().
//		GroupBy(warning.FieldReason).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (wq *WarningQuery) GroupBy(field string, fields ...string) *WarningGroupBy {
	group := &WarningGroupBy{config: wq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Reason string `json:"reason,omitempty"`
//	}
//
//	client.Warning.Query().
//		Select(warning.FieldReason).
//		Scan(ctx, &v)
//
func (wq *WarningQuery) Select(field string, fields ...string) *WarningSelect {
	selector := &WarningSelect{config: wq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wq.sqlQuery(), nil
	}
	return selector
}

func (wq *WarningQuery) prepareQuery(ctx context.Context) error {
	if wq.path != nil {
		prev, err := wq.path(ctx)
		if err != nil {
			return err
		}
		wq.sql = prev
	}
	return nil
}

func (wq *WarningQuery) sqlAll(ctx context.Context) ([]*Warning, error) {
	var (
		nodes       = []*Warning{}
		withFKs     = wq.withFKs
		_spec       = wq.querySpec()
		loadedTypes = [3]bool{
			wq.withUser != nil,
			wq.withIssuedBy != nil,
			wq.withGuild != nil,
		}
	)
	if wq.withUser != nil || wq.withIssuedBy != nil || wq.withGuild != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, warning.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Warning{config: wq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
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
	if err := sqlgraph.QueryNodes(ctx, wq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := wq.withUser; query != nil {
		ids := make([]discord.UserID, 0, len(nodes))
		nodeids := make(map[discord.UserID][]*Warning)
		for i := range nodes {
			if fk := nodes[i].warning_user; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "warning_user" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	if query := wq.withIssuedBy; query != nil {
		ids := make([]discord.UserID, 0, len(nodes))
		nodeids := make(map[discord.UserID][]*Warning)
		for i := range nodes {
			if fk := nodes[i].warning_issued_by; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "warning_issued_by" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.IssuedBy = n
			}
		}
	}

	if query := wq.withGuild; query != nil {
		ids := make([]discord.GuildID, 0, len(nodes))
		nodeids := make(map[discord.GuildID][]*Warning)
		for i := range nodes {
			if fk := nodes[i].warning_guild; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(guild.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "warning_guild" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Guild = n
			}
		}
	}

	return nodes, nil
}

func (wq *WarningQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wq.querySpec()
	return sqlgraph.CountNodes(ctx, wq.driver, _spec)
}

func (wq *WarningQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (wq *WarningQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   warning.Table,
			Columns: warning.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: warning.FieldID,
			},
		},
		From:   wq.sql,
		Unique: true,
	}
	if ps := wq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, warning.ValidColumn)
			}
		}
	}
	return _spec
}

func (wq *WarningQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(wq.driver.Dialect())
	t1 := builder.Table(warning.Table)
	selector := builder.Select(t1.Columns(warning.Columns...)...).From(t1)
	if wq.sql != nil {
		selector = wq.sql
		selector.Select(selector.Columns(warning.Columns...)...)
	}
	for _, p := range wq.predicates {
		p(selector)
	}
	for _, p := range wq.order {
		p(selector, warning.ValidColumn)
	}
	if offset := wq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WarningGroupBy is the builder for group-by Warning entities.
type WarningGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wgb *WarningGroupBy) Aggregate(fns ...AggregateFunc) *WarningGroupBy {
	wgb.fns = append(wgb.fns, fns...)
	return wgb
}

// Scan applies the group-by query and scan the result into the given value.
func (wgb *WarningGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wgb.path(ctx)
	if err != nil {
		return err
	}
	wgb.sql = query
	return wgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wgb *WarningGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (wgb *WarningGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: WarningGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wgb *WarningGroupBy) StringsX(ctx context.Context) []string {
	v, err := wgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (wgb *WarningGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{warning.Label}
	default:
		err = fmt.Errorf("ent: WarningGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wgb *WarningGroupBy) StringX(ctx context.Context) string {
	v, err := wgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (wgb *WarningGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: WarningGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wgb *WarningGroupBy) IntsX(ctx context.Context) []int {
	v, err := wgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (wgb *WarningGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{warning.Label}
	default:
		err = fmt.Errorf("ent: WarningGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wgb *WarningGroupBy) IntX(ctx context.Context) int {
	v, err := wgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (wgb *WarningGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: WarningGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wgb *WarningGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (wgb *WarningGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{warning.Label}
	default:
		err = fmt.Errorf("ent: WarningGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wgb *WarningGroupBy) Float64X(ctx context.Context) float64 {
	v, err := wgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (wgb *WarningGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: WarningGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wgb *WarningGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (wgb *WarningGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{warning.Label}
	default:
		err = fmt.Errorf("ent: WarningGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wgb *WarningGroupBy) BoolX(ctx context.Context) bool {
	v, err := wgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wgb *WarningGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wgb.fields {
		if !warning.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wgb *WarningGroupBy) sqlQuery() *sql.Selector {
	selector := wgb.sql
	columns := make([]string, 0, len(wgb.fields)+len(wgb.fns))
	columns = append(columns, wgb.fields...)
	for _, fn := range wgb.fns {
		columns = append(columns, fn(selector, warning.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(wgb.fields...)
}

// WarningSelect is the builder for select fields of Warning entities.
type WarningSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ws *WarningSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ws.path(ctx)
	if err != nil {
		return err
	}
	ws.sql = query
	return ws.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ws *WarningSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ws.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ws *WarningSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: WarningSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ws *WarningSelect) StringsX(ctx context.Context) []string {
	v, err := ws.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ws *WarningSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ws.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{warning.Label}
	default:
		err = fmt.Errorf("ent: WarningSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ws *WarningSelect) StringX(ctx context.Context) string {
	v, err := ws.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ws *WarningSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: WarningSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ws *WarningSelect) IntsX(ctx context.Context) []int {
	v, err := ws.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ws *WarningSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ws.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{warning.Label}
	default:
		err = fmt.Errorf("ent: WarningSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ws *WarningSelect) IntX(ctx context.Context) int {
	v, err := ws.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ws *WarningSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: WarningSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ws *WarningSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ws.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ws *WarningSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ws.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{warning.Label}
	default:
		err = fmt.Errorf("ent: WarningSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ws *WarningSelect) Float64X(ctx context.Context) float64 {
	v, err := ws.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ws *WarningSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: WarningSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ws *WarningSelect) BoolsX(ctx context.Context) []bool {
	v, err := ws.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ws *WarningSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ws.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{warning.Label}
	default:
		err = fmt.Errorf("ent: WarningSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ws *WarningSelect) BoolX(ctx context.Context) bool {
	v, err := ws.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ws *WarningSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ws.fields {
		if !warning.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := ws.sqlQuery().Query()
	if err := ws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ws *WarningSelect) sqlQuery() sql.Querier {
	selector := ws.sql
	selector.Select(selector.Columns(ws.fields...)...)
	return selector
}
