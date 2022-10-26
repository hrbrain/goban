package element

import (
	"fmt"

	"github.com/pkg/errors"
)

type NumericElements []NumericElement

func NewNumericElements(numericElements []NumericElement) NumericElements {
	return numericElements
}

func (ne NumericElements) Len() int {
	return len(ne)
}

// GetElement returns element at index
func (ne NumericElements) GetElement(index int) (Element, error) {
	if index < 0 || ne.Len() <= index {
		return nil, fmt.Errorf("index out of range, index: %d", index)
	}
	return ne[index], nil
}

// AddElement adds element to NumericElements
func (ne NumericElements) AddElement(e Element) (Elements, error) {
	numericElement, ok := e.(NumericElement)
	if !ok {
		return nil, fmt.Errorf("invalid element type e: %v", e)
	}
	return append(ne, numericElement), nil
}

// Floats convert numeric elements into float64 slice
func (ne NumericElements) Floats() ([]float64, error) {
	floats := make([]float64, ne.Len())
	for i := 0; i < ne.Len(); i++ {
		element, err := ne.GetElement(i)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get element")
		}
		floatValue, err := element.Float()
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert to float value")
		}
		floats[i] = floatValue
	}
	return floats, nil
}

// GetGroupedElement returns single element if all elements are same
func (ne NumericElements) GetGroupedElement() (Element, error) {
	var groupedElement Element
	for i, element := range ne {
		if i == 0 {
			groupedElement = element
			continue
		}
		if !groupedElement.Equal(element) {
			err := fmt.Errorf("elements are not grouped, groupedElement: %v, element: %v", groupedElement, element)
			return nil, err
		}
	}
	return groupedElement, nil
}

// Delete returns empty NumericElements
func (ne NumericElements) Delete() Elements {
	return NumericElements{}
}

// Append appends elements to NumericElements
func (ne NumericElements) Append(elements2 Elements) (Elements, error) {
	numericElements2, ok := elements2.(NumericElements)
	if !ok {
		return nil, fmt.Errorf("invalid element type elements2: %v", elements2)
	}
	return append(ne, numericElements2...), nil
}
