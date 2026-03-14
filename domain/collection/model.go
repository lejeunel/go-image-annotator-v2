package collection

type Collection struct {
	ID          CollectionID
	Name        string
	Description string
}

func NewCollection(name string) *Collection {
	return &Collection{ID: NewCollectionID(), Name: name}
}
