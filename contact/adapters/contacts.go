package adapters

type ContactHandler struct {
	data []string
}

func (c *ContactHandler) Get() []string {
	return c.data
}

func (c *ContactHandler) Save(s string) string {
	c.data = append(c.data, s)
	return c.data[len(c.data)-1]
}
