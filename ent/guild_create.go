// Code generated by entc, DO NOT EDIT.

package ent

import (
	"baritone/bot/commands/guildconfig"
	"baritone/ent/guild"
	"baritone/ent/warning"
	"context"
	"errors"
	"fmt"

	"github.com/diamondburned/arikawa/discord"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// GuildCreate is the builder for creating a Guild entity.
type GuildCreate struct {
	config
	mutation *GuildMutation
	hooks    []Hook
}

// SetConfig sets the config field.
func (gc *GuildCreate) SetConfig(value guildconfig.GuildConfig) *GuildCreate {
	gc.mutation.SetConfig(value)
	return gc
}

// SetID sets the id field.
func (gc *GuildCreate) SetID(di discord.GuildID) *GuildCreate {
	gc.mutation.SetID(di)
	return gc
}

// AddWarningIDs adds the warnings edge to Warning by ids.
func (gc *GuildCreate) AddWarningIDs(ids ...int) *GuildCreate {
	gc.mutation.AddWarningIDs(ids...)
	return gc
}

// AddWarnings adds the warnings edges to Warning.
func (gc *GuildCreate) AddWarnings(w ...*Warning) *GuildCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return gc.AddWarningIDs(ids...)
}

// Mutation returns the GuildMutation object of the builder.
func (gc *GuildCreate) Mutation() *GuildMutation {
	return gc.mutation
}

// Save creates the Guild in the database.
func (gc *GuildCreate) Save(ctx context.Context) (*Guild, error) {
	var (
		err  error
		node *Guild
	)
	if len(gc.hooks) == 0 {
		if err = gc.check(); err != nil {
			return nil, err
		}
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GuildMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gc.check(); err != nil {
				return nil, err
			}
			gc.mutation = mutation
			node, err = gc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			mut = gc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GuildCreate) SaveX(ctx context.Context) *Guild {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (gc *GuildCreate) check() error {
	if _, ok := gc.mutation.Config(); !ok {
		return &ValidationError{Name: "config", err: errors.New("ent: missing required field \"config\"")}
	}
	return nil
}

func (gc *GuildCreate) sqlSave(ctx context.Context) (*Guild, error) {
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = discord.GuildID(id)
	}
	return _node, nil
}

func (gc *GuildCreate) createSpec() (*Guild, *sqlgraph.CreateSpec) {
	var (
		_node = &Guild{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: guild.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: guild.FieldID,
			},
		}
	)
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gc.mutation.Config(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: guild.FieldConfig,
		})
		_node.Config = value
	}
	if nodes := gc.mutation.WarningsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   guild.WarningsTable,
			Columns: []string{guild.WarningsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: warning.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GuildCreateBulk is the builder for creating a bulk of Guild entities.
type GuildCreateBulk struct {
	config
	builders []*GuildCreate
}

// Save creates the Guild entities in the database.
func (gcb *GuildCreateBulk) Save(ctx context.Context) ([]*Guild, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Guild, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GuildMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = discord.GuildID(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (gcb *GuildCreateBulk) SaveX(ctx context.Context) []*Guild {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
