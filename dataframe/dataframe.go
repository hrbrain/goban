package dataframe

import (
	"github.com/hrbrain/goban/series"
	"github.com/pkg/errors"
)

type DataFrame struct {
	Columns     Columns
	RecordCount int
}

func NewDataFrame(columns Columns) DataFrame {
	return DataFrame{
		Columns:     columns,
		RecordCount: columns.RecordCount(),
	}
}

func (df DataFrame) IncrementRecordCount(i int) DataFrame {
	df.RecordCount += i
	return df
}

func (df DataFrame) GetColumnByNameAndMethod(name series.Name, method series.AggregationMethod) (series.Series, error) {
	s, err := df.GetColumns().FindSeriesBy(name, method)
	if err != nil {
		return series.Series{}, errors.Wrap(err, "")
	}
	return s, nil
}

func (df DataFrame) GetColumnNames() []series.Name {
	return df.GetColumns().Names()
}

func (df DataFrame) GetColumns() Columns {
	return df.Columns
}

func (df DataFrame) UpdateColumn(series series.Series) (DataFrame, error) {
	columns, err := df.GetColumns().ReplaceByName(series)
	if err != nil {
		return DataFrame{}, errors.Wrap(err, "")
	}
	return df.UpdateColumns(columns), nil
}

func (df DataFrame) UpdateColumns(columns Columns) DataFrame {
	df.Columns = columns
	df.RecordCount = columns.RecordCount()
	return df
}

func (df DataFrame) GetRecordCount() int {
	return df.RecordCount
}

// Append concatenate two dataframes
func (df DataFrame) Append(df2 DataFrame) (DataFrame, error) {
	var newColumns Columns

	// check if columns length are same
	columns1 := df.GetColumns()
	columns2 := df2.GetColumns()
	if columns1.Len() != columns2.Len() {
		return DataFrame{}, errors.New("columns length mismatch")
	}

	for _, col1 := range columns1 {
		for _, col2 := range columns2 {
			if col1.GetName() == col2.GetName() && col1.GetAggregatedMethod() == col2.GetAggregatedMethod() {
				if col1.GetType() != col2.GetType() {
					return DataFrame{}, errors.New("column type mismatch")
				}
				newCol, err := col1.Append(col2)
				if err != nil {
					return DataFrame{}, errors.Wrap(err, "failed to append column")
				}
				newColumns, err = newColumns.Append(newCol)
				if err != nil {
					return DataFrame{}, errors.Wrap(err, "failed to append columns")
				}
			}
		}
	}
	return NewDataFrame(newColumns), nil
}

// Records convert dataframe into a list of records
func (df DataFrame) Records() (Records, error) {
	records := NewRecords(nil)
	columns := df.GetColumns()
	for i := 0; i < df.GetRecordCount(); i++ {
		record := NewRecord()
		for _, s := range columns {
			var err error
			element, err := s.GetElement(i)
			if err != nil {
				return nil, errors.Wrap(err, "")
			}
			record, err = record.AddField(s.GetName(), element)
			if err != nil {
				return nil, errors.Wrap(err, "")
			}
		}
		records = append(records, record)
	}
	return records, nil
}

// LoadRecord load a record into dataframe
func (df DataFrame) LoadRecord(r Record) (DataFrame, error) {
	newDataFrame := NewDataFrame(nil)
	for _, seriesName := range r.GetSeriesNames() {
		element, err := r.GetElement(seriesName)
		if err != nil {
			return DataFrame{}, errors.Wrap(err, "failed to get element")
		}
		s, err := df.GetColumnByNameAndMethod(seriesName, series.None)
		if err != nil {
			return DataFrame{}, errors.Wrap(err, "failed to get column")
		}
		s, err = s.AddElement(element)
		if err != nil {
			return DataFrame{}, errors.Wrap(err, "failed to add element")
		}
		newDataFrame, err = df.UpdateColumn(s)
		if err != nil {
			return DataFrame{}, errors.Wrap(err, "failed to update column")
		}
	}
	return newDataFrame, nil
}

