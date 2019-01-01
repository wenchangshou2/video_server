package taskrunner

import (
	"log"
	"testing"
)

func TestRunner(t *testing.T) {
	d := func(dc dataChan) error {
		for i := 0; i < 30; i++ {
			dc <- i
			log.Printf("Dispatcher send:%d", i)
		}
		return nil
	}
	e := func(dc dataChan) {

	}
}
