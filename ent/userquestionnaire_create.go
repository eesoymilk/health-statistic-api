// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eesoymilk/health-statistic-api/ent/answer"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaire"
	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/eesoymilk/health-statistic-api/ent/userquestionnaire"
	"github.com/google/uuid"
)

// UserQuestionnaireCreate is the builder for creating a UserQuestionnaire entity.
type UserQuestionnaireCreate struct {
	config
	mutation *UserQuestionnaireMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (uqc *UserQuestionnaireCreate) SetCreatedAt(t time.Time) *UserQuestionnaireCreate {
	uqc.mutation.SetCreatedAt(t)
	return uqc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uqc *UserQuestionnaireCreate) SetNillableCreatedAt(t *time.Time) *UserQuestionnaireCreate {
	if t != nil {
		uqc.SetCreatedAt(*t)
	}
	return uqc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (uqc *UserQuestionnaireCreate) SetUserID(id uuid.UUID) *UserQuestionnaireCreate {
	uqc.mutation.SetUserID(id)
	return uqc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (uqc *UserQuestionnaireCreate) SetNillableUserID(id *uuid.UUID) *UserQuestionnaireCreate {
	if id != nil {
		uqc = uqc.SetUserID(*id)
	}
	return uqc
}

// SetUser sets the "user" edge to the User entity.
func (uqc *UserQuestionnaireCreate) SetUser(u *User) *UserQuestionnaireCreate {
	return uqc.SetUserID(u.ID)
}

// SetQuestionnaireID sets the "questionnaire" edge to the Questionnaire entity by ID.
func (uqc *UserQuestionnaireCreate) SetQuestionnaireID(id int) *UserQuestionnaireCreate {
	uqc.mutation.SetQuestionnaireID(id)
	return uqc
}

// SetNillableQuestionnaireID sets the "questionnaire" edge to the Questionnaire entity by ID if the given value is not nil.
func (uqc *UserQuestionnaireCreate) SetNillableQuestionnaireID(id *int) *UserQuestionnaireCreate {
	if id != nil {
		uqc = uqc.SetQuestionnaireID(*id)
	}
	return uqc
}

// SetQuestionnaire sets the "questionnaire" edge to the Questionnaire entity.
func (uqc *UserQuestionnaireCreate) SetQuestionnaire(q *Questionnaire) *UserQuestionnaireCreate {
	return uqc.SetQuestionnaireID(q.ID)
}

// AddAnswerIDs adds the "answers" edge to the Answer entity by IDs.
func (uqc *UserQuestionnaireCreate) AddAnswerIDs(ids ...int) *UserQuestionnaireCreate {
	uqc.mutation.AddAnswerIDs(ids...)
	return uqc
}

// AddAnswers adds the "answers" edges to the Answer entity.
func (uqc *UserQuestionnaireCreate) AddAnswers(a ...*Answer) *UserQuestionnaireCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uqc.AddAnswerIDs(ids...)
}

// Mutation returns the UserQuestionnaireMutation object of the builder.
func (uqc *UserQuestionnaireCreate) Mutation() *UserQuestionnaireMutation {
	return uqc.mutation
}

// Save creates the UserQuestionnaire in the database.
func (uqc *UserQuestionnaireCreate) Save(ctx context.Context) (*UserQuestionnaire, error) {
	uqc.defaults()
	return withHooks(ctx, uqc.sqlSave, uqc.mutation, uqc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uqc *UserQuestionnaireCreate) SaveX(ctx context.Context) *UserQuestionnaire {
	v, err := uqc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uqc *UserQuestionnaireCreate) Exec(ctx context.Context) error {
	_, err := uqc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uqc *UserQuestionnaireCreate) ExecX(ctx context.Context) {
	if err := uqc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uqc *UserQuestionnaireCreate) defaults() {
	if _, ok := uqc.mutation.CreatedAt(); !ok {
		v := userquestionnaire.DefaultCreatedAt()
		uqc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uqc *UserQuestionnaireCreate) check() error {
	if _, ok := uqc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserQuestionnaire.created_at"`)}
	}
	return nil
}

func (uqc *UserQuestionnaireCreate) sqlSave(ctx context.Context) (*UserQuestionnaire, error) {
	if err := uqc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uqc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uqc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uqc.mutation.id = &_node.ID
	uqc.mutation.done = true
	return _node, nil
}

func (uqc *UserQuestionnaireCreate) createSpec() (*UserQuestionnaire, *sqlgraph.CreateSpec) {
	var (
		_node = &UserQuestionnaire{config: uqc.config}
		_spec = sqlgraph.NewCreateSpec(userquestionnaire.Table, sqlgraph.NewFieldSpec(userquestionnaire.FieldID, field.TypeInt))
	)
	if value, ok := uqc.mutation.CreatedAt(); ok {
		_spec.SetField(userquestionnaire.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := uqc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userquestionnaire.UserTable,
			Columns: []string{userquestionnaire.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_questionnaires = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uqc.mutation.QuestionnaireIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userquestionnaire.QuestionnaireTable,
			Columns: []string{userquestionnaire.QuestionnaireColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionnaire.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.questionnaire_responses = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uqc.mutation.AnswersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userquestionnaire.AnswersTable,
			Columns: []string{userquestionnaire.AnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserQuestionnaireCreateBulk is the builder for creating many UserQuestionnaire entities in bulk.
type UserQuestionnaireCreateBulk struct {
	config
	builders []*UserQuestionnaireCreate
}

// Save creates the UserQuestionnaire entities in the database.
func (uqcb *UserQuestionnaireCreateBulk) Save(ctx context.Context) ([]*UserQuestionnaire, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uqcb.builders))
	nodes := make([]*UserQuestionnaire, len(uqcb.builders))
	mutators := make([]Mutator, len(uqcb.builders))
	for i := range uqcb.builders {
		func(i int, root context.Context) {
			builder := uqcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserQuestionnaireMutation)
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
					_, err = mutators[i+1].Mutate(root, uqcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uqcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, uqcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uqcb *UserQuestionnaireCreateBulk) SaveX(ctx context.Context) []*UserQuestionnaire {
	v, err := uqcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uqcb *UserQuestionnaireCreateBulk) Exec(ctx context.Context) error {
	_, err := uqcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uqcb *UserQuestionnaireCreateBulk) ExecX(ctx context.Context) {
	if err := uqcb.Exec(ctx); err != nil {
		panic(err)
	}
}
