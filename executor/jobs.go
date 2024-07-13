package executor

import (
	"context"
	"errors"
	"threadpool/log"
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
	ExecuteFunc func() error
	Status      STATUS
	Msg         string
}

func NewJob(ExecuteFunc func() error) *Job {
	return &Job{
		JobId:       GenerateUUID(),
		Status:      TODO,
		ExecuteFunc: ExecuteFunc,
	}
}

func (j *Job) GetId() string {
	return j.JobId
}

func (j *Job) Execute(ctx context.Context) error {
	if j.ExecuteFunc == nil {
		return errors.New("Execute Func not provided for job")
	}
	j.Status = INPROGRESS
	err := j.ExecuteFunc()
	j.Status = SUCCESS
	if err != nil {
		j.Msg = err.Error()
		j.Status = FAILED
	}
	log.Print(log.Info, "Job Id: %s completed with status %d", j.JobId, j.Status)
	return err
}
