package sched

func ExampleSched() {
	Routine()
	// Output:
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
}
