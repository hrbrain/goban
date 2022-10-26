package element

import (
	"errors"
	"fmt"
	"strings"
)

type StringElement struct {
	Value  string
	IsNull bool
}

func NewStringElement(s string, isNull bool) StringElement {
	return StringElement{s, isNull}
}

// Float try to convert string element to float but always return error
func (se StringElement) Float() (float64, error) {
	return 0, fmt.Errorf("string element cannot be converted into float, se.Value: %s", se.Value)
}

// String convert string element to string
func (se StringElement) String(_ bool) (string, error) {
	if se.IsNA() {
		return "", errors.New("can't convert NA to string")
	}
	return se.Value, nil
}

// ToElements return elements which contains only one element
func (se StringElement) ToElements() Elements {
	return StringElements{se}
}

// Equal compare two elements
func (se StringElement) Equal(e Element) bool {
	// If both elements are NA, they are same no matter what value they have
	if se.IsNA() && e.IsNA() {
		return true
	}
	return se == e
}

// IsNA return true if element is null
func (se StringElement) IsNA() bool {
	return se.IsNull
}

// Split string element by separator
func (se StringElement) Split(separator string, limit int) StringListElement {
	if limit <= 1 {
		return NewStringListElement([]string{se.Value})
	}
	stringList := strings.SplitN(se.Value, separator, limit)
	return NewStringListElement(stringList)
}
