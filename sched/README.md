# Goroutine Schedule

```go
	// ready to say foo
	// ready to say bar
	// foo  |  10
	// ready to say foo
	// bar  |  10
	// ready to say bar
	// foo  |  11
	// ready to say foo
	// bar  |  11
	// ready to say bar
	// foo  |  12
```

## Remark
we call a single function called Routine in every goroutine.
The function has two stages: fisrt stage prepare to say its 
arguments; second stage actually say its arguments. Two stages
are sperated by a golang primitives of switching routines, which
means the context shall be saved the recover later in certain 
moment. We declare and initialize a local variable *a* as contextual.
