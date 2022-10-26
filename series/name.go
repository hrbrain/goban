package series

type Name string

func NewName(s string) Name {
	return Name(s)
}

func (n Name) String() string {
	return string(n)
}
