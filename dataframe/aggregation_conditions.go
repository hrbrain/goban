package dataframe

import (
	"fmt"

	"github.com/hrbrain/goban/series"
)

type AggregationConditions []AggregationCondition

func (acs AggregationConditions) Len() int {
	return len(acs)
}

func (acs AggregationConditions) GetCondition(index int) (AggregationCondition, error) {
	if index < 0 || acs.Len() <= index {
		return AggregationCondition{}, fmt.Errorf("index out of range, index: %d", index)
	}
	return acs[index], nil
}

func (acs AggregationConditions) GetConditionByName(name series.Name) (AggregationCondition, error) {
	for _, ac := range acs {
		if ac.GetColumnName() == name {
			return ac, nil
		}
	}
	return AggregationCondition{}, fmt.Errorf("aggregation condition not found, name: %s", name)
}

func (acs AggregationConditions) Append(ac AggregationCondition) AggregationConditions {
	return append(acs, ac)
}
