package dataframe

import (
	"fmt"

	"github.com/hrbrain/goban/element"
	"github.com/hrbrain/goban/series"
)

type Record map[series.Name]element.Element

func NewRecord() Record {
	return map[series.Name]element.Element{}
}

// GetSeriesNames does not guarantee the order of series names
func (r Record) GetSeriesNames() []series.Name {
	names := make([]series.Name, 0, len(r))
	for name := range r {
		names = append(names, name)
	}
	return names
}

func (r Record) Has(n series.Name) bool {
	if _, ok := r[n]; ok {
		return true
	}
	return false
}

func (r Record) HasNAElement(names ...series.Name) bool {
	if len(names) == 0 {
		names = r.GetSeriesNames()
	}
	for _, name := range names {
		if v, ok := r[name]; ok && v.IsNA() {
			return true
		}
	}
	return false
}

func (r Record) AddField(name series.Name, e element.Element) (Record, error) {
	if r.Has(name) {
		return Record{}, fmt.Errorf("duplicated series name, name: %s", name)
	}
	r[name] = e
	return r, nil
}

func (r Record) GetElement(name series.Name) (element.Element, error) {
	if r.Has(name) {
		return r[name], nil
	}
	return nil, fmt.Errorf("column name not found in record, name: %s", name)
}
