package errgroupexample

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/sync/errgroup"
)

type Command string

const (
	Success Command = "SUCCESS"
	Fail    Command = "FAIL"
)

var commands []Command = []Command{
	Success,
	Success,
	Success,
	Success,
	Success,
	Success,
	Success,
	Fail,
	Success,
	Success,
	Success,
	Success,
}

func fetch(workerId int, command Command) error {

	if command == Fail {
		fmt.Println("Worker: ", workerId, " : Error")
		return errors.New("Foo")
	}

	fmt.Println("Worker: ", workerId, " : Finished")
	return nil
}

func Run() {
	g := new(errgroup.Group)
	for i, comm := range commands {
		g.Go(func() error {
			err := fetch(i, comm)
			if err != nil {
				return err
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatalln("ERROR: ", err)
	}
	log.Println("No error")
}
