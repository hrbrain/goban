package element

type Element interface {
	Float() (float64, error)
	String(isForce bool) (string, error)
	ToElements() Elements
	Equal(element Element) bool
	IsNA() bool
}
