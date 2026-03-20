package label

type Label struct {
	Id          LabelId
	Name        string
	Description string
}

func NewLabel(name string) *Label {
	return &Label{Id: NewLabelID(), Name: name}
}
