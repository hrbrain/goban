package element

import (
	"fmt"
	"math"
)

type NumericElement struct {
	Value  float64
	IsNull bool
}

func NewNumericElement(f float64, isNull bool) NumericElement {
	return NumericElement{f, isNull}
}

// Float convert numeric element to float
func (ne NumericElement) Float() (float64, error) {
	if ne.IsNA() {
		return math.NaN(), fmt.Errorf("can't convert NaN to float")
	}
	return ne.Value, nil
}

// String convert numeric element to string
func (ne NumericElement) String() (string, error) {
	if ne.IsNA() {
		return "", nil
	}
	return fmt.Sprintf("%f", ne.Value), nil
}

// ToElements return elements which contains only one element
func (ne NumericElement) ToElements() Elements {
	return NumericElements{ne}
}

// Equal compare two elements
func (ne NumericElement) Equal(e Element) bool {
	// If both elements are NA, they are same no matter what value they have
	if ne.IsNA() && e.IsNA() {
		return true
	}
	return ne == e
}

// IsNA return true if element is NA
func (ne NumericElement) IsNA() bool {
	return ne.IsNull
}
