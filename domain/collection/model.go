package collection

type Collection struct {
	Id          CollectionId
	Name        string
	Description string
}

func NewCollection(name string) *Collection {
	return &Collection{Id: NewCollectionID(), Name: name}
}
