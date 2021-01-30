package typeof

type pencil struct{}

func (p pencil) Write([]byte) (int, error) { return 0, nil }
