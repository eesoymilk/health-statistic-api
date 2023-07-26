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
	"github.com/eesoymilk/health-statistic-api/ent/questionnaire"
	"github.com/google/uuid"
)

// QuestionCreate is the builder for creating a Question entity.
type QuestionCreate struct {
	config
	mutation *QuestionMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (qc *QuestionCreate) SetType(q question.Type) *QuestionCreate {
	qc.mutation.SetType(q)
	return qc
}

// SetBody sets the "body" field.
func (qc *QuestionCreate) SetBody(s string) *QuestionCreate {
	qc.mutation.SetBody(s)
	return qc
}

// SetOrder sets the "order" field.
func (qc *QuestionCreate) SetOrder(i int) *QuestionCreate {
	qc.mutation.SetOrder(i)
	return qc
}

// SetID sets the "id" field.
func (qc *QuestionCreate) SetID(u uuid.UUID) *QuestionCreate {
	qc.mutation.SetID(u)
	return qc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (qc *QuestionCreate) SetNillableID(u *uuid.UUID) *QuestionCreate {
	if u != nil {
		qc.SetID(*u)
	}
	return qc
}

// SetQuestionnaireID sets the "questionnaire" edge to the Questionnaire entity by ID.
func (qc *QuestionCreate) SetQuestionnaireID(id uuid.UUID) *QuestionCreate {
	qc.mutation.SetQuestionnaireID(id)
	return qc
}

// SetQuestionnaire sets the "questionnaire" edge to the Questionnaire entity.
func (qc *QuestionCreate) SetQuestionnaire(q *Questionnaire) *QuestionCreate {
	return qc.SetQuestionnaireID(q.ID)
}

// AddChoiceIDs adds the "choices" edge to the Choice entity by IDs.
func (qc *QuestionCreate) AddChoiceIDs(ids ...uuid.UUID) *QuestionCreate {
	qc.mutation.AddChoiceIDs(ids...)
	return qc
}

// AddChoices adds the "choices" edges to the Choice entity.
func (qc *QuestionCreate) AddChoices(c ...*Choice) *QuestionCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return qc.AddChoiceIDs(ids...)
}

// AddAnswerIDs adds the "answers" edge to the Answer entity by IDs.
func (qc *QuestionCreate) AddAnswerIDs(ids ...uuid.UUID) *QuestionCreate {
	qc.mutation.AddAnswerIDs(ids...)
	return qc
}

// AddAnswers adds the "answers" edges to the Answer entity.
func (qc *QuestionCreate) AddAnswers(a ...*Answer) *QuestionCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return qc.AddAnswerIDs(ids...)
}

// Mutation returns the QuestionMutation object of the builder.
func (qc *QuestionCreate) Mutation() *QuestionMutation {
	return qc.mutation
}

// Save creates the Question in the database.
func (qc *QuestionCreate) Save(ctx context.Context) (*Question, error) {
	qc.defaults()
	return withHooks(ctx, qc.sqlSave, qc.mutation, qc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (qc *QuestionCreate) SaveX(ctx context.Context) *Question {
	v, err := qc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qc *QuestionCreate) Exec(ctx context.Context) error {
	_, err := qc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qc *QuestionCreate) ExecX(ctx context.Context) {
	if err := qc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (qc *QuestionCreate) defaults() {
	if _, ok := qc.mutation.ID(); !ok {
		v := question.DefaultID()
		qc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (qc *QuestionCreate) check() error {
	if _, ok := qc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Question.type"`)}
	}
	if v, ok := qc.mutation.GetType(); ok {
		if err := question.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Question.type": %w`, err)}
		}
	}
	if _, ok := qc.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New(`ent: missing required field "Question.body"`)}
	}
	if v, ok := qc.mutation.Body(); ok {
		if err := question.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`ent: validator failed for field "Question.body": %w`, err)}
		}
	}
	if _, ok := qc.mutation.Order(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required field "Question.order"`)}
	}
	if v, ok := qc.mutation.Order(); ok {
		if err := question.OrderValidator(v); err != nil {
			return &ValidationError{Name: "order", err: fmt.Errorf(`ent: validator failed for field "Question.order": %w`, err)}
		}
	}
	if _, ok := qc.mutation.QuestionnaireID(); !ok {
		return &ValidationError{Name: "questionnaire", err: errors.New(`ent: missing required edge "Question.questionnaire"`)}
	}
	return nil
}

func (qc *QuestionCreate) sqlSave(ctx context.Context) (*Question, error) {
	if err := qc.check(); err != nil {
		return nil, err
	}
	_node, _spec := qc.createSpec()
	if err := sqlgraph.CreateNode(ctx, qc.driver, _spec); err != nil {
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
	qc.mutation.id = &_node.ID
	qc.mutation.done = true
	return _node, nil
}

func (qc *QuestionCreate) createSpec() (*Question, *sqlgraph.CreateSpec) {
	var (
		_node = &Question{config: qc.config}
		_spec = sqlgraph.NewCreateSpec(question.Table, sqlgraph.NewFieldSpec(question.FieldID, field.TypeUUID))
	)
	if id, ok := qc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := qc.mutation.GetType(); ok {
		_spec.SetField(question.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := qc.mutation.Body(); ok {
		_spec.SetField(question.FieldBody, field.TypeString, value)
		_node.Body = value
	}
	if value, ok := qc.mutation.Order(); ok {
		_spec.SetField(question.FieldOrder, field.TypeInt, value)
		_node.Order = value
	}
	if nodes := qc.mutation.QuestionnaireIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.QuestionnaireTable,
			Columns: []string{question.QuestionnaireColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionnaire.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.questionnaire_questions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := qc.mutation.ChoicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.ChoicesTable,
			Columns: []string{question.ChoicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(choice.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := qc.mutation.AnswersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswersTable,
			Columns: []string{question.AnswersColumn},
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

// QuestionCreateBulk is the builder for creating many Question entities in bulk.
type QuestionCreateBulk struct {
	config
	builders []*QuestionCreate
}

// Save creates the Question entities in the database.
func (qcb *QuestionCreateBulk) Save(ctx context.Context) ([]*Question, error) {
	specs := make([]*sqlgraph.CreateSpec, len(qcb.builders))
	nodes := make([]*Question, len(qcb.builders))
	mutators := make([]Mutator, len(qcb.builders))
	for i := range qcb.builders {
		func(i int, root context.Context) {
			builder := qcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*QuestionMutation)
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
					_, err = mutators[i+1].Mutate(root, qcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, qcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, qcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (qcb *QuestionCreateBulk) SaveX(ctx context.Context) []*Question {
	v, err := qcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qcb *QuestionCreateBulk) Exec(ctx context.Context) error {
	_, err := qcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qcb *QuestionCreateBulk) ExecX(ctx context.Context) {
	if err := qcb.Exec(ctx); err != nil {
		panic(err)
	}
}
