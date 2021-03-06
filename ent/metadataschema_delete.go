// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/responserms/server/ent/metadataschema"
	"github.com/responserms/server/ent/predicate"
)

// MetadataSchemaDelete is the builder for deleting a MetadataSchema entity.
type MetadataSchemaDelete struct {
	config
	hooks    []Hook
	mutation *MetadataSchemaMutation
}

// Where adds a new predicate to the delete builder.
func (msd *MetadataSchemaDelete) Where(ps ...predicate.MetadataSchema) *MetadataSchemaDelete {
	msd.mutation.predicates = append(msd.mutation.predicates, ps...)
	return msd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (msd *MetadataSchemaDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(msd.hooks) == 0 {
		affected, err = msd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MetadataSchemaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			msd.mutation = mutation
			affected, err = msd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(msd.hooks) - 1; i >= 0; i-- {
			mut = msd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, msd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (msd *MetadataSchemaDelete) ExecX(ctx context.Context) int {
	n, err := msd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (msd *MetadataSchemaDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: metadataschema.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: metadataschema.FieldID,
			},
		},
	}
	if ps := msd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, msd.driver, _spec)
}

// MetadataSchemaDeleteOne is the builder for deleting a single MetadataSchema entity.
type MetadataSchemaDeleteOne struct {
	msd *MetadataSchemaDelete
}

// Exec executes the deletion query.
func (msdo *MetadataSchemaDeleteOne) Exec(ctx context.Context) error {
	n, err := msdo.msd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{metadataschema.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (msdo *MetadataSchemaDeleteOne) ExecX(ctx context.Context) {
	msdo.msd.ExecX(ctx)
}
