package dataframe

import (
	"fmt"

	"github.com/hrbrain/goban/series"
)

// Columns are a list of series
type Columns []series.Series

func NewColumns(series []series.Series) (Columns, error) {
	var recordCount int
	for i, s := range series {
		if i == 0 {
			recordCount = s.Len()
			continue
		}
		if recordCount != s.Len() {
			return nil, fmt.Errorf("series length mismatch recordCount: %d, series.len(): %d", recordCount, s.Len())
		}
	}
	return series, nil
}

func (columns Columns) Len() int {
	return len(columns)
}

func (columns Columns) RecordCount() int {
	if columns.Len() == 0 {
		return 0
	}
	return columns[0].Len()
}

func (columns Columns) Names() []series.Name {
	columnNames := make([]series.Name, 0, columns.Len())
	for _, series := range columns {
		columnNames = append(columnNames, series.GetName())
	}
	return columnNames
}

// Append new series into columns
func (columns Columns) Append(series series.Series) (Columns, error) {
	// no need to check length when columns are empty
	if columns.Len() == 0 {
		return append(columns, series), nil
	}

	// check if the length of the series is the same as the length of the columns
	if columns.RecordCount() != series.Len() {
		return nil, fmt.Errorf("record count mismatch, columns.RecordCount(): %d, series.Len(): %d", columns.RecordCount(), series.Len())
	}

	// check if the name and aggregation method of the series is duplicated
	for _, c := range columns {
		if c.GetName() == series.GetName() && c.GetAggregatedMethod() == series.GetAggregatedMethod() {
			return nil, fmt.Errorf("duplicate column name and aggregated method, name: %s, aggregatedMethod: %s", series.GetName(), series.GetAggregatedMethod())
		}
	}
	return append(columns, series), nil
}

func (columns Columns) isValidIndex(index int) bool {
	return 0 <= index && index < columns.Len()
}

func (columns Columns) GetSeries(index int) (series.Series, error) {
	if !columns.isValidIndex(index) {
		return series.Series{}, fmt.Errorf("index out of range, index: %d", index)
	}
	return columns[index], nil
}

func (columns Columns) HasSeriesName(name series.Name) (series.Series, bool) {
	for i, nameInColumns := range columns.Names() {
		if nameInColumns == name {
			return columns[i], true
		}
	}
	return series.Series{}, false
}

func (columns Columns) FindSeriesBy(name series.Name, method series.AggregationMethod) (series.Series, error) {
	for _, column := range columns {
		if column.GetName() == name && column.GetAggregatedMethod() == method {
			return column, nil
		}
	}
	return series.Series{}, fmt.Errorf("series name not found in columns, name: %s", name)
}

// Replace column by index.
// Length of replacing series and length of other columns can be inconsistent
func (columns Columns) Replace(index int, series series.Series) (Columns, error) {
	if !columns.isValidIndex(index) {
		return nil, fmt.Errorf("index out of range, index: %d", index)
	}
	columns[index] = series
	return columns, nil
}

// ReplaceByName replace column by its name.
// Length of replacing series and length of other columns can be inconsistent
func (columns Columns) ReplaceByName(s series.Series) (Columns, error) {
	for i, nameInColumns := range columns.Names() {
		if nameInColumns == s.GetName() {
			columns[i] = s
			return columns, nil
		}
	}
	return Columns{}, fmt.Errorf("series name not found, s.Name: %s", s.GetName())
}

// Delete delete data keeping its schema
func (columns Columns) Delete() (Columns, error) {
	ss := make([]series.Series, columns.Len())
	for i, s := range columns {
		ss[i] = s.Delete()
	}
	return NewColumns(ss)
}

func (columns Columns) IsEmpty() bool {
	return columns.Len() == 0
}
