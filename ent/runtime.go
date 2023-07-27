// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/eesoymilk/health-statistic-api/ent/answer"
	"github.com/eesoymilk/health-statistic-api/ent/choice"
	"github.com/eesoymilk/health-statistic-api/ent/mycard"
	"github.com/eesoymilk/health-statistic-api/ent/notification"
	"github.com/eesoymilk/health-statistic-api/ent/price"
	"github.com/eesoymilk/health-statistic-api/ent/question"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaire"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaireresponse"
	"github.com/eesoymilk/health-statistic-api/ent/schema"
	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	answerFields := schema.Answer{}.Fields()
	_ = answerFields
	// answerDescCreatedAt is the schema descriptor for created_at field.
	answerDescCreatedAt := answerFields[1].Descriptor()
	// answer.DefaultCreatedAt holds the default value on creation for the created_at field.
	answer.DefaultCreatedAt = answerDescCreatedAt.Default.(func() time.Time)
	// answerDescID is the schema descriptor for id field.
	answerDescID := answerFields[0].Descriptor()
	// answer.DefaultID holds the default value on creation for the id field.
	answer.DefaultID = answerDescID.Default.(func() uuid.UUID)
	choiceFields := schema.Choice{}.Fields()
	_ = choiceFields
	// choiceDescBody is the schema descriptor for body field.
	choiceDescBody := choiceFields[1].Descriptor()
	// choice.BodyValidator is a validator for the "body" field. It is called by the builders before save.
	choice.BodyValidator = choiceDescBody.Validators[0].(func(string) error)
	// choiceDescOrder is the schema descriptor for order field.
	choiceDescOrder := choiceFields[2].Descriptor()
	// choice.OrderValidator is a validator for the "order" field. It is called by the builders before save.
	choice.OrderValidator = choiceDescOrder.Validators[0].(func(int) error)
	// choiceDescID is the schema descriptor for id field.
	choiceDescID := choiceFields[0].Descriptor()
	// choice.DefaultID holds the default value on creation for the id field.
	choice.DefaultID = choiceDescID.Default.(func() uuid.UUID)
	mycardFields := schema.MyCard{}.Fields()
	_ = mycardFields
	// mycardDescCardPassword is the schema descriptor for card_password field.
	mycardDescCardPassword := mycardFields[1].Descriptor()
	// mycard.CardPasswordValidator is a validator for the "card_password" field. It is called by the builders before save.
	mycard.CardPasswordValidator = func() func(string) error {
		validators := mycardDescCardPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(card_password string) error {
			for _, fn := range fns {
				if err := fn(card_password); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// mycardDescCreatedAt is the schema descriptor for created_at field.
	mycardDescCreatedAt := mycardFields[2].Descriptor()
	// mycard.DefaultCreatedAt holds the default value on creation for the created_at field.
	mycard.DefaultCreatedAt = mycardDescCreatedAt.Default.(func() time.Time)
	// mycardDescID is the schema descriptor for id field.
	mycardDescID := mycardFields[0].Descriptor()
	// mycard.IDValidator is a validator for the "id" field. It is called by the builders before save.
	mycard.IDValidator = func() func(string) error {
		validators := mycardDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	notificationFields := schema.Notification{}.Fields()
	_ = notificationFields
	// notificationDescID is the schema descriptor for id field.
	notificationDescID := notificationFields[0].Descriptor()
	// notification.DefaultID holds the default value on creation for the id field.
	notification.DefaultID = notificationDescID.Default.(func() uuid.UUID)
	priceFields := schema.Price{}.Fields()
	_ = priceFields
	// priceDescName is the schema descriptor for name field.
	priceDescName := priceFields[1].Descriptor()
	// price.NameValidator is a validator for the "name" field. It is called by the builders before save.
	price.NameValidator = priceDescName.Validators[0].(func(string) error)
	// priceDescDescription is the schema descriptor for description field.
	priceDescDescription := priceFields[2].Descriptor()
	// price.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	price.DescriptionValidator = priceDescDescription.Validators[0].(func(string) error)
	// priceDescCreatedAt is the schema descriptor for created_at field.
	priceDescCreatedAt := priceFields[3].Descriptor()
	// price.DefaultCreatedAt holds the default value on creation for the created_at field.
	price.DefaultCreatedAt = priceDescCreatedAt.Default.(func() time.Time)
	// priceDescID is the schema descriptor for id field.
	priceDescID := priceFields[0].Descriptor()
	// price.DefaultID holds the default value on creation for the id field.
	price.DefaultID = priceDescID.Default.(func() uuid.UUID)
	questionFields := schema.Question{}.Fields()
	_ = questionFields
	// questionDescBody is the schema descriptor for body field.
	questionDescBody := questionFields[2].Descriptor()
	// question.BodyValidator is a validator for the "body" field. It is called by the builders before save.
	question.BodyValidator = questionDescBody.Validators[0].(func(string) error)
	// questionDescOrder is the schema descriptor for order field.
	questionDescOrder := questionFields[3].Descriptor()
	// question.OrderValidator is a validator for the "order" field. It is called by the builders before save.
	question.OrderValidator = questionDescOrder.Validators[0].(func(int) error)
	// questionDescID is the schema descriptor for id field.
	questionDescID := questionFields[0].Descriptor()
	// question.DefaultID holds the default value on creation for the id field.
	question.DefaultID = questionDescID.Default.(func() uuid.UUID)
	questionnaireFields := schema.Questionnaire{}.Fields()
	_ = questionnaireFields
	// questionnaireDescCreatedAt is the schema descriptor for created_at field.
	questionnaireDescCreatedAt := questionnaireFields[2].Descriptor()
	// questionnaire.DefaultCreatedAt holds the default value on creation for the created_at field.
	questionnaire.DefaultCreatedAt = questionnaireDescCreatedAt.Default.(func() time.Time)
	// questionnaireDescID is the schema descriptor for id field.
	questionnaireDescID := questionnaireFields[0].Descriptor()
	// questionnaire.DefaultID holds the default value on creation for the id field.
	questionnaire.DefaultID = questionnaireDescID.Default.(func() uuid.UUID)
	questionnaireresponseFields := schema.QuestionnaireResponse{}.Fields()
	_ = questionnaireresponseFields
	// questionnaireresponseDescCreatedAt is the schema descriptor for created_at field.
	questionnaireresponseDescCreatedAt := questionnaireresponseFields[1].Descriptor()
	// questionnaireresponse.DefaultCreatedAt holds the default value on creation for the created_at field.
	questionnaireresponse.DefaultCreatedAt = questionnaireresponseDescCreatedAt.Default.(func() time.Time)
	// questionnaireresponseDescID is the schema descriptor for id field.
	questionnaireresponseDescID := questionnaireresponseFields[0].Descriptor()
	// questionnaireresponse.DefaultID holds the default value on creation for the id field.
	questionnaireresponse.DefaultID = questionnaireresponseDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[2].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescBirthYear is the schema descriptor for birth_year field.
	userDescBirthYear := userFields[3].Descriptor()
	// user.BirthYearValidator is a validator for the "birth_year" field. It is called by the builders before save.
	user.BirthYearValidator = func() func(int) error {
		validators := userDescBirthYear.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(birth_year int) error {
			for _, fn := range fns {
				if err := fn(birth_year); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescHeight is the schema descriptor for height field.
	userDescHeight := userFields[4].Descriptor()
	// user.HeightValidator is a validator for the "height" field. It is called by the builders before save.
	user.HeightValidator = userDescHeight.Validators[0].(func(float64) error)
	// userDescWeight is the schema descriptor for weight field.
	userDescWeight := userFields[5].Descriptor()
	// user.WeightValidator is a validator for the "weight" field. It is called by the builders before save.
	user.WeightValidator = userDescWeight.Validators[0].(func(float64) error)
}
