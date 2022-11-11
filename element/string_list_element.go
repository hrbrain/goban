package element

import (
	"fmt"
	"reflect"
	"strings"
)

type StringListElement []string

func NewStringListElement(stringList []string) StringListElement {
	return stringList
}

func (sle StringListElement) Float() (float64, error) {
	return 0, fmt.Errorf("string list element cannot be converted into float, sle: %s", sle)
}

func (sle StringListElement) String() (string, error) {
	return "", fmt.Errorf("string list element cannot be converted into string, sle: %s", sle)
}

func (sle StringListElement) ToElements() Elements {
	return StringListElements{sle}
}

func (sle StringListElement) Equal(e2 Element) bool {
	// If both elements are NA, they are same no matter what value they have
	if sle.IsNA() && e2.IsNA() {
		return true
	}
	return reflect.DeepEqual(sle, e2)
}

func (sle StringListElement) IsNA() bool {
	return len(sle) == 0
}

func (sle StringListElement) Slice(start int, end int) (StringListElement, error) {
	if start < 0 || start > end {
		return StringListElement{}, fmt.Errorf("invalid start: %d, end: %d", start, end)
	}
	if end > len(sle) {
		return sle, nil
	}
	return sle[start:end], nil
}

func (sle StringListElement) Join(separator string) StringElement {
	joinedStr := strings.Join(sle, separator)
	return NewStringElement(joinedStr, false)
}
