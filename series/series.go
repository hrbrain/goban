package series

import (
	"fmt"

	"github.com/hrbrain/goban/element"
	"github.com/pkg/errors"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
)

// Series list of same type of elements
type Series struct {
	Name             Name
	Elements         element.Elements
	AggregatedMethod AggregationMethod
}

func NewSeries(name Name, elements element.Elements, aggregatedMethod AggregationMethod) (Series, error) {
	if elements == nil {
		return Series{}, errors.New("nil elements are not allowed")
	}
	return Series{
		Name:             name,
		Elements:         elements,
		AggregatedMethod: aggregatedMethod,
	}, nil
}

func (s Series) GetName() Name {
	return s.Name
}

func (s Series) GetType() Type {
	switch s.Elements.(type) {
	case element.StringElements:
		return StringType
	case element.NumericElements:
		return NumericType
	case element.StringListElements:
		return StringListType
	}
	return UnknownType
}

func (s Series) Len() int {
	return s.Elements.Len()
}

func (s Series) GetElement(index int) (element.Element, error) {
	return s.Elements.GetElement(index)
}

// GetGroupedElement returns a single element if all elements in the series are same
func (s Series) GetGroupedElement() (element.Element, error) {
	return s.Elements.GetGroupedElement()
}

func (s Series) UpdateElements(elements element.Elements) (Series, error) {
	if elements == nil {
		return Series{}, errors.New("nil elements are not allowed")
	}
	s.Elements = elements
	return s, nil
}

func (s Series) AddElement(element element.Element) (Series, error) {
	elements, err := s.Elements.AddElement(element)
	if err != nil {
		return Series{}, errors.Wrap(err, "failed to add element")
	}
	return s.UpdateElements(elements)
}

// Floats convert elements to float64 slice
func (s Series) Floats() ([]float64, error) {
	return s.Elements.Floats()
}

// Mean Calculate the mean of all elements
func (s Series) Mean() (float64, error) {
	floatsNums, err := s.Floats()
	if err != nil {
		return 0, errors.Wrap(err, "failed to convert to floats")
	}
	return stat.Mean(floatsNums, nil), nil
}

// Sum Calculate the sum of all elements
func (s Series) Sum() (float64, error) {
	floatsNums, err := s.Floats()
	if err != nil {
		return 0, errors.Wrap(err, "failed to convert to floats")
	}
	return floats.Sum(floatsNums), nil
}

// Split string elements by separator
func (s Series) Split(separator string, limit int) (Series, error) {
	stringElements, ok := s.Elements.(element.StringElements)
	if !ok {
		return Series{}, errors.New("this series is not string elements")
	}
	stringListElements, err := stringElements.Split(separator, limit)
	if err != nil {
		return Series{}, errors.Wrap(err, "failed to split elements")
	}
	return s.UpdateElements(stringListElements)
}

// CanAggregateWith Check if the series can be aggregated with the method
func (s Series) CanAggregateWith(method AggregationMethod) error {
	switch s.GetType() {
	case StringType:
		switch method {
		case Count, None:
			return nil
		}
	case NumericType:
		switch method {
		case Count, Mean, Sum, None:
			return nil
		}
	}
	return fmt.Errorf("you cannot aggregate with this method for this type, method: %s, type: %s", method, s.GetType())
}

func (s Series) Aggregate(method AggregationMethod) (element.Element, error) {

	if err := s.CanAggregateWith(method); err != nil {
		return nil, errors.Wrap(err, "")
	}

	switch method {
	case Count:
		return element.NewNumericElement(float64(s.Len()), false), nil
	case Mean:
		mean, err := s.Mean()
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		return element.NewNumericElement(mean, false), nil
	case Sum:
		sum, err := s.Sum()
		if err != nil {
			return nil, errors.Wrap(err, "failed to sum")
		}
		return element.NewNumericElement(sum, false), nil
	default:
		return s.GetGroupedElement()
	}
}

// Delete delete elements with keeping its schema
func (s Series) Delete() Series {
	s.Elements = s.Elements.Delete()
	return s
}

func (s Series) Append(s2 Series) (Series, error) {
	if s.GetName() != s2.GetName() {
		return Series{}, errors.New("series name is different")
	}
	if s.GetType() != s2.GetType() {
		return Series{}, errors.New("series type is different")
	}
	if s.GetAggregatedMethod() != s2.GetAggregatedMethod() {
		return Series{}, errors.New("series aggregated method is different")
	}
	elements, err := s.Elements.Append(s2.Elements)
	if err != nil {
		return Series{}, errors.Wrap(err, "failed to append elements")
	}
	return s.UpdateElements(elements)
}

// Slice make a series with sliced subsets of string elements
func (s Series) Slice(start int, end int) (Series, error) {
	stringListElements, ok := s.Elements.(element.StringListElements)
	if !ok {
		return Series{}, errors.New("this series is not string list elements")
	}
	newStringElements, err := stringListElements.Slice(start, end)
	if err != nil {
		return Series{}, errors.Wrap(err, "failed to slice elements")
	}
	return s.UpdateElements(newStringElements)
}

// Join string elements with separator
func (s Series) Join(separator string) (Series, error) {
	stringListElements, ok := s.Elements.(element.StringListElements)
	if !ok {
		return Series{}, errors.New("this series is not string list elements")
	}
	newStringElements := stringListElements.Join(separator)
	return s.UpdateElements(newStringElements)
}

func (s Series) GetAggregatedMethod() AggregationMethod {
	return s.AggregatedMethod
}
