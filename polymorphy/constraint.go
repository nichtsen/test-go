package polymorphy

import "fmt"

type Executor interface {
	Run()
}
type Stopor interface {
	Stop()
}

type Handler func(args ...any) error
type ExecutorHandle func(fun string, args int) error

type FooExecutor struct {
	Name string
}

func (f *FooExecutor) Run() {
	fmt.Println("foo executor is running")
}

type BarExecutor struct {
	FooExecutor
}

type runner interface {
	FooExecutor | Handler
}

type T1[T []interface{ Executor }] string

func Run[T Executor](r T) {
	fmt.Printf("\n%T", r)
}
func Exec[T Handler | ExecutorHandle](r ...T) {
	fmt.Printf("\n%T", r)
}

func baz() {
	type s string
	var a s = "a"
	fmt.Println(a)
}
