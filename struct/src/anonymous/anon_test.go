package anonymous

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestFoo(t *testing.T) {
	foo := new(Foo)
	foo.Bar()
}

func TestJob(t *testing.T) {
	logger := log.New(os.Stdout, "JOB: ", log.Ldate|log.Ltime)
	job := &Job{
		Command: "some Command",
		Logger:  logger,
	}
	job.Start()
	fmt.Println(job.Command)
}
