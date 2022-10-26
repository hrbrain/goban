package dataframe

import (
	"github.com/hrbrain/goban/element"
	"github.com/hrbrain/goban/series"
	"github.com/pkg/errors"
)

type Groups []Group

func NewGroups(groups []Group) Groups {
	return groups
}

func (groups Groups) Update(name series.Name, element element.Element, dataFrame DataFrame) Groups {
	for i, g := range groups {
		if g.GetSeriesName() == name && g.GetElement().Equal(element) {
			groups[i].DataFrame = dataFrame
			return groups
		}
	}
	newGroup := NewGroup(name, element, dataFrame)
	return groups.Append(newGroup)
}

func (groups Groups) FindGroup(name series.Name, element element.Element) DataFrame {
	for _, group := range groups {
		if group.SeriesName == name && group.Element.Equal(element) {
			return group.DataFrame
		}
	}
	return DataFrame{}
}

func (groups Groups) Len() int {
	return len(groups)
}

func (groups Groups) Aggregate(conditions AggregationConditions) (DataFrame, error) {
	var AggregatedDataframe DataFrame
	for i, group := range groups {
		AggregatedDataframeInGroup, err := group.GetDataframe().Aggregate(conditions)
		if err != nil {
			return DataFrame{}, errors.Wrap(err, "failed to aggregate dataframe")
		}

		if i == 0 {
			AggregatedDataframe = AggregatedDataframeInGroup
			continue
		}

		AggregatedDataframe, err = AggregatedDataframe.Append(AggregatedDataframeInGroup)
		if err != nil {
			return DataFrame{}, errors.Wrap(err, "failed to append dataframe")
		}
	}
	return AggregatedDataframe, nil
}

func (groups Groups) Append(group Group) Groups {
	return append(groups, group)
}
