// Code generated by entc, DO NOT EDIT.

package ent

import (
	"baritone/ent/predicate"
	"baritone/ent/warning"
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// WarningDelete is the builder for deleting a Warning entity.
type WarningDelete struct {
	config
	hooks      []Hook
	mutation   *WarningMutation
	predicates []predicate.Warning
}

// Where adds a new predicate to the delete builder.
func (wd *WarningDelete) Where(ps ...predicate.Warning) *WarningDelete {
	wd.predicates = append(wd.predicates, ps...)
	return wd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wd *WarningDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wd.hooks) == 0 {
		affected, err = wd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WarningMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wd.mutation = mutation
			affected, err = wd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wd.hooks) - 1; i >= 0; i-- {
			mut = wd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (wd *WarningDelete) ExecX(ctx context.Context) int {
	n, err := wd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wd *WarningDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: warning.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: warning.FieldID,
			},
		},
	}
	if ps := wd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, wd.driver, _spec)
}

// WarningDeleteOne is the builder for deleting a single Warning entity.
type WarningDeleteOne struct {
	wd *WarningDelete
}

// Exec executes the deletion query.
func (wdo *WarningDeleteOne) Exec(ctx context.Context) error {
	n, err := wdo.wd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{warning.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wdo *WarningDeleteOne) ExecX(ctx context.Context) {
	wdo.wd.ExecX(ctx)
}