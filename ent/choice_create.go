// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eesoymilk/health-statistic-api/ent/answer"
	"github.com/eesoymilk/health-statistic-api/ent/choice"
	"github.com/eesoymilk/health-statistic-api/ent/question"
	"github.com/google/uuid"
)

// ChoiceCreate is the builder for creating a Choice entity.
type ChoiceCreate struct {
	config
	mutation *ChoiceMutation
	hooks    []Hook
}

// SetBody sets the "body" field.
func (cc *ChoiceCreate) SetBody(s string) *ChoiceCreate {
	cc.mutation.SetBody(s)
	return cc
}

// SetOrder sets the "order" field.
func (cc *ChoiceCreate) SetOrder(i int) *ChoiceCreate {
	cc.mutation.SetOrder(i)
	return cc
}

// SetID sets the "id" field.
func (cc *ChoiceCreate) SetID(u uuid.UUID) *ChoiceCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *ChoiceCreate) SetNillableID(u *uuid.UUID) *ChoiceCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// SetQuesionID sets the "quesion" edge to the Question entity by ID.
func (cc *ChoiceCreate) SetQuesionID(id uuid.UUID) *ChoiceCreate {
	cc.mutation.SetQuesionID(id)
	return cc
}

// SetNillableQuesionID sets the "quesion" edge to the Question entity by ID if the given value is not nil.
func (cc *ChoiceCreate) SetNillableQuesionID(id *uuid.UUID) *ChoiceCreate {
	if id != nil {
		cc = cc.SetQuesionID(*id)
	}
	return cc
}

// SetQuesion sets the "quesion" edge to the Question entity.
func (cc *ChoiceCreate) SetQuesion(q *Question) *ChoiceCreate {
	return cc.SetQuesionID(q.ID)
}

// AddAnswerIDs adds the "answer" edge to the Answer entity by IDs.
func (cc *ChoiceCreate) AddAnswerIDs(ids ...uuid.UUID) *ChoiceCreate {
	cc.mutation.AddAnswerIDs(ids...)
	return cc
}

// AddAnswer adds the "answer" edges to the Answer entity.
func (cc *ChoiceCreate) AddAnswer(a ...*Answer) *ChoiceCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cc.AddAnswerIDs(ids...)
}

// Mutation returns the ChoiceMutation object of the builder.
func (cc *ChoiceCreate) Mutation() *ChoiceMutation {
	return cc.mutation
}

// Save creates the Choice in the database.
func (cc *ChoiceCreate) Save(ctx context.Context) (*Choice, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ChoiceCreate) SaveX(ctx context.Context) *Choice {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ChoiceCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ChoiceCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ChoiceCreate) defaults() {
	if _, ok := cc.mutation.ID(); !ok {
		v := choice.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ChoiceCreate) check() error {
	if _, ok := cc.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New(`ent: missing required field "Choice.body"`)}
	}
	if v, ok := cc.mutation.Body(); ok {
		if err := choice.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`ent: validator failed for field "Choice.body": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Order(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required field "Choice.order"`)}
	}
	if v, ok := cc.mutation.Order(); ok {
		if err := choice.OrderValidator(v); err != nil {
			return &ValidationError{Name: "order", err: fmt.Errorf(`ent: validator failed for field "Choice.order": %w`, err)}
		}
	}
	return nil
}

func (cc *ChoiceCreate) sqlSave(ctx context.Context) (*Choice, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ChoiceCreate) createSpec() (*Choice, *sqlgraph.CreateSpec) {
	var (
		_node = &Choice{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(choice.Table, sqlgraph.NewFieldSpec(choice.FieldID, field.TypeUUID))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.Body(); ok {
		_spec.SetField(choice.FieldBody, field.TypeString, value)
		_node.Body = value
	}
	if value, ok := cc.mutation.Order(); ok {
		_spec.SetField(choice.FieldOrder, field.TypeInt, value)
		_node.Order = value
	}
	if nodes := cc.mutation.QuesionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   choice.QuesionTable,
			Columns: []string{choice.QuesionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.question_choices = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.AnswerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   choice.AnswerTable,
			Columns: choice.AnswerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ChoiceCreateBulk is the builder for creating many Choice entities in bulk.
type ChoiceCreateBulk struct {
	config
	builders []*ChoiceCreate
}

// Save creates the Choice entities in the database.
func (ccb *ChoiceCreateBulk) Save(ctx context.Context) ([]*Choice, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Choice, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChoiceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ChoiceCreateBulk) SaveX(ctx context.Context) []*Choice {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ChoiceCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ChoiceCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
