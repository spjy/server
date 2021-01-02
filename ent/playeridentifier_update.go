// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/responserms/server/ent/metadata"
	"github.com/responserms/server/ent/player"
	"github.com/responserms/server/ent/playeridentifier"
	"github.com/responserms/server/ent/predicate"
)

// PlayerIdentifierUpdate is the builder for updating PlayerIdentifier entities.
type PlayerIdentifierUpdate struct {
	config
	hooks    []Hook
	mutation *PlayerIdentifierMutation
}

// Where adds a new predicate for the builder.
func (piu *PlayerIdentifierUpdate) Where(ps ...predicate.PlayerIdentifier) *PlayerIdentifierUpdate {
	piu.mutation.predicates = append(piu.mutation.predicates, ps...)
	return piu
}

// SetUpdatedAt sets the updated_at field.
func (piu *PlayerIdentifierUpdate) SetUpdatedAt(t time.Time) *PlayerIdentifierUpdate {
	piu.mutation.SetUpdatedAt(t)
	return piu
}

// SetType sets the type field.
func (piu *PlayerIdentifierUpdate) SetType(s string) *PlayerIdentifierUpdate {
	piu.mutation.SetType(s)
	return piu
}

// SetIdentifier sets the identifier field.
func (piu *PlayerIdentifierUpdate) SetIdentifier(s string) *PlayerIdentifierUpdate {
	piu.mutation.SetIdentifier(s)
	return piu
}

// SetMetadataID sets the metadata edge to Metadata by id.
func (piu *PlayerIdentifierUpdate) SetMetadataID(id int) *PlayerIdentifierUpdate {
	piu.mutation.SetMetadataID(id)
	return piu
}

// SetNillableMetadataID sets the metadata edge to Metadata by id if the given value is not nil.
func (piu *PlayerIdentifierUpdate) SetNillableMetadataID(id *int) *PlayerIdentifierUpdate {
	if id != nil {
		piu = piu.SetMetadataID(*id)
	}
	return piu
}

// SetMetadata sets the metadata edge to Metadata.
func (piu *PlayerIdentifierUpdate) SetMetadata(m *Metadata) *PlayerIdentifierUpdate {
	return piu.SetMetadataID(m.ID)
}

// SetPlayerID sets the player edge to Player by id.
func (piu *PlayerIdentifierUpdate) SetPlayerID(id int) *PlayerIdentifierUpdate {
	piu.mutation.SetPlayerID(id)
	return piu
}

// SetNillablePlayerID sets the player edge to Player by id if the given value is not nil.
func (piu *PlayerIdentifierUpdate) SetNillablePlayerID(id *int) *PlayerIdentifierUpdate {
	if id != nil {
		piu = piu.SetPlayerID(*id)
	}
	return piu
}

// SetPlayer sets the player edge to Player.
func (piu *PlayerIdentifierUpdate) SetPlayer(p *Player) *PlayerIdentifierUpdate {
	return piu.SetPlayerID(p.ID)
}

// Mutation returns the PlayerIdentifierMutation object of the builder.
func (piu *PlayerIdentifierUpdate) Mutation() *PlayerIdentifierMutation {
	return piu.mutation
}

// ClearMetadata clears the "metadata" edge to type Metadata.
func (piu *PlayerIdentifierUpdate) ClearMetadata() *PlayerIdentifierUpdate {
	piu.mutation.ClearMetadata()
	return piu
}

