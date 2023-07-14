// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eesoymilk/health-statistic-api/ent/predicate"
	"github.com/eesoymilk/health-statistic-api/ent/userquestionnaire"
)

// UserQuestionnaireDelete is the builder for deleting a UserQuestionnaire entity.
type UserQuestionnaireDelete struct {
	config
	hooks    []Hook
	mutation *UserQuestionnaireMutation
}

// Where appends a list predicates to the UserQuestionnaireDelete builder.
func (uqd *UserQuestionnaireDelete) Where(ps ...predicate.UserQuestionnaire) *UserQuestionnaireDelete {
	uqd.mutation.Where(ps...)
	return uqd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uqd *UserQuestionnaireDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, uqd.sqlExec, uqd.mutation, uqd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (uqd *UserQuestionnaireDelete) ExecX(ctx context.Context) int {
	n, err := uqd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uqd *UserQuestionnaireDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userquestionnaire.Table, sqlgraph.NewFieldSpec(userquestionnaire.FieldID, field.TypeInt))
	if ps := uqd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, uqd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	uqd.mutation.done = true
	return affected, err
}

// UserQuestionnaireDeleteOne is the builder for deleting a single UserQuestionnaire entity.
type UserQuestionnaireDeleteOne struct {
	uqd *UserQuestionnaireDelete
}

// Where appends a list predicates to the UserQuestionnaireDelete builder.
func (uqdo *UserQuestionnaireDeleteOne) Where(ps ...predicate.UserQuestionnaire) *UserQuestionnaireDeleteOne {
	uqdo.uqd.mutation.Where(ps...)
	return uqdo
}

// Exec executes the deletion query.
func (uqdo *UserQuestionnaireDeleteOne) Exec(ctx context.Context) error {
	n, err := uqdo.uqd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userquestionnaire.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uqdo *UserQuestionnaireDeleteOne) ExecX(ctx context.Context) {
	if err := uqdo.Exec(ctx); err != nil {
		panic(err)
	}
}
