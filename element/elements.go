package element

type Elements interface {
	Len() int
	GetElement(int) (Element, error)
	AddElement(element Element) (Elements, error)
	Floats() ([]float64, error)
	GetGroupedElement() (Element, error)
	Delete() Elements
	Append(elements Elements) (Elements, error)
}
