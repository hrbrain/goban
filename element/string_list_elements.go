package element

import (
	"fmt"

	"github.com/pkg/errors"
)

type StringListElements []StringListElement

func NewStringListElements(StringListElements []StringListElement) StringListElements {
	return StringListElements
}

func (se StringListElements) Len() int {
	return len(se)
}

func (se StringListElements) GetElement(index int) (Element, error) {
	if index < 0 || se.Len() <= index {
		return nil, fmt.Errorf("index out of range, index: %d", index)
	}
	return se[index], nil
}

func (se StringListElements) AddElement(e Element) (Elements, error) {
	stringListElement, ok := e.(StringListElement)
	if !ok {
		return nil, fmt.Errorf("invalid element type e: %v", e)
	}
	return append(se, stringListElement), nil
}

func (se StringListElements) Floats() ([]float64, error) {
	return nil, fmt.Errorf("string elements cannot be converted into floats")
}

func (se StringListElements) GetGroupedElement() (Element, error) {
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

func (se StringListElements) Delete() Elements {
	return StringListElements{}
}

func (se StringListElements) Append(elements2 Elements) (Elements, error) {
	StringListElements2, ok := elements2.(StringListElements)
	if !ok {
		return nil, fmt.Errorf("invalid element type elements2: %v", elements2)
	}
	return append(se, StringListElements2...), nil
}

func (se StringListElements) Slice(start int, end int) (StringListElements, error) {
	newStringListElements := make(StringListElements, len(se))
	for i, stringListElement := range se {
		newStringListElement, err := stringListElement.Slice(start, end)
		if err != nil {
			return nil, errors.Wrap(err, "failed to slice stringListElement")
		}
		newStringListElements[i] = newStringListElement
	}
	return newStringListElements, nil
}

func (se StringListElements) Join(separator string) StringElements {
	stringElements := make(StringElements, se.Len())
	for i, stringListElement := range se {
		stringElements[i] = stringListElement.Join(separator)
	}
	return stringElements
}
