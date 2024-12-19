import asyncio
from datetime import timedelta
from temporalio.client import Client
from your_workflows import YourWorkflow
from temporalio.common import RetryPolicy

async def main():
    client = await Client.connect("localhost:7233")

    result = await client.execute_workflow(
        YourWorkflow.run,
        "args",
        id="workflow-id",
        task_queue="task-queue",
        # Set Workflow Timeout duration
        execution_timeout=timedelta(seconds=2),
        # run_timeout=timedelta(seconds=2),
        # task_timeout=timedelta(seconds=2),
    )
