// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/responserms/server/ent"
)

// The ActivationFunc type is an adapter to allow the use of ordinary
// function as Activation mutator.
type ActivationFunc func(context.Context, *ent.ActivationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ActivationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ActivationMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ActivationMutation", m)
	}
	return f(ctx, mv)
}

// The AuthSessionFunc type is an adapter to allow the use of ordinary
// function as AuthSession mutator.
type AuthSessionFunc func(context.Context, *ent.AuthSessionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AuthSessionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AuthSessionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AuthSessionMutation", m)
	}
	return f(ctx, mv)
}

// The MapLayerFunc type is an adapter to allow the use of ordinary
// function as MapLayer mutator.
type MapLayerFunc func(context.Context, *ent.MapLayerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MapLayerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.MapLayerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MapLayerMutation", m)
	}
	return f(ctx, mv)
}

// The MapTypeFunc type is an adapter to allow the use of ordinary
// function as MapType mutator.
type MapTypeFunc func(context.Context, *ent.MapTypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MapTypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.MapTypeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MapTypeMutation", m)
	}
	return f(ctx, mv)
}

// The MetadataFunc type is an adapter to allow the use of ordinary
// function as Metadata mutator.
type MetadataFunc func(context.Context, *ent.MetadataMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MetadataFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.MetadataMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MetadataMutation", m)
	}
	return f(ctx, mv)
}

// The MetadataSchemaFunc type is an adapter to allow the use of ordinary
// function as MetadataSchema mutator.
type MetadataSchemaFunc func(context.Context, *ent.MetadataSchemaMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MetadataSchemaFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.MetadataSchemaMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MetadataSchemaMutation", m)
	}
	return f(ctx, mv)
}

// The PlayerFunc type is an adapter to allow the use of ordinary
// function as Player mutator.
type PlayerFunc func(context.Context, *ent.PlayerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PlayerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PlayerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PlayerMutation", m)
	}
	return f(ctx, mv)
}

// The PlayerIdentifierFunc type is an adapter to allow the use of ordinary
// function as PlayerIdentifier mutator.
type PlayerIdentifierFunc func(context.Context, *ent.PlayerIdentifierMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PlayerIdentifierFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PlayerIdentifierMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PlayerIdentifierMutation", m)
	}
	return f(ctx, mv)
}

// The ServerFunc type is an adapter to allow the use of ordinary
// function as Server mutator.
type ServerFunc func(context.Context, *ent.ServerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ServerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ServerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ServerMutation", m)
	}
	return f(ctx, mv)
}

// The ServerTypeFunc type is an adapter to allow the use of ordinary
// function as ServerType mutator.
type ServerTypeFunc func(context.Context, *ent.ServerTypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ServerTypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ServerTypeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ServerTypeMutation", m)
	}
	return f(ctx, mv)
}

// The SessionTokenFunc type is an adapter to allow the use of ordinary
// function as SessionToken mutator.
type SessionTokenFunc func(context.Context, *ent.SessionTokenMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SessionTokenFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SessionTokenMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SessionTokenMutation", m)
	}
	return f(ctx, mv)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
