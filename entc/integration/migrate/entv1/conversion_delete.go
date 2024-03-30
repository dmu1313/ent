// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv1

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/migrate/entv1/conversion"
	"entgo.io/ent/entc/integration/migrate/entv1/predicate"
	"entgo.io/ent/schema/field"
)

// ConversionDelete is the builder for deleting a Conversion entity.
type ConversionDelete struct {
	config
	hooks    []Hook
	mutation *ConversionMutation
}

// Where appends a list predicates to the ConversionDelete builder.
func (cd *ConversionDelete) Where(ps ...predicate.Conversion) *ConversionDelete {
	cd.mutation.Where(ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *ConversionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cd.sqlExec, cd.mutation, cd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *ConversionDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

// Mutation returns the ConversionMutation object of the builder.
func (cd *ConversionDelete) Mutation() *ConversionMutation {
	return cd.mutation
}

func (cd *ConversionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(conversion.Table, sqlgraph.NewFieldSpec(conversion.FieldID, field.TypeInt))
	if ps := cd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cd.mutation.done = true
	return affected, err
}

// ConversionDeleteOne is the builder for deleting a single Conversion entity.
type ConversionDeleteOne struct {
	cd *ConversionDelete
}

// Where appends a list predicates to the ConversionDelete builder.
func (cdo *ConversionDeleteOne) Where(ps ...predicate.Conversion) *ConversionDeleteOne {
	cdo.cd.mutation.Where(ps...)
	return cdo
}

// Exec executes the deletion query.
func (cdo *ConversionDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{conversion.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *ConversionDeleteOne) ExecX(ctx context.Context) {
	if err := cdo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Mutation returns the ConversionMutation object of the builder.
func (cdo *ConversionDeleteOne) Mutation() *ConversionMutation {
	return cdo.cd.mutation
}
