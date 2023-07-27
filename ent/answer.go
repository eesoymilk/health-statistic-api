// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/eesoymilk/health-statistic-api/ent/answer"
	"github.com/eesoymilk/health-statistic-api/ent/question"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaireresponse"
	"github.com/google/uuid"
)

// Answer is the model entity for the Answer schema.
type Answer struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Body holds the value of the "body" field.
	Body *string `json:"body,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AnswerQuery when eager-loading is set.
	Edges                          AnswerEdges `json:"-"`
	question_answers               *uuid.UUID
	questionnaire_response_answers *uuid.UUID
	selectValues                   sql.SelectValues
}

// AnswerEdges holds the relations/edges for other nodes in the graph.
type AnswerEdges struct {
	// Chosen holds the value of the chosen edge.
	Chosen []*Choice `json:"chosen,omitempty"`
	// Question holds the value of the question edge.
	Question *Question `json:"question,omitempty"`
	// QuestionnaireResponse holds the value of the questionnaire_response edge.
	QuestionnaireResponse *QuestionnaireResponse `json:"questionnaire_response,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ChosenOrErr returns the Chosen value or an error if the edge
// was not loaded in eager-loading.
func (e AnswerEdges) ChosenOrErr() ([]*Choice, error) {
	if e.loadedTypes[0] {
		return e.Chosen, nil
	}
	return nil, &NotLoadedError{edge: "chosen"}
}

// QuestionOrErr returns the Question value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AnswerEdges) QuestionOrErr() (*Question, error) {
	if e.loadedTypes[1] {
		if e.Question == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: question.Label}
		}
		return e.Question, nil
	}
	return nil, &NotLoadedError{edge: "question"}
}

// QuestionnaireResponseOrErr returns the QuestionnaireResponse value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AnswerEdges) QuestionnaireResponseOrErr() (*QuestionnaireResponse, error) {
	if e.loadedTypes[2] {
		if e.QuestionnaireResponse == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: questionnaireresponse.Label}
		}
		return e.QuestionnaireResponse, nil
	}
	return nil, &NotLoadedError{edge: "questionnaire_response"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Answer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case answer.FieldBody:
			values[i] = new(sql.NullString)
		case answer.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case answer.FieldID:
			values[i] = new(uuid.UUID)
		case answer.ForeignKeys[0]: // question_answers
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case answer.ForeignKeys[1]: // questionnaire_response_answers
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Answer fields.
func (a *Answer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case answer.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case answer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case answer.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				a.Body = new(string)
				*a.Body = value.String
			}
		case answer.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field question_answers", values[i])
			} else if value.Valid {
				a.question_answers = new(uuid.UUID)
				*a.question_answers = *value.S.(*uuid.UUID)
			}
		case answer.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field questionnaire_response_answers", values[i])
			} else if value.Valid {
				a.questionnaire_response_answers = new(uuid.UUID)
				*a.questionnaire_response_answers = *value.S.(*uuid.UUID)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Answer.
// This includes values selected through modifiers, order, etc.
func (a *Answer) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryChosen queries the "chosen" edge of the Answer entity.
func (a *Answer) QueryChosen() *ChoiceQuery {
	return NewAnswerClient(a.config).QueryChosen(a)
}

// QueryQuestion queries the "question" edge of the Answer entity.
func (a *Answer) QueryQuestion() *QuestionQuery {
	return NewAnswerClient(a.config).QueryQuestion(a)
}

// QueryQuestionnaireResponse queries the "questionnaire_response" edge of the Answer entity.
func (a *Answer) QueryQuestionnaireResponse() *QuestionnaireResponseQuery {
	return NewAnswerClient(a.config).QueryQuestionnaireResponse(a)
}

// Update returns a builder for updating this Answer.
// Note that you need to call Answer.Unwrap() before calling this method if this Answer
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Answer) Update() *AnswerUpdateOne {
	return NewAnswerClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Answer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Answer) Unwrap() *Answer {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Answer is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Answer) String() string {
	var builder strings.Builder
	builder.WriteString("Answer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := a.Body; v != nil {
		builder.WriteString("body=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (a *Answer) MarshalJSON() ([]byte, error) {
	type Alias Answer
	return json.Marshal(&struct {
		*Alias
		AnswerEdges
	}{
		Alias:       (*Alias)(a),
		AnswerEdges: a.Edges,
	})
}

// Answers is a parsable slice of Answer.
type Answers []*Answer
