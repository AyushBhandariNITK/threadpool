package executor

import (
	"context"
	"errors"
)

type STATUS int

const (
	TODO STATUS = iota
	INPROGRESS
	SUCCESS
	FAILED
)

type Job struct {
	JobId       string
	ExecuteFunc func(context.Context) error
	Status      STATUS
	Msg         string
}

func NewJob() *Job {
	return &Job{
		JobId:  GenerateUUID(),
		Status: TODO,
	}
}

func (j *Job) Execute(ctx context.Context) (err error) {
	if j.ExecuteFunc == nil {
		return errors.New("Execute Func not provided for job")
	}
	j.Status = INPROGRESS
	err = j.ExecuteFunc(ctx)
	j.Status = SUCCESS
	if err != nil {
		j.Msg = err.Error()
		j.Status = FAILED
	}
	return err
}
