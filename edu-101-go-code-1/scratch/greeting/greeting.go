package greeting

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"

	"temporal101/scratch/models"
)

func GreetSomeone(ctx workflow.Context, request models.TranslateRequest) (models.TranslateOutput, error) {

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
		RetryPolicy:         &temporal.RetryPolicy{MaximumAttempts: 2},
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var result models.TranslateOutput
	err := workflow.ExecuteActivity(ctx, GetTranslation, request).Get(ctx, &result)
	if err != nil {
		return models.TranslateOutput{}, err
	}

	return result, nil
}

func GetTranslation(ctx context.Context, request models.TranslateRequest) (models.TranslateOutput, error) {
	// logger := workflow.GetLogger(ctx)

	path := fmt.Sprintf("http://localhost:3333/translate?name=%s&lang=%s", request.Name, request.Language)
	resp, err := http.Get(path)
	if err != nil {
		log.Fatal("Failed to get a response", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.TranslateOutput{}, err
	}
	translation := string(body)
	status := resp.StatusCode
	if status >= 400 {
		message := fmt.Sprintf("HTTP Error %d: %s", status, translation)
		return models.TranslateOutput{}, errors.New(message)
	}

	return models.TranslateOutput{
		Greeting: translation,
	}, nil
}
