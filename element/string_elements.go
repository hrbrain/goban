package element

import (
	"fmt"
)

type StringElements []StringElement

func NewStringElements(stringElements []StringElement) StringElements {
	return stringElements
}

// Len returns the length of the elements.
func (se StringElements) Len() int {
	return len(se)
}

// GetElement returns the element at the given index.
func (se StringElements) GetElement(index int) (Element, error) {
	if index < 0 || se.Len() <= index {
		return nil, fmt.Errorf("index out of range, index: %d", index)
	}
	return se[index], nil
}

// AddElement adds the given element to the elements.
func (se StringElements) AddElement(e Element) (Elements, error) {
	stringElement, ok := e.(StringElement)
	if !ok {
		return nil, fmt.Errorf("invalid element type e: %v", e)
	}
	return append(se, stringElement), nil
}

// Floats trys to return the elements as a slice of float64 but always returns an error.
func (se StringElements) Floats() ([]float64, error) {
	return nil, fmt.Errorf("string elements cannot be converted into floats")
}

// GetGroupedElement returns a single element if all the elements are the same.
func (se StringElements) GetGroupedElement() (Element, error) {
	var groupedElement Element
	for i, element := range se {
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

// Delete returns an empty elements.
func (se StringElements) Delete() Elements {
	return StringElements{}
}

// Append the given elements to the elements.
func (se StringElements) Append(elements2 Elements) (Elements, error) {
	stringElements2, ok := elements2.(StringElements)
	if !ok {
		return nil, fmt.Errorf("invalid element type elements2: %v", elements2)
	}
	return append(se, stringElements2...), nil
}

// Split string in each element into a string list and return a list of string list elements.
func (se StringElements) Split(separator string, limit int) (StringListElements, error) {
	stringListElements := make(StringListElements, se.Len())
	for i, stringElement := range se {
		stringListElements[i] = stringElement.Split(separator, limit)
	}
	return stringListElements, nil
}
