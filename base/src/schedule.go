// dacx
package main

import (
	"context"
	"log"
	"time"

	// schedule "github.com/temporalio/documentation-samples-go/feature/schedules"
	wf "temporal_learning/src/helloworld"

	"go.temporal.io/sdk/client"
)

func main() {
	ctx := context.Background()
	temporalClient, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal Client", err)
	}
	defer temporalClient.Close()

		// Create Schedule and Workflow IDs
		scheduleID := "schedule_id1"
		workflowID := "schedule_workflow_id"

		// Create the schedule.
		scheduleHandle, err := temporalClient.ScheduleClient().Create(ctx, client.ScheduleOptions{
			ID: scheduleID,
			Spec: client.ScheduleSpec{
				Intervals: []client.ScheduleIntervalSpec{
					{
						Every:  time.Minute * 10,
						Offset: time.Minute * time.Duration(time.Now().Minute()%10),
					},
				},
			},
			TriggerImmediately: true,
			Action: &client.ScheduleWorkflowAction{
				ID:        workflowID,
				Workflow:  wf.Workflow,
				TaskQueue: "schedule",
				Args:      []interface{}{"Test"},
			},
		})
		if err != nil {
		log.Fatalln("Unable to create schedule", err)
	}
	log.Println("Schedule created", "ScheduleID", scheduleID)
	_, _ = scheduleHandle.Describe(ctx)
}

/*
Schedules are initiated with the `create` call.
The user generates a unique Schedule ID for each new Schedule.

To create a Schedule in Go, use `Create()` on the [Client](/concepts/what-is-a-temporal-client).
Schedules must be initialized with a Schedule ID, [Spec](/concepts/what-is-a-schedule#spec), and [Action](/concepts/what-is-a-schedule#action) in `client.ScheduleOptions{}`.
*/

/* @dacx
id: how-to-create-a-schedule-in-go
title: How to create a Schedule in Go
label: Create Schedule
description: To create a Schedule in Go, use `Create()` on the Client.
tags:
- go sdk
- code sample
- schedule
- create
lines: 12, 23-34, 40, 42-48
@dacx */
