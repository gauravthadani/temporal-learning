package greeting

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"go.temporal.io/sdk/workflow"
)

func GreetSomeone(ctx workflow.Context, name string) (string, error) {

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	err := workflow.ExecuteActivity(ctx, GetTranslation, name).Get(ctx, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func GetTranslation(ctx context.Context, name string) (string, error) {
	// logger := workflow.GetLogger(ctx)

	path := fmt.Sprintf("http://localhost:3333/translate?name=%s", name)
	resp, err := http.Get(path)
	if err != nil {
		log.Fatal("Failed to get a response", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	translation := string(body)
	status := resp.StatusCode
	if status >= 400 {
		message := fmt.Sprintf("HTTP Error %d: %s", status, translation)
		return "", errors.New(message)
	}

	return translation, nil
}
