package dataframe

import (
	"github.com/hrbrain/goban/element"
	"github.com/hrbrain/goban/series"
)

type Group struct {
	SeriesName series.Name
	Element    element.Element
	DataFrame  DataFrame
}

func NewGroup(name series.Name, element element.Element, dataframe DataFrame) Group {
	return Group{
		SeriesName: name,
		Element:    element,
		DataFrame:  dataframe,
	}
}

func (g Group) GetElement() element.Element {
	return g.Element
}

func (g Group) GetSeriesName() series.Name {
	return g.SeriesName
}

func (g Group) GetDataframe() DataFrame {
	return g.DataFrame
}
