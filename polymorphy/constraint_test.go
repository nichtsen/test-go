package polymorphy

import (
	"testing"
)

func testHandle(args ...any) error {
	return nil
}

var bar = BarExecutor{}
var foo = &FooExecutor{}

func TestRun(t *testing.T) {
	Run[Executor](foo)
	Exec[Handler](testHandle)
}
