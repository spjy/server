// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/responserms/server/ent/predicate"
	"github.com/responserms/server/ent/servertype"
)

// ServerTypeDelete is the builder for deleting a ServerType entity.
type ServerTypeDelete struct {
	config
	hooks    []Hook
	mutation *ServerTypeMutation
}

// Where adds a new predicate to the delete builder.
func (std *ServerTypeDelete) Where(ps ...predicate.ServerType) *ServerTypeDelete {
	std.mutation.predicates = append(std.mutation.predicates, ps...)
	return std
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (std *ServerTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(std.hooks) == 0 {
		affected, err = std.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServerTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			std.mutation = mutation
			affected, err = std.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(std.hooks) - 1; i >= 0; i-- {
			mut = std.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, std.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (std *ServerTypeDelete) ExecX(ctx context.Context) int {
	n, err := std.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (std *ServerTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: servertype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: servertype.FieldID,
			},
		},
	}
	if ps := std.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, std.driver, _spec)
}

// ServerTypeDeleteOne is the builder for deleting a single ServerType entity.
type ServerTypeDeleteOne struct {
	std *ServerTypeDelete
}

// Exec executes the deletion query.
func (stdo *ServerTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := stdo.std.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{servertype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (stdo *ServerTypeDeleteOne) ExecX(ctx context.Context) {
	stdo.std.ExecX(ctx)
}
