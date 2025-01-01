package main

import (
	"context"
	"log"
	"temporal101/scratch/greeting"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}

	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "programatic-workflow",
		TaskQueue: "greeting-tasks",
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, greeting.GreetSomeone, "Thadani")
	if err != nil {
		log.Fatalln("Unable to initiate Workflow execution", err)
	}

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)

}
