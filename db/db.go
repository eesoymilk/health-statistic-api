package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/eesoymilk/health-statistic-api/ent/migrate"
	"github.com/eesoymilk/health-statistic-api/ent/question"
	"github.com/eesoymilk/health-statistic-api/types"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func ReadRegistrationQuestionnaire() (*types.QuestionnaireWithId, error) {
	data, err := os.ReadFile("./db/registration_questionnaire.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	// Parse the JSON
	var questionnaireData types.QuestionnaireWithId
	err = json.Unmarshal(data, &questionnaireData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	// Print the json
	jsonBytes, err := json.MarshalIndent(questionnaireData, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error marshaling to JSON: %v", err)
	}
	log.Printf("data: %v", string(jsonBytes))

	return &questionnaireData, nil
}

func CreateRegistrationQuestionnaire(
	ctx context.Context,
	db *ent.Client,
) error {
	// Parse the JSON
	questionnaireData, err := ReadRegistrationQuestionnaire()
	if err != nil {
		return err
	}

	id, err := uuid.Parse(questionnaireData.ID)
	if err != nil {
		return fmt.Errorf("failed parsing id: %v", err)
	}

	questionnaireNode, err := db.Questionnaire.Create().
		SetID(id).
		SetName(questionnaireData.Name).
		Save(ctx)
	if err != nil {
		// This error occurs if registration questionnaire is already created
		return fmt.Errorf(
			"failed to create registration questionnaire: %v",
			err,
		)
	}

	for i, questionData := range questionnaireData.Questions {
		questionNode, err := db.Question.Create().
			SetType(question.Type(questionData.Type)).
			SetBody(questionData.Body).
			SetOrder(i).
			SetQuestionnaire(questionnaireNode).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create question: %v", err)
		}

		for i, choice := range questionData.Choices {
			log.Default().Printf("creating options: %v", choice)
			_, err := db.Choice.Create().
				SetBody(choice).
				SetQuesion(questionNode).
				SetOrder(i).
				Save(ctx)
			if err != nil {
				return fmt.Errorf("failed to create choice: %v", err)
			}
		}
	}

	return nil
}

func Migrate(ctx context.Context, db *ent.Client) error {
	if err := db.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		return fmt.Errorf(
			"failed creating schema resources: %v",
			err,
		)
	}

	// Parse the JSON
	err := CreateRegistrationQuestionnaire(ctx, db)
	if err != nil {
		return fmt.Errorf(
			"failed creating registration questionnaire: %v",
			err,
		)
	}

	return nil
}

func New() *ent.Client {
	var (
		postgresHost     = os.Getenv("POSTGRES_HOST")
		postgresPort     = os.Getenv("POSTGRES_PORT")
		postgresUser     = os.Getenv("POSTGRES_USER")
		postgresDb       = os.Getenv("POSTGRES_DB")
		postgresPassword = os.Getenv("POSTGRES_PASSWORD")
	)

	connectionStr := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		postgresHost,
		postgresPort,
		postgresUser,
		postgresDb,
		postgresPassword,
	)

	log.Println(connectionStr)

	db, err := ent.Open("postgres", connectionStr)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	return db
}
