package dataframe

import (
	"github.com/hrbrain/goban/series"
)

type AggregationCondition struct {
	ColumnName series.Name
	Method     series.AggregationMethod
}

func NewAggregationCondition(columnName series.Name, method series.AggregationMethod) AggregationCondition {
	return AggregationCondition{
		ColumnName: columnName,
		Method:     method,
	}
}

func (ac AggregationCondition) GetMethod() series.AggregationMethod {
	return ac.Method
}

func (ac AggregationCondition) GetColumnName() series.Name {
	return ac.ColumnName
}
