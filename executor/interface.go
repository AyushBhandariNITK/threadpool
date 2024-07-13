package executor

import "context"

type Task interface {
	Execute(context.Context) error
}

type Pool interface {
	Submit(context.Context, Job) error
}
