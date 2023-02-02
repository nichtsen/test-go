package polymorphy

type Client struct{}

func (c *Client) Run(args ...string) {
	for _, v := range args {
		if l := GetLighter(v); l != nil {
			l.shining()
		}
	}
}
