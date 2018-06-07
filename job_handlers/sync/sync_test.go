package sync_test

import (
	"fmt"
	"testing"
	"distributed-programming-abstractions"
	jh "distributed-programming-abstractions/job_handlers/sync"
)

func TestRun(t *test.T) {
	var msg = make(chan string)
	jh = jh.New()
	jh.Confirm(func(j job.Job) {
		msg <- fmt.Sprintln("Confirmed job: ", j)
	})
	jh.Process(func(j job.Job) {
		fmt.Println("Processing job: ", j)
	})
	jh.Submit("Job 1")
	jh.Submit("Job 2")
	jh.Submit("Job 3")

	fmt.Println(<-msg)
	fmt.Println(<-msg)
	fmt.Println(<-msg)
}