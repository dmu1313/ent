// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgefield/ent/info"
	"entgo.io/ent/entc/integration/edgefield/ent/predicate"
	"entgo.io/ent/schema/field"
)

// InfoDelete is the builder for deleting a Info entity.
type InfoDelete struct {
	config
	hooks    []Hook
	mutation *InfoMutation
}

// Where appends a list predicates to the InfoDelete builder.
func (id *InfoDelete) Where(ps ...predicate.Info) *InfoDelete {
	id.mutation.Where(ps...)
	return id
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (id *InfoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, id.sqlExec, id.mutation, id.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (id *InfoDelete) ExecX(ctx context.Context) int {
	n, err := id.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

// Mutation returns the InfoMutation object of the builder.
func (id *InfoDelete) Mutation() *InfoMutation {
	return id.mutation
}

func (id *InfoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(info.Table, sqlgraph.NewFieldSpec(info.FieldID, field.TypeInt))
	if ps := id.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, id.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	id.mutation.done = true
	return affected, err
}

// InfoDeleteOne is the builder for deleting a single Info entity.
type InfoDeleteOne struct {
	id *InfoDelete
}

// Where appends a list predicates to the InfoDelete builder.
func (ido *InfoDeleteOne) Where(ps ...predicate.Info) *InfoDeleteOne {
	ido.id.mutation.Where(ps...)
	return ido
}

// Exec executes the deletion query.
func (ido *InfoDeleteOne) Exec(ctx context.Context) error {
	n, err := ido.id.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{info.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ido *InfoDeleteOne) ExecX(ctx context.Context) {
	if err := ido.Exec(ctx); err != nil {
		panic(err)
	}
}

// Mutation returns the InfoMutation object of the builder.
func (ido *InfoDeleteOne) Mutation() *InfoMutation {
	return ido.id.mutation
}
