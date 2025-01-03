package main

import (
	"context"
	"log"
	"temporal101/scratch/greeting"
	"temporal101/scratch/models"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}

	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:          "programatic-workflow",
		TaskQueue:   "greeting-tasks",
		RetryPolicy: &temporal.RetryPolicy{},
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, greeting.GreetSomeone, models.TranslateRequest{
		Name:     "Thadani",
		Language: "french",
	})
	if err != nil {
		log.Fatalln("Unable to initiate Workflow execution", err)
	}

	var result models.TranslateOutput
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Printf("Workflow result: %s", result.Greeting)

}
