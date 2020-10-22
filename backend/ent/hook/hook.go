// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/moominzie/user-record/ent"
)

// The BillFunc type is an adapter to allow the use of ordinary
// function as Bill mutator.
type BillFunc func(context.Context, *ent.BillMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BillFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.BillMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BillMutation", m)
	}
	return f(ctx, mv)
}

// The BillingstatusFunc type is an adapter to allow the use of ordinary
// function as Billingstatus mutator.
type BillingstatusFunc func(context.Context, *ent.BillingstatusMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BillingstatusFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.BillingstatusMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BillingstatusMutation", m)
	}
	return f(ctx, mv)
}

// The BranchFunc type is an adapter to allow the use of ordinary
// function as Branch mutator.
type BranchFunc func(context.Context, *ent.BranchMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BranchFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.BranchMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BranchMutation", m)
	}
	return f(ctx, mv)
}

// The BuildingFunc type is an adapter to allow the use of ordinary
// function as Building mutator.
type BuildingFunc func(context.Context, *ent.BuildingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BuildingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.BuildingMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BuildingMutation", m)
	}
	return f(ctx, mv)
}

// The DeviceFunc type is an adapter to allow the use of ordinary
// function as Device mutator.
type DeviceFunc func(context.Context, *ent.DeviceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DeviceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DeviceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DeviceMutation", m)
	}
	return f(ctx, mv)
}

// The EmployeeFunc type is an adapter to allow the use of ordinary
// function as Employee mutator.
type EmployeeFunc func(context.Context, *ent.EmployeeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EmployeeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.EmployeeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EmployeeMutation", m)
	}
	return f(ctx, mv)
}

// The FacultyFunc type is an adapter to allow the use of ordinary
// function as Faculty mutator.
type FacultyFunc func(context.Context, *ent.FacultyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FacultyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FacultyMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FacultyMutation", m)
	}
	return f(ctx, mv)
}

// The RepairInvoiceFunc type is an adapter to allow the use of ordinary
// function as RepairInvoice mutator.
type RepairInvoiceFunc func(context.Context, *ent.RepairInvoiceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RepairInvoiceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RepairInvoiceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RepairInvoiceMutation", m)
	}
	return f(ctx, mv)
}

// The ReturninvoiceFunc type is an adapter to allow the use of ordinary
// function as Returninvoice mutator.
type ReturninvoiceFunc func(context.Context, *ent.ReturninvoiceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ReturninvoiceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ReturninvoiceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ReturninvoiceMutation", m)
	}
	return f(ctx, mv)
}

// The RoomFunc type is an adapter to allow the use of ordinary
// function as Room mutator.
type RoomFunc func(context.Context, *ent.RoomMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RoomFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RoomMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RoomMutation", m)
	}
	return f(ctx, mv)
}

// The StatusRFunc type is an adapter to allow the use of ordinary
// function as StatusR mutator.
type StatusRFunc func(context.Context, *ent.StatusRMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StatusRFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.StatusRMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StatusRMutation", m)
	}
	return f(ctx, mv)
}

// The StatustFunc type is an adapter to allow the use of ordinary
// function as Statust mutator.
type StatustFunc func(context.Context, *ent.StatustMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StatustFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.StatustMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StatustMutation", m)
	}
	return f(ctx, mv)
}

// The SymptomFunc type is an adapter to allow the use of ordinary
// function as Symptom mutator.
type SymptomFunc func(context.Context, *ent.SymptomMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SymptomFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SymptomMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SymptomMutation", m)
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
//	Hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
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

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(_ context.Context, m ent.Mutation) (ent.Value, error) {
			return nil, fmt.Errorf("%s operation is not allowed", m.Op())
		})
	}
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