// LoadRecords load records into dataframe
func (df DataFrame) LoadRecords(records Records) (DataFrame, error) {
	var newDataframe DataFrame
	for _, record := range records {
		var err error
		newDataframe, err = df.LoadRecord(record)
		if err != nil {
			return DataFrame{}, errors.Wrap(err, "")
		}
	}
	return newDataframe, nil
}

// GroupBy group by a designated column
func (df DataFrame) GroupBy(columnName series.Name) (Groups, error) {
	// convert the dataframe into a list of records
	records, err := df.Records()
	if err != nil {
		return Groups{}, errors.Wrap(err, "failed to convert to records")
	}

	// convert the list of records into a list of groups
	groups := NewGroups(nil)
	for _, r := range records {
		elem, err := r.GetElement(columnName)
		if err != nil {
			return Groups{}, errors.Wrap(err, "failed to get element")
		}
		dataFrame := groups.FindGroup(columnName, elem)
		if dataFrame.IsEmpty() {
			// if group does not exist, put the same structure dataframe as the original dataframe
			dataFrame, err = df.Delete()
			if err != nil {
				return Groups{}, errors.Wrap(err, "failed to delete dataframe")
			}
		}

		// add the record to the existing group
		dataFrame, err = dataFrame.LoadRecord(r)
		if err != nil {
			return Groups{}, errors.Wrap(err, "failed to load record")
		}
		groups = groups.Update(columnName, elem, dataFrame)
	}
	return groups, nil
}

// Delete delete data keeping its schema
func (df DataFrame) Delete() (DataFrame, error) {
	columns, err := df.GetColumns().Delete()
	if err != nil {
		return DataFrame{}, errors.Wrap(err, "")
	}
	return NewDataFrame(columns), nil
}

func (df DataFrame) Aggregate(conditions AggregationConditions) (DataFrame, error) {
	newColumns, err := NewColumns(nil)
	if err != nil {
		return DataFrame{}, errors.Wrap(err, "")
	}

	for _, condition := range conditions {
		col, err := df.GetColumnByNameAndMethod(condition.GetColumnName(), series.None) // only aggregate non aggregated column
		if err != nil {
			return DataFrame{}, errors.Wrap(err, "failed to find column")
		}
		aggregationMethod := condition.GetMethod()
		aggregatedValue, err := col.Aggregate(aggregationMethod)
		if err != nil {
			return DataFrame{}, errors.Wrap(err, "failed to aggregate series")
		}

		newSeries, err := series.NewSeries(col.GetName(), aggregatedValue.ToElements(), aggregationMethod)
		if err != nil {
			return DataFrame{}, errors.Wrap(err, "failed to make new series")
		}
		newColumns, err = newColumns.Append(newSeries)
		if err != nil {
			return DataFrame{}, errors.Wrap(err, "failed to append new column")
		}
	}
	return NewDataFrame(newColumns), nil
}

// DropNA Drop rows that have NA values
func (df DataFrame) DropNA(names ...series.Name) (DataFrame, error) {
	records, err := df.Records()
	if err != nil {
		return DataFrame{}, errors.Wrap(err, "failed to convert dataframe to records")
	}

	// Loop over each record and drop records with null element
	newRecords := NewRecords(nil)
	for _, r := range records {
		if r.HasNAElement(names...) {
			continue
		}
		newRecords = newRecords.Append(r)
	}

	// Convert new records to dataframe
	newDataframe, err := df.Delete()
	if err != nil {
		return DataFrame{}, errors.Wrap(err, "failed to delete columns")
	}
	newDataframe, err = newDataframe.LoadRecords(newRecords)
	if err != nil {
		return DataFrame{}, errors.Wrap(err, "failed to convert records to dataframe")
	}
	return newDataframe, nil
}

func (df DataFrame) IsEmpty() bool {
	return df.GetColumns().IsEmpty()
}