// ClearPlayer clears the "player" edge to type Player.
func (piu *PlayerIdentifierUpdate) ClearPlayer() *PlayerIdentifierUpdate {
	piu.mutation.ClearPlayer()
	return piu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (piu *PlayerIdentifierUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	piu.defaults()
	if len(piu.hooks) == 0 {
		affected, err = piu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlayerIdentifierMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			piu.mutation = mutation
			affected, err = piu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(piu.hooks) - 1; i >= 0; i-- {
			mut = piu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, piu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (piu *PlayerIdentifierUpdate) SaveX(ctx context.Context) int {
	affected, err := piu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (piu *PlayerIdentifierUpdate) Exec(ctx context.Context) error {
	_, err := piu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piu *PlayerIdentifierUpdate) ExecX(ctx context.Context) {
	if err := piu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (piu *PlayerIdentifierUpdate) defaults() {
	if _, ok := piu.mutation.UpdatedAt(); !ok {
		v := playeridentifier.UpdateDefaultUpdatedAt()
		piu.mutation.SetUpdatedAt(v)
	}
}

func (piu *PlayerIdentifierUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   playeridentifier.Table,
			Columns: playeridentifier.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: playeridentifier.FieldID,
			},
		},
	}
	if ps := piu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: playeridentifier.FieldUpdatedAt,
		})
	}
	if value, ok := piu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playeridentifier.FieldType,
		})
	}
	if value, ok := piu.mutation.Identifier(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playeridentifier.FieldIdentifier,
		})
	}
	if piu.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   playeridentifier.MetadataTable,
			Columns: []string{playeridentifier.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metadata.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piu.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   playeridentifier.MetadataTable,
			Columns: []string{playeridentifier.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metadata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if piu.mutation.PlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   playeridentifier.PlayerTable,
			Columns: []string{playeridentifier.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: player.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piu.mutation.PlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   playeridentifier.PlayerTable,
			Columns: []string{playeridentifier.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: player.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, piu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{playeridentifier.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PlayerIdentifierUpdateOne is the builder for updating a single PlayerIdentifier entity.
type PlayerIdentifierUpdateOne struct {
	config
	hooks    []Hook
	mutation *PlayerIdentifierMutation
}

// SetUpdatedAt sets the updated_at field.
func (piuo *PlayerIdentifierUpdateOne) SetUpdatedAt(t time.Time) *PlayerIdentifierUpdateOne {
	piuo.mutation.SetUpdatedAt(t)
	return piuo
}

// SetType sets the type field.
func (piuo *PlayerIdentifierUpdateOne) SetType(s string) *PlayerIdentifierUpdateOne {
	piuo.mutation.SetType(s)
	return piuo
}

// SetIdentifier sets the identifier field.
func (piuo *PlayerIdentifierUpdateOne) SetIdentifier(s string) *PlayerIdentifierUpdateOne {
	piuo.mutation.SetIdentifier(s)
	return piuo
}

// SetMetadataID sets the metadata edge to Metadata by id.
func (piuo *PlayerIdentifierUpdateOne) SetMetadataID(id int) *PlayerIdentifierUpdateOne {
	piuo.mutation.SetMetadataID(id)
	return piuo
}

// SetNillableMetadataID sets the metadata edge to Metadata by id if the given value is not nil.
func (piuo *PlayerIdentifierUpdateOne) SetNillableMetadataID(id *int) *PlayerIdentifierUpdateOne {
	if id != nil {
		piuo = piuo.SetMetadataID(*id)
	}
	return piuo
}

// SetMetadata sets the metadata edge to Metadata.
func (piuo *PlayerIdentifierUpdateOne) SetMetadata(m *Metadata) *PlayerIdentifierUpdateOne {
	return piuo.SetMetadataID(m.ID)
}

// SetPlayerID sets the player edge to Player by id.
func (piuo *PlayerIdentifierUpdateOne) SetPlayerID(id int) *PlayerIdentifierUpdateOne {
	piuo.mutation.SetPlayerID(id)
	return piuo
}

// SetNillablePlayerID sets the player edge to Player by id if the given value is not nil.
func (piuo *PlayerIdentifierUpdateOne) SetNillablePlayerID(id *int) *PlayerIdentifierUpdateOne {
	if id != nil {
		piuo = piuo.SetPlayerID(*id)
	}
	return piuo
}

// SetPlayer sets the player edge to Player.
func (piuo *PlayerIdentifierUpdateOne) SetPlayer(p *Player) *PlayerIdentifierUpdateOne {
	return piuo.SetPlayerID(p.ID)
}

// Mutation returns the PlayerIdentifierMutation object of the builder.
func (piuo *PlayerIdentifierUpdateOne) Mutation() *PlayerIdentifierMutation {
	return piuo.mutation
}

// ClearMetadata clears the "metadata" edge to type Metadata.
func (piuo *PlayerIdentifierUpdateOne) ClearMetadata() *PlayerIdentifierUpdateOne {
	piuo.mutation.ClearMetadata()
	return piuo
}

// ClearPlayer clears the "player" edge to type Player.
func (piuo *PlayerIdentifierUpdateOne) ClearPlayer() *PlayerIdentifierUpdateOne {
	piuo.mutation.ClearPlayer()
	return piuo
}

// Save executes the query and returns the updated entity.
func (piuo *PlayerIdentifierUpdateOne) Save(ctx context.Context) (*PlayerIdentifier, error) {
	var (
		err  error
		node *PlayerIdentifier
	)
	piuo.defaults()
	if len(piuo.hooks) == 0 {
		node, err = piuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlayerIdentifierMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			piuo.mutation = mutation
			node, err = piuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(piuo.hooks) - 1; i >= 0; i-- {
			mut = piuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, piuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (piuo *PlayerIdentifierUpdateOne) SaveX(ctx context.Context) *PlayerIdentifier {
	node, err := piuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (piuo *PlayerIdentifierUpdateOne) Exec(ctx context.Context) error {
	_, err := piuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piuo *PlayerIdentifierUpdateOne) ExecX(ctx context.Context) {
	if err := piuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (piuo *PlayerIdentifierUpdateOne) defaults() {
	if _, ok := piuo.mutation.UpdatedAt(); !ok {
		v := playeridentifier.UpdateDefaultUpdatedAt()
		piuo.mutation.SetUpdatedAt(v)
	}
}

func (piuo *PlayerIdentifierUpdateOne) sqlSave(ctx context.Context) (_node *PlayerIdentifier, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   playeridentifier.Table,
			Columns: playeridentifier.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: playeridentifier.FieldID,
			},
		},
	}
	id, ok := piuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing PlayerIdentifier.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := piuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: playeridentifier.FieldUpdatedAt,
		})
	}
	if value, ok := piuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playeridentifier.FieldType,
		})
	}
	if value, ok := piuo.mutation.Identifier(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playeridentifier.FieldIdentifier,
		})
	}
	if piuo.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   playeridentifier.MetadataTable,
			Columns: []string{playeridentifier.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metadata.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piuo.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   playeridentifier.MetadataTable,
			Columns: []string{playeridentifier.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metadata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if piuo.mutation.PlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   playeridentifier.PlayerTable,
			Columns: []string{playeridentifier.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: player.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piuo.mutation.PlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   playeridentifier.PlayerTable,
			Columns: []string{playeridentifier.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: player.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PlayerIdentifier{config: piuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, piuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{playeridentifier.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
