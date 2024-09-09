package main

import (
	"bitbucket.org/shisheodev/on-demand-automation/engine"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("Workflow automation project initialised!")

	e := engine.Init()

	e.ExecuteWorkflow(ctx, "workflowID")
}
