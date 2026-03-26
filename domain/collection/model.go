package collection

type Collection struct {
	Id          CollectionId
	Name        string
	Description string
}

func NewCollection(name string, opts ...Option) *Collection {
	c := &Collection{Id: NewCollectionId(), Name: name}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

type Option func(*Collection)

func WithDescription(d string) Option {
	return func(c *Collection) {
		c.Description = d
	}
}
