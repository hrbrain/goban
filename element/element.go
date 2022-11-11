package element

type Element interface {
	Float() (float64, error)
	String() (string, error)
	ToElements() Elements
	Equal(element Element) bool
	IsNA() bool
}
