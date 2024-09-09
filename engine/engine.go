package engine

import (
	"context"
)

type Engine interface {
	RegisterWorkflow(ctx context.Context) (workflowID string, err error)
	ExecuteWorkflow(ctx context.Context, workflowID string)
}

type impl struct{}

func Init() Engine {
	return &impl{}
}
